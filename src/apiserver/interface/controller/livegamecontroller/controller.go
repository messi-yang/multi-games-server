package livegamecontroller

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/redissub"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/gzipper"
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

	go controller.liveGameAppService.SendItemsUpdatedServerEvent(socketPresenter)

	IntEventSubscriberUnsubscriber := redissub.New().Subscribe(
		intevent.CreateLiveGameClientChannel(liveGameId, playerIdVm),
		func(message []byte) {
			intEvent, err := jsonmarshaller.Unmarshal[intevent.GenericintEvent](message)
			if err != nil {
				return
			}

			switch intEvent.Name {
			case intevent.GameJoinedintEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.GameJoinedintEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendGameJoinedServerEvent(socketPresenter, viewmodel.NewPlayerVm(myPlayer), event.Camera, event.MapSize, event.View)
			case intevent.CameraChangedintEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.CameraChangedintEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendCameraChangedServerEvent(socketPresenter, event.Camera)
			case intevent.ViewChangedintEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.ViewChangedintEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendViewChangedServerEvent(socketPresenter, event.View)
			case intevent.ViewUpdatedintEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.ViewUpdatedintEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendViewUpdatedServerEvent(socketPresenter, event.View)
			}

		})
	defer IntEventSubscriberUnsubscriber()

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

			message, err := gzipper.Ungzip(compressedMessage)
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
