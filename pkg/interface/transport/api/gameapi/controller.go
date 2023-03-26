package gameapi

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gameapiservice"
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

	playersUpdatedIntEventUnsubscriber := redisinteventsubscriber.New[gameapiservice.PlayersUpdatedIntEvent]().Subscribe(
		gameapiservice.NewPlayersUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gameapiservice.PlayersUpdatedIntEvent) {
			err = gameAppService.GetPlayersAroundPlayer(gameapiservice.GetPlayersQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			})
			if err != nil {
				disconnect()
			}
		},
	)
	defer playersUpdatedIntEventUnsubscriber()

	unitsUpdatedIntEventTypeUnsubscriber := redisinteventsubscriber.New[gameapiservice.UnitsUpdatedIntEvent]().Subscribe(
		gameapiservice.NewUnitsUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gameapiservice.UnitsUpdatedIntEvent) {
			err = gameAppService.GetUnitsVisibleByPlayer(gameapiservice.GetUnitsVisibleByPlayerQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			})
			if err != nil {
				disconnect()
			}
		},
	)
	defer unitsUpdatedIntEventTypeUnsubscriber()

	visionBoundUpdatedIntEventTypeUnsubscriber := redisinteventsubscriber.New[gameapiservice.VisionBoundUpdatedIntEvent]().Subscribe(
		gameapiservice.NewVisionBoundUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gameapiservice.VisionBoundUpdatedIntEvent) {
			err = gameAppService.GetUnitsVisibleByPlayer(gameapiservice.GetUnitsVisibleByPlayerQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			})
			if err != nil {
				disconnect()
			}
		},
	)
	defer visionBoundUpdatedIntEventTypeUnsubscriber()

	err = gameAppService.AddPlayer(gameapiservice.AddPlayerCommand{
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

			genericRequestDto, err := jsonutil.Unmarshal[gameapiservice.GenericRequestDto](message)
			if err != nil {
				return
			}

			switch genericRequestDto.Type {
			case gameapiservice.PingRequestDtoType:
				continue
			case gameapiservice.MoveRequestDtoType:
				requestDto, err := jsonutil.Unmarshal[gameapiservice.MoveRequestDto](message)
				if err != nil {
					return
				}

				err = gameAppService.MovePlayer(gameapiservice.MovePlayerCommand{
					WorldId:   worldIdDto,
					PlayerId:  playerIdDto,
					Direction: requestDto.Direction,
				})
				if err != nil {
					return
				}
			case gameapiservice.PlaceItemRequestDtoType:
				requestDto, err := jsonutil.Unmarshal[gameapiservice.PlaceItemRequestDto](message)
				if err != nil {
					return
				}

				err = gameAppService.PlaceItem(gameapiservice.PlaceItemCommand{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
					ItemId:   requestDto.ItemId,
				})
				if err != nil {
					return
				}
			case gameapiservice.DestroyItemRequestDtoType:
				err = gameAppService.DestroyItem(gameapiservice.DestroyItemCommand{
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

		_ = gameAppService.RemovePlayer(gameapiservice.RemovePlayerCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		})
		return
	}
}
