package socketcontroller

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/redissub"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/gzipper"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/jsonmarshaller"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/samber/lo"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type LiveGameSocketController struct {
	gameRepo           gamemodel.GameRepo
	liveGameAppService appservice.LiveGameAppService
}

func NewLiveGameSocketController(
	gameRepo gamemodel.GameRepo,
	liveGameAppService appservice.LiveGameAppService,
) *LiveGameSocketController {
	return &LiveGameSocketController{
		gameRepo:           gameRepo,
		liveGameAppService: liveGameAppService,
	}
}

func (controller *LiveGameSocketController) HandleLiveGameConnection(c *gin.Context) {
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

	socketPresenter := newSocketPresenter(socketConn, &sync.RWMutex{})

	go controller.liveGameAppService.SendItemsUpdatedServerEvent(socketPresenter)

	intEventUnsubscriber := redissub.New().Subscribe(
		intevent.CreateLiveGameClientChannel(liveGameId, playerIdVm),
		func(message []byte) {
			intEvent, err := jsonmarshaller.Unmarshal[intevent.GenericIntEvent](message)
			if err != nil {
				return
			}

			switch intEvent.Name {
			case intevent.GameJoinedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.GameJoinedIntEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendGameJoinedServerEvent(socketPresenter, event.Player, event.MapSize, event.View)
			case intevent.PlayerUpdatedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.PlayerUpdatedIntEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendPlayerUpdatedServerEvent(socketPresenter, event.Player)
			case intevent.PlayersUpdatedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.PlayersUpdatedIntEvent](message)
				if err != nil {
					return
				}
				myPlayerVm, exists := lo.Find(event.Players, func(playerVm viewmodel.PlayerVm) bool {
					return playerVm.Id == playerIdVm
				})
				if !exists {
					return
				}

				otherPlayerVms := lo.Filter(event.Players, func(playerVm viewmodel.PlayerVm, _ int) bool {
					return playerVm.Id != playerIdVm
				})
				controller.liveGameAppService.SendPlayersUpdatedServerEvent(socketPresenter, myPlayerVm, otherPlayerVms)
			case intevent.ViewUpdatedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.ViewUpdatedIntEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendViewUpdatedServerEvent(socketPresenter, event.View)
			}

		})
	defer intEventUnsubscriber()

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

			genericClientEvent, err := jsonmarshaller.Unmarshal[appservice.GenericClientEvent](message)
			if err != nil {
				controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
				continue
			}

			switch genericClientEvent.Type {
			case appservice.PingClientEventType:
				continue
			case appservice.ChangeCameraClientEventType:
				command, err := jsonmarshaller.Unmarshal[appservice.ChangeCameraClientEvent](message)
				if err != nil {
					controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
					continue
				}
				controller.liveGameAppService.RequestToChangeCamera(liveGameId, playerIdVm, command.Payload.Camera)
			case appservice.BuildItemClientEventType:
				command, err := jsonmarshaller.Unmarshal[appservice.BuildItemClientEvent](message)
				if err != nil {
					controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
					continue
				}

				controller.liveGameAppService.RequestToBuildItem(liveGameId, command.Payload.Location, command.Payload.ItemId)
			case appservice.DestroyItemClientEventType:
				command, err := jsonmarshaller.Unmarshal[appservice.DestroyItemClientEvent](message)
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

		controller.liveGameAppService.RequestToLeaveGame(liveGameId, playerIdVm)
		return
	}
}
