package gamesocketapi

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/gzip"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/json"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gamesocketappservice"
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

type Controller struct {
	gameAppService gamesocketappservice.Service
}

func NewController(
	gameAppService gamesocketappservice.Service,
) *Controller {
	return &Controller{
		gameAppService: gameAppService,
	}
}

func (controller *Controller) HandleGameConnection(c *gin.Context) {
	socketConn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Error(err)
		return
	}
	defer socketConn.Close()
	closeConnFlag := make(chan bool)

	gameIdDto := "20716447-6514-4eac-bd05-e558ca72bf3c"

	playerIdDto := uuid.New().String()

	socketPresenter := newPresenter(socketConn, &sync.RWMutex{})

	intEventUnsubscriber := redissub.New().Subscribe(
		gamesocketappservice.CreateGamePlayerChannel(gameIdDto, playerIdDto),
		func(message []byte) {
			intEvent, err := json.Unmarshal[gamesocketappservice.GameSocketIntEvent](message)
			if err != nil {
				return
			}

			switch intEvent.Name {
			case gamesocketappservice.PlayersUpdatedGameSocketIntEventName:
				event, _ := json.Unmarshal[gamesocketappservice.PlayersUpdatedIntEvent](message)

				query, err := gamesocketappservice.NewGetPlayersQuery(event.GameId, playerIdDto)
				if err != nil {
					return
				}

				controller.gameAppService.GetPlayers(socketPresenter, query)
			case gamesocketappservice.ViewUpdatedGameSocketIntEventName:
				event, _ := json.Unmarshal[gamesocketappservice.ViewUpdatedIntEvent](message)

				query, err := gamesocketappservice.NewGetViewQuery(event.GameId, playerIdDto)
				if err != nil {
					return
				}

				controller.gameAppService.GetView(socketPresenter, query)
			}

		})
	defer intEventUnsubscriber()

	command, err := gamesocketappservice.NewAddPlayerCommand(gameIdDto, playerIdDto)
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
				controller.gameAppService.GetError(socketPresenter, err.Error())
				return
			}

			message, err := gzip.Ungzip(compressedMessage)
			if err != nil {
				controller.gameAppService.GetError(socketPresenter, err.Error())
				continue
			}

			genericRequestDto, err := json.Unmarshal[gamesocketappservice.GenericRequestDto](message)
			if err != nil {
				controller.gameAppService.GetError(socketPresenter, err.Error())
				continue
			}

			switch genericRequestDto.Type {
			case gamesocketappservice.PingRequestDtoType:
				continue
			case gamesocketappservice.MoveRequestDtoType:
				requestDto, err := json.Unmarshal[gamesocketappservice.MoveRequestDto](message)
				if err != nil {
					controller.gameAppService.GetError(socketPresenter, err.Error())
					continue
				}

				command, err := gamesocketappservice.NewMovePlayerCommand(gameIdDto, playerIdDto, requestDto.Direction)
				if err != nil {
					controller.gameAppService.GetError(socketPresenter, err.Error())
					continue
				}

				controller.gameAppService.MovePlayer(socketPresenter, command)
			case gamesocketappservice.PlaceItemRequestDtoType:
				requestDto, err := json.Unmarshal[gamesocketappservice.PlaceItemRequestDto](message)
				if err != nil {
					controller.gameAppService.GetError(socketPresenter, err.Error())
					continue
				}

				command, err := gamesocketappservice.NewPlaceItemCommand(gameIdDto, playerIdDto, requestDto.Location, requestDto.ItemId)
				if err != nil {
					controller.gameAppService.GetError(socketPresenter, err.Error())
					continue
				}

				controller.gameAppService.PlaceItem(command)
			case gamesocketappservice.DestroyItemRequestDtoType:
				requestDto, err := json.Unmarshal[gamesocketappservice.DestroyItemRequestDto](message)
				if err != nil {
					controller.gameAppService.GetError(socketPresenter, err.Error())
					continue
				}

				command, err := gamesocketappservice.NewDestroyItemCommand(gameIdDto, playerIdDto, requestDto.Location)
				if err != nil {
					controller.gameAppService.GetError(socketPresenter, err.Error())
					continue
				}

				controller.gameAppService.DestroyItem(command)
			default:
			}
		}
	}()

	for {
		<-closeConnFlag

		command, _ := gamesocketappservice.NewRemovePlayerCommand(gameIdDto, playerIdDto)
		controller.gameAppService.RemovePlayer(command)
		return
	}
}
