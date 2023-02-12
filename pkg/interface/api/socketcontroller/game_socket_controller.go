package socketcontroller

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/library/gzipper"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/library/jsonmarshaller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gamesocketservice"
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
	gameAppService gamesocketservice.Service
}

func NewGameSocketController(
	gameAppService gamesocketservice.Service,
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
		gamesocketservice.CreateGameIntEventChannel(gameIdVm),
		func(message []byte) {
			intEvent, err := jsonmarshaller.Unmarshal[gamesocketservice.GameSocketIntEvent](message)
			if err != nil {
				return
			}

			switch intEvent.Name {
			case gamesocketservice.PlayerUpdatedGameSocketIntEventName:
				event, err := jsonmarshaller.Unmarshal[gamesocketservice.PlayerUpdatedIntEvent](message)
				if err != nil {
					return
				}

				controller.gameAppService.HandlePlayerUpdatedEvent(socketPresenter, event)
			case gamesocketservice.UnitUpdatedGameSocketIntEventName:
				event, err := jsonmarshaller.Unmarshal[gamesocketservice.UnitUpdatedIntEvent](message)
				if err != nil {
					return
				}
				controller.gameAppService.HandleUnitUpdatedEvent(socketPresenter, playerIdVm, event)
			}

		})
	defer intEventUnsubscriber()

	command, err := gamesocketservice.NewAddPlayerCommand(gameIdVm, playerIdVm)
	if err != nil {
		return
	}
	controller.gameAppService.AddPlayer(socketPresenter, command)

	go func() {
		defer func() {
			closeConnFlag <- true
		}()

		for {
			_, compressedMessage, err := socketConn.ReadMessage()
			if err != nil {
				controller.gameAppService.SendErroredResponseDto(socketPresenter, err.Error())
				return
			}

			message, err := gzipper.Ungzip(compressedMessage)
			if err != nil {
				controller.gameAppService.SendErroredResponseDto(socketPresenter, err.Error())
				continue
			}

			genericRequestDto, err := jsonmarshaller.Unmarshal[gamesocketservice.RequestDto](message)
			if err != nil {
				controller.gameAppService.SendErroredResponseDto(socketPresenter, err.Error())
				continue
			}

			switch genericRequestDto.Type {
			case gamesocketservice.PingRequestDtoType:
				continue
			case gamesocketservice.MoveRequestDtoType:
				requestDto, err := jsonmarshaller.Unmarshal[gamesocketservice.MoveRequestDto](message)
				if err != nil {
					controller.gameAppService.SendErroredResponseDto(socketPresenter, err.Error())
					continue
				}

				command, err := gamesocketservice.NewMovePlayerCommand(gameIdVm, playerIdVm, requestDto.Direction)
				if err != nil {
					controller.gameAppService.SendErroredResponseDto(socketPresenter, err.Error())
					continue
				}

				controller.gameAppService.MovePlayer(socketPresenter, command)
			case gamesocketservice.PlaceItemRequestDtoType:
				requestDto, err := jsonmarshaller.Unmarshal[gamesocketservice.PlaceItemRequestDto](message)
				if err != nil {
					controller.gameAppService.SendErroredResponseDto(socketPresenter, err.Error())
					continue
				}

				command, err := gamesocketservice.NewPlaceItemCommand(gameIdVm, playerIdVm, requestDto.Location, requestDto.ItemId)
				if err != nil {
					controller.gameAppService.SendErroredResponseDto(socketPresenter, err.Error())
					continue
				}

				controller.gameAppService.PlaceItem(command)
			case gamesocketservice.DestroyItemRequestDtoType:
				requestDto, err := jsonmarshaller.Unmarshal[gamesocketservice.DestroyItemRequestDto](message)
				if err != nil {
					controller.gameAppService.SendErroredResponseDto(socketPresenter, err.Error())
					continue
				}

				command, err := gamesocketservice.NewDestroyItemCommand(gameIdVm, playerIdVm, requestDto.Location)
				if err != nil {
					controller.gameAppService.SendErroredResponseDto(socketPresenter, err.Error())
					continue
				}

				controller.gameAppService.DestroyItem(command)
			default:
			}
		}
	}()

	for {
		<-closeConnFlag

		command, _ := gamesocketservice.NewRemovePlayerCommand(gameIdVm, playerIdVm)
		controller.gameAppService.RemovePlayer(command)
		return
	}
}
