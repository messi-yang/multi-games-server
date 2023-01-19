package livegamecontroller

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/messaging/redisintgreventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/gzipprovider"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/jsonmarshaller"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Controller struct {
	gameRepo           gamemodel.GameRepo
	liveGameAppService livegameappservice.Service
	playerRepo         playermodel.Repo
}

func NewController(
	gameRepo gamemodel.GameRepo,
	liveGameAppService livegameappservice.Service,
	playerRepo playermodel.Repo,
) *Controller {
	return &Controller{
		gameRepo:           gameRepo,
		liveGameAppService: liveGameAppService,
		playerRepo:         playerRepo,
	}
}

func (controller *Controller) HandleLiveGameConnection(c *gin.Context) {
	socketConn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Error(err)
		return
	}
	defer socketConn.Close()
	closeConnFlag := make(chan bool)

	games, err := controller.gameRepo.GetAll()
	if err != nil {
		return
	}
	gameId := games[0].GetId()

	liveGameId := gameId.ToString()

	playerIdVm := uuid.New().String()
	playerId, _ := playermodel.NewPlayerIdVo(playerIdVm)
	myPlayer := playermodel.NewPlayerAgg(playerId, "Hello")
	controller.playerRepo.Add(myPlayer)

	socketPresenter := newSocketPresenter(socketConn, &sync.RWMutex{})
	gzipCompressor := gzipprovider.New()

	go controller.liveGameAppService.SendItemsUpdatedServerEvent(socketPresenter)

	intgrEventSubscriberUnsubscriber := redisintgreventsubscriber.New().Subscribe(
		intgrevent.CreateLiveGameClientChannel(liveGameId, playerIdVm),
		func(message []byte) {
			intgrEvent, err := jsonmarshaller.Unmarshal[intgrevent.GenericIntgrEvent](message)
			if err != nil {
				return
			}

			switch intgrEvent.Name {
			case intgrevent.GameJoinedIntgrEventName:
				event, err := jsonmarshaller.Unmarshal[intgrevent.GameJoinedIntgrEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendGameJoinedServerEvent(socketPresenter, viewmodel.NewPlayerVm(myPlayer), event.Camera, event.MapSize, event.View)
			case intgrevent.CameraChangedIntgrEventName:
				event, err := jsonmarshaller.Unmarshal[intgrevent.CameraChangedIntgrEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendCameraChangedServerEvent(socketPresenter, event.Camera)
			case intgrevent.ViewChangedIntgrEventName:
				event, err := jsonmarshaller.Unmarshal[intgrevent.ViewChangedIntgrEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendViewChangedServerEvent(socketPresenter, event.View)
			case intgrevent.ViewUpdatedIntgrEventName:
				event, err := jsonmarshaller.Unmarshal[intgrevent.ViewUpdatedIntgrEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendViewUpdatedServerEvent(socketPresenter, event.View)
			}

		})
	defer intgrEventSubscriberUnsubscriber()

	controller.liveGameAppService.RequestToJoinGame(liveGameId, playerIdVm)

	go func() {
		defer func() {
			closeConnFlag <- true
		}()

		for {
			_, compressedMessage, err := socketConn.ReadMessage()
			if err != nil {
				controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
				return
			}

			message, err := gzipCompressor.Ungzip(compressedMessage)
			if err != nil {
				controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
				continue
			}

			commandType, err := livegameappservice.ParseClientEventType(message)
			if err != nil {
				controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
				continue
			}

			switch commandType {
			case livegameappservice.PingClientEventType:
				continue
			case livegameappservice.ChangeCameraClientEventType:
				command, err := livegameappservice.ParseClientEvent[livegameappservice.ChangeCameraClientEvent](message)
				if err != nil {
					controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
					continue
				}
				controller.liveGameAppService.RequestToChangeCamera(liveGameId, playerIdVm, command.Payload.Camera)
			case livegameappservice.BuildItemClientEventType:
				command, err := livegameappservice.ParseClientEvent[livegameappservice.BuildItemClientEvent](message)
				if err != nil {
					controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
					continue
				}

				controller.liveGameAppService.RequestToBuildItem(liveGameId, command.Payload.Location, command.Payload.ItemId)
			case livegameappservice.DestroyItemClientEventType:
				command, err := livegameappservice.ParseClientEvent[livegameappservice.DestroyItemClientEvent](message)
				if err != nil {
					controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
					continue
				}

				controller.liveGameAppService.RequestToDestroyItem(liveGameId, command.Payload.Location)
			default:
			}
		}
	}()

	for {
		<-closeConnFlag

		controller.playerRepo.Remove(myPlayer.GetId())
		controller.liveGameAppService.RequestToLeaveGame(liveGameId, playerIdVm)
		return
	}
}
