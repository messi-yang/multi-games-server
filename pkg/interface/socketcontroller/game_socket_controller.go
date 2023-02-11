package socketcontroller

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/redissub"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/library/gzipper"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/library/jsonmarshaller"
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

type GameSocketController struct {
	gameAppService appservice.GameAppService
}

func NewGameSocketController(
	gameAppService appservice.GameAppService,
) *GameSocketController {
	return &GameSocketController{
		gameAppService: gameAppService,
	}
}

func (controller *GameSocketController) HandleGameConnection(c *gin.Context) {
	socketConn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Error(err)
		return
	}
	defer socketConn.Close()
	closeConnFlag := make(chan bool)

	gameIdVm := "20716447-6514-4eac-bd05-e558ca72bf3c"

	playerIdVm := uuid.New().String()

	socketPresenter := newSocketPresenter(socketConn, &sync.RWMutex{})

	go controller.gameAppService.SendItemsUpdatedServerEvent(socketPresenter)

	intEventUnsubscriber := redissub.New().Subscribe(
		intevent.CreateGameClientChannel(gameIdVm, playerIdVm),
		func(message []byte) {
			intEvent, err := jsonmarshaller.Unmarshal[intevent.GenericIntEvent](message)
			if err != nil {
				return
			}

			switch intEvent.Name {
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
				controller.gameAppService.SendPlayersUpdatedServerEvent(socketPresenter, myPlayerVm, otherPlayerVms)
			case intevent.ViewUpdatedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.ViewUpdatedIntEvent](message)
				if err != nil {
					return
				}
				controller.gameAppService.SendViewUpdatedServerEvent(socketPresenter, event.View)
			}

		})
	defer intEventUnsubscriber()

	controller.gameAppService.JoinGame(socketPresenter, gameIdVm, playerIdVm)

	go func() {
		defer func() {
			closeConnFlag <- true
		}()

		for {
			_, compressedMessage, err := socketConn.ReadMessage()
			if err != nil {
				controller.gameAppService.SendErroredServerEvent(socketPresenter, err.Error())
				return
			}

			message, err := gzipper.Ungzip(compressedMessage)
			if err != nil {
				controller.gameAppService.SendErroredServerEvent(socketPresenter, err.Error())
				continue
			}

			genericClientEvent, err := jsonmarshaller.Unmarshal[appservice.GenericClientEvent](message)
			if err != nil {
				controller.gameAppService.SendErroredServerEvent(socketPresenter, err.Error())
				continue
			}

			switch genericClientEvent.Type {
			case appservice.PingClientEventType:
				continue
			case appservice.MoveClientEventType:
				command, err := jsonmarshaller.Unmarshal[appservice.MoveClientEvent](message)
				if err != nil {
					controller.gameAppService.SendErroredServerEvent(socketPresenter, err.Error())
					continue
				}
				controller.gameAppService.MovePlayer(gameIdVm, playerIdVm, command.Payload.Direction)
			case appservice.PlaceItemClientEventType:
				command, err := jsonmarshaller.Unmarshal[appservice.PlaceItemClientEvent](message)
				if err != nil {
					controller.gameAppService.SendErroredServerEvent(socketPresenter, err.Error())
					continue
				}

				controller.gameAppService.PlaceItem(gameIdVm, playerIdVm, command.Payload.Location, command.Payload.ItemId)
			case appservice.DestroyItemClientEventType:
				command, err := jsonmarshaller.Unmarshal[appservice.DestroyItemClientEvent](message)
				if err != nil {
					controller.gameAppService.SendErroredServerEvent(socketPresenter, err.Error())
					continue
				}

				controller.gameAppService.DestroyItem(gameIdVm, playerIdVm, command.Payload.Location)
			default:
			}
		}
	}()

	for {
		<-closeConnFlag

		controller.gameAppService.LeaveGame(gameIdVm, playerIdVm)
		return
	}
}
