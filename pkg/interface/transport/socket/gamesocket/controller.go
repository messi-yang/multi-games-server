package gamesocket

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/gzip"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/json"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gamesocketappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/messaging/redisinteventsubscriber"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
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
	redisClient    *redis.Client
}

func NewController(
	gameAppService gamesocketappservice.Service,
	redisClient *redis.Client,
) *Controller {
	return &Controller{
		gameAppService: gameAppService,
		redisClient:    redisClient,
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

	worldIdDto := c.Request.URL.Query().Get("id")

	playerIdDto := uuid.New().String()

	socketPresenter := newPresenter(socketConn, &sync.RWMutex{})

	playersUpdatedIntEventUnsubscriber := redisinteventsubscriber.New[gamesocketappservice.PlayersUpdatedIntEvent](
		controller.redisClient,
	).Subscribe(
		gamesocketappservice.NewPlayersUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gamesocketappservice.PlayersUpdatedIntEvent) {
			controller.gameAppService.GetPlayersAroundPlayer(socketPresenter, gamesocketappservice.GetPlayersQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			})
		},
	)
	defer playersUpdatedIntEventUnsubscriber()

	unitsUpdatedIntEventTypeUnsubscriber := redisinteventsubscriber.New[gamesocketappservice.UnitsUpdatedIntEvent](
		controller.redisClient,
	).Subscribe(
		gamesocketappservice.NewUnitsUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gamesocketappservice.UnitsUpdatedIntEvent) {
			controller.gameAppService.GetUnitsVisibleByPlayer(socketPresenter, gamesocketappservice.GetUnitsVisibleByPlayerQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			})
		},
	)
	defer unitsUpdatedIntEventTypeUnsubscriber()

	controller.gameAppService.AddPlayer(socketPresenter, gamesocketappservice.AddPlayerCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	})

	go func() {
		defer disconnect()

		for {
			_, compressedMessage, err := socketConn.ReadMessage()
			if err != nil {
				return
			}

			message, err := gzip.Ungzip(compressedMessage)
			if err != nil {
				return
			}

			genericRequestDto, err := json.Unmarshal[gamesocketappservice.GenericRequestDto](message)
			if err != nil {
				return
			}

			switch genericRequestDto.Type {
			case gamesocketappservice.PingRequestDtoType:
				continue
			case gamesocketappservice.MoveRequestDtoType:
				requestDto, err := json.Unmarshal[gamesocketappservice.MoveRequestDto](message)
				if err != nil {
					return
				}

				controller.gameAppService.MovePlayer(socketPresenter, gamesocketappservice.MovePlayerCommand{
					WorldId:   worldIdDto,
					PlayerId:  playerIdDto,
					Direction: requestDto.Direction,
				})
			case gamesocketappservice.PlaceItemRequestDtoType:
				requestDto, err := json.Unmarshal[gamesocketappservice.PlaceItemRequestDto](message)
				if err != nil {
					return
				}

				controller.gameAppService.PlaceItem(socketPresenter, gamesocketappservice.PlaceItemCommand{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
					ItemId:   requestDto.ItemId,
				})
			case gamesocketappservice.DestroyItemRequestDtoType:
				controller.gameAppService.DestroyItem(socketPresenter, gamesocketappservice.DestroyItemCommand{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
				})
			default:
			}
		}
	}()

	for {
		<-closeConnFlag

		controller.gameAppService.RemovePlayer(socketPresenter, gamesocketappservice.RemovePlayerCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		})
		return
	}
}
