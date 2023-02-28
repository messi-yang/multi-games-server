package gamesocketapi

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/gzip"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/json"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gamesocketappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/messaging/intevent/redisinteventsubscriber"
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
	disconnect := func() {
		closeConnFlag <- true
	}

	gameIdDto := c.Request.URL.Query().Get("id")

	playerIdDto := uuid.New().String()

	socketPresenter := newPresenter(socketConn, &sync.RWMutex{})

	playersUpdatedIntEventUnsubscriber := redisinteventsubscriber.New[gamesocketappservice.PlayersUpdatedIntEvent]().Subscribe(
		gamesocketappservice.NewPlayersUpdatedIntEventChannel(gameIdDto, playerIdDto),
		func(intEvent gamesocketappservice.PlayersUpdatedIntEvent) {
			controller.gameAppService.GetPlayersAroundPlayer(socketPresenter, gamesocketappservice.GetPlayersQuery{
				GameId:   gameIdDto,
				PlayerId: playerIdDto,
			})
		},
	)
	defer playersUpdatedIntEventUnsubscriber()

	unitsUpdatedIntEventTypeUnsubscriber := redisinteventsubscriber.New[gamesocketappservice.UnitsUpdatedIntEvent]().Subscribe(
		gamesocketappservice.NewUnitsUpdatedIntEventChannel(gameIdDto, playerIdDto),
		func(intEvent gamesocketappservice.UnitsUpdatedIntEvent) {
			controller.gameAppService.GetUnitsVisibleByPlayer(socketPresenter, gamesocketappservice.GetUnitsVisibleByPlayerQuery{
				GameId:   gameIdDto,
				PlayerId: playerIdDto,
			})
		},
	)
	defer unitsUpdatedIntEventTypeUnsubscriber()

	controller.gameAppService.AddPlayer(socketPresenter, gamesocketappservice.AddPlayerCommand{
		GameId:   gameIdDto,
		PlayerId: playerIdDto,
	})

	go func() {
		defer disconnect()

		for {
			_, compressedMessage, err := socketConn.ReadMessage()
			if err != nil {
				disconnect()
			}

			message, err := gzip.Ungzip(compressedMessage)
			if err != nil {
				disconnect()
			}

			genericRequestDto, err := json.Unmarshal[gamesocketappservice.GenericRequestDto](message)
			if err != nil {
				disconnect()
			}

			switch genericRequestDto.Type {
			case gamesocketappservice.PingRequestDtoType:
				continue
			case gamesocketappservice.MoveRequestDtoType:
				requestDto, err := json.Unmarshal[gamesocketappservice.MoveRequestDto](message)
				if err != nil {
					disconnect()
				}

				controller.gameAppService.MovePlayer(socketPresenter, gamesocketappservice.MovePlayerCommand{
					GameId:    gameIdDto,
					PlayerId:  playerIdDto,
					Direction: requestDto.Direction,
				})
			case gamesocketappservice.PlaceItemRequestDtoType:
				requestDto, err := json.Unmarshal[gamesocketappservice.PlaceItemRequestDto](message)
				if err != nil {
					disconnect()
				}

				controller.gameAppService.PlaceItem(socketPresenter, gamesocketappservice.PlaceItemCommand{
					GameId:   gameIdDto,
					PlayerId: playerIdDto,
					ItemId:   requestDto.ItemId,
				})
			case gamesocketappservice.DestroyItemRequestDtoType:
				controller.gameAppService.DestroyItem(socketPresenter, gamesocketappservice.DestroyItemCommand{
					GameId:   gameIdDto,
					PlayerId: playerIdDto,
				})
			default:
			}
		}
	}()

	for {
		<-closeConnFlag

		controller.gameAppService.RemovePlayer(socketPresenter, gamesocketappservice.RemovePlayerCommand{
			GameId:   gameIdDto,
			PlayerId: playerIdDto,
		})
		return
	}
}
