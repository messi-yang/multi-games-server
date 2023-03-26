package gameapi

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gamesocketappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/gziputil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/jsonutil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/messaging/redisinteventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/transport/api"
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
		return
	}
	defer socketConn.Close()

	closeConnFlag := make(chan bool)
	disconnect := func() {
		closeConnFlag <- true
	}

	worldIdDto := c.Request.URL.Query().Get("id")

	playerIdDto := uuid.New().String()

	socketPresenter := api.NewSocketPresenter(socketConn, &sync.RWMutex{})

	playersUpdatedIntEventUnsubscriber := redisinteventsubscriber.New[gamesocketappservice.PlayersUpdatedIntEvent](
		controller.redisClient,
	).Subscribe(
		gamesocketappservice.NewPlayersUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gamesocketappservice.PlayersUpdatedIntEvent) {
			err = controller.gameAppService.GetPlayersAroundPlayer(socketPresenter, gamesocketappservice.GetPlayersQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			})
			if err != nil {
				disconnect()
			}
		},
	)
	defer playersUpdatedIntEventUnsubscriber()

	unitsUpdatedIntEventTypeUnsubscriber := redisinteventsubscriber.New[gamesocketappservice.UnitsUpdatedIntEvent](
		controller.redisClient,
	).Subscribe(
		gamesocketappservice.NewUnitsUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gamesocketappservice.UnitsUpdatedIntEvent) {
			err = controller.gameAppService.GetUnitsVisibleByPlayer(socketPresenter, gamesocketappservice.GetUnitsVisibleByPlayerQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			})
			if err != nil {
				disconnect()
			}
		},
	)
	defer unitsUpdatedIntEventTypeUnsubscriber()

	visionBoundUpdatedIntEventTypeUnsubscriber := redisinteventsubscriber.New[gamesocketappservice.VisionBoundUpdatedIntEvent](
		controller.redisClient,
	).Subscribe(
		gamesocketappservice.NewVisionBoundUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gamesocketappservice.VisionBoundUpdatedIntEvent) {
			err = controller.gameAppService.GetUnitsVisibleByPlayer(socketPresenter, gamesocketappservice.GetUnitsVisibleByPlayerQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			})
			if err != nil {
				disconnect()
			}
		},
	)
	defer visionBoundUpdatedIntEventTypeUnsubscriber()

	err = controller.gameAppService.AddPlayer(socketPresenter, gamesocketappservice.AddPlayerCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	})
	if err != nil {
		return
	}

	go func() {
		defer disconnect()

		for {
			_, compressedMessage, err := socketConn.ReadMessage()
			if err != nil {
				return
			}

			message, err := gziputil.Ungzip(compressedMessage)
			if err != nil {
				return
			}

			genericRequestDto, err := jsonutil.Unmarshal[gamesocketappservice.GenericRequestDto](message)
			if err != nil {
				return
			}

			switch genericRequestDto.Type {
			case gamesocketappservice.PingRequestDtoType:
				continue
			case gamesocketappservice.MoveRequestDtoType:
				requestDto, err := jsonutil.Unmarshal[gamesocketappservice.MoveRequestDto](message)
				if err != nil {
					return
				}

				err = controller.gameAppService.MovePlayer(socketPresenter, gamesocketappservice.MovePlayerCommand{
					WorldId:   worldIdDto,
					PlayerId:  playerIdDto,
					Direction: requestDto.Direction,
				})
				if err != nil {
					return
				}
			case gamesocketappservice.PlaceItemRequestDtoType:
				requestDto, err := jsonutil.Unmarshal[gamesocketappservice.PlaceItemRequestDto](message)
				if err != nil {
					return
				}

				err = controller.gameAppService.PlaceItem(socketPresenter, gamesocketappservice.PlaceItemCommand{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
					ItemId:   requestDto.ItemId,
				})
				if err != nil {
					return
				}
			case gamesocketappservice.DestroyItemRequestDtoType:
				err = controller.gameAppService.DestroyItem(socketPresenter, gamesocketappservice.DestroyItemCommand{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
				})
				if err != nil {
					return
				}
			default:
			}
		}
	}()

	for {
		<-closeConnFlag

		_ = controller.gameAppService.RemovePlayer(socketPresenter, gamesocketappservice.RemovePlayerCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		})
		return
	}
}
