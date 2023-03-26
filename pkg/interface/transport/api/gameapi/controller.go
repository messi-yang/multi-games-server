package gameapi

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/gziputil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/jsonutil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/messaging/redisinteventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/transport/api"
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

func gameConnectionHandler(c *gin.Context) {
	worldIdDto, err := uuid.Parse(c.Request.URL.Query().Get("id"))
	if err != nil {
		return
	}
	socketConn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer socketConn.Close()

	socketPresenter := api.NewSocketPresenter(socketConn, &sync.RWMutex{})
	gameAppService, err := newGameAppService(socketPresenter)
	if err != nil {
		return
	}

	closeConnFlag := make(chan bool)
	disconnect := func() {
		closeConnFlag <- true
	}

	playerIdDto := uuid.New()

	playersUpdatedIntEventUnsubscriber := redisinteventsubscriber.New[gameappservice.PlayersUpdatedIntEvent]().Subscribe(
		gameappservice.NewPlayersUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gameappservice.PlayersUpdatedIntEvent) {
			err = gameAppService.GetPlayersAroundPlayer(gameappservice.GetPlayersQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			})
			if err != nil {
				disconnect()
			}
		},
	)
	defer playersUpdatedIntEventUnsubscriber()

	unitsUpdatedIntEventTypeUnsubscriber := redisinteventsubscriber.New[gameappservice.UnitsUpdatedIntEvent]().Subscribe(
		gameappservice.NewUnitsUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gameappservice.UnitsUpdatedIntEvent) {
			err = gameAppService.GetUnitsVisibleByPlayer(gameappservice.GetUnitsVisibleByPlayerQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			})
			if err != nil {
				disconnect()
			}
		},
	)
	defer unitsUpdatedIntEventTypeUnsubscriber()

	visionBoundUpdatedIntEventTypeUnsubscriber := redisinteventsubscriber.New[gameappservice.VisionBoundUpdatedIntEvent]().Subscribe(
		gameappservice.NewVisionBoundUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gameappservice.VisionBoundUpdatedIntEvent) {
			err = gameAppService.GetUnitsVisibleByPlayer(gameappservice.GetUnitsVisibleByPlayerQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			})
			if err != nil {
				disconnect()
			}
		},
	)
	defer visionBoundUpdatedIntEventTypeUnsubscriber()

	err = gameAppService.AddPlayer(gameappservice.AddPlayerCommand{
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

			genericRequestDto, err := jsonutil.Unmarshal[gameappservice.GenericRequestDto](message)
			if err != nil {
				return
			}

			switch genericRequestDto.Type {
			case gameappservice.PingRequestDtoType:
				continue
			case gameappservice.MoveRequestDtoType:
				requestDto, err := jsonutil.Unmarshal[gameappservice.MoveRequestDto](message)
				if err != nil {
					return
				}

				err = gameAppService.MovePlayer(gameappservice.MovePlayerCommand{
					WorldId:   worldIdDto,
					PlayerId:  playerIdDto,
					Direction: requestDto.Direction,
				})
				if err != nil {
					return
				}
			case gameappservice.PlaceItemRequestDtoType:
				requestDto, err := jsonutil.Unmarshal[gameappservice.PlaceItemRequestDto](message)
				if err != nil {
					return
				}

				err = gameAppService.PlaceItem(gameappservice.PlaceItemCommand{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
					ItemId:   requestDto.ItemId,
				})
				if err != nil {
					return
				}
			case gameappservice.DestroyItemRequestDtoType:
				err = gameAppService.DestroyItem(gameappservice.DestroyItemCommand{
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

		_ = gameAppService.RemovePlayer(gameappservice.RemovePlayerCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		})
		return
	}
}
