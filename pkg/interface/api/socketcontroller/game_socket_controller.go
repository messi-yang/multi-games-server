package socketcontroller

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/library/gzipper"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/library/jsonmarshaller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/messaging/redissub"
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

	intEventUnsubscriber := redissub.New().Subscribe(
		intevent.CreateGameChannel(gameIdVm),
		func(message []byte) {
			intEvent, err := jsonmarshaller.Unmarshal[intevent.GenericIntEvent](message)
			if err != nil {
				return
			}

			switch intEvent.Name {
			case intevent.PlayerUpdatedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.PlayerUpdatedIntEvent](message)
				if err != nil {
					return
				}

				controller.gameAppService.HandlePlayerUpdatedEvent(socketPresenter, event)
			case intevent.UnitUpdatedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.UnitUpdatedIntEvent](message)
				if err != nil {
					return
				}
				controller.gameAppService.HandleUnitUpdatedEvent(socketPresenter, playerIdVm, event)
			}

		})
	defer intEventUnsubscriber()

	controller.gameAppService.AddPlayer(socketPresenter, gameIdVm, playerIdVm)

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
				controller.gameAppService.MovePlayer(socketPresenter, gameIdVm, playerIdVm, command.Payload.Direction)
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

		controller.gameAppService.RemovePlayer(gameIdVm, playerIdVm)
		return
	}
}
