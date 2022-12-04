package gamehandler

import (
	"encoding/json"
	"net/http"
	"sync"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/port/adapter/notification/redis"
	commonredisdto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis/dto"
	commonservice "github.com/dum-dum-genius/game-of-liberty-computer/server/common/service"
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

type HandlerConfiguration struct {
	GameApplicationService     service.GameApplicationService
	LiveGameApplicationService service.LiveGameApplicationService
}

func NewHandler(configuration HandlerConfiguration) func(c *gin.Context) {
	return func(c *gin.Context) {
		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.Error(err)
			return
		}
		defer conn.Close()
		closeConnFlag := make(chan bool)

		gameId, err := configuration.GameApplicationService.GetFirstGameId()
		if err != nil {
			return
		}

		liveGameId := livegamemodel.NewLiveGameId(gameId.GetId())
		playerId := gamecommonmodel.NewPlayerId(uuid.New())
		socketConnLock := &sync.RWMutex{}

		redisGameInfoUpdatedSubscriber, _ := redis.NewRedisGameInfoUpdatedSubscriber(liveGameId, playerId)
		redisGameInfoUpdatedSubscriberUnsubscriber := redisGameInfoUpdatedSubscriber.Subscribe(
			func(event commonredisdto.RedisGameInfoUpdatedEvent) {
				sendJSONMessageToClient(conn, socketConnLock, presenter.PresentInformationUpdatedEvent(event.Dimension))
			},
		)
		defer redisGameInfoUpdatedSubscriberUnsubscriber()

		redisAreaZoomedSubscriber, _ := redis.NewRedisAreaZoomedSubscriber(liveGameId, playerId)
		redisAreaZoomedSubscriberUnsubscriber := redisAreaZoomedSubscriber.Subscribe(
			func(event commonredisdto.RedisAreaZoomedEvent) {
				sendJSONMessageToClient(conn, socketConnLock, presenter.PresentAreaZoomedEvent(event.Area, event.UnitBlock))
			},
		)
		defer redisAreaZoomedSubscriberUnsubscriber()

		redisZoomedAreaUpdatedSubscriber, _ := redis.NewRedisZoomedAreaUpdatedSubscriber(liveGameId, playerId)
		redisZoomedAreaUpdatedSubscriberUnsubscriber := redisZoomedAreaUpdatedSubscriber.Subscribe(
			func(event commonredisdto.RedisZoomedAreaUpdatedEvent) {
				sendJSONMessageToClient(conn, socketConnLock, presenter.PresentZoomedAreaUpdatedEvent(event.Area, event.UnitBlock))
			},
		)
		defer redisZoomedAreaUpdatedSubscriberUnsubscriber()

		configuration.LiveGameApplicationService.RequestToAddPlayer(liveGameId, playerId)

		go func() {
			defer func() {
				closeConnFlag <- true
			}()

			for {
				_, compressedMessage, err := conn.ReadMessage()
				if err != nil {
					sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
					break
				}

				gzipCompressor := commonservice.NewGzipService()
				message, err := gzipCompressor.Ungzip(compressedMessage)
				if err != nil {
					sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
					break
				}

				eventType, err := presenter.ParseEventType(message)
				if err != nil {
					sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
					break
				}

				switch eventType {
				case RedisZoomAreaRequestedEventType:
					area, err := presenter.ParseZoomAreaRequestedEvent(message)
					if err != nil {
						sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
						return
					}

					configuration.LiveGameApplicationService.RequestToZoomArea(liveGameId, playerId, area)
				case RedisReviveUnitsRequestedEventType:
					coordinates, err := presenter.ParseReviveUnitsRequestedEvent(message)
					if err != nil {
						sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
						return
					}

					configuration.LiveGameApplicationService.RequestToReviveUnits(liveGameId, coordinates)
				default:
				}
			}
		}()

		for {
			<-closeConnFlag

			configuration.LiveGameApplicationService.RequestToRemovePlayer(liveGameId, playerId)
			return
		}
	}
}

func sendJSONMessageToClient(conn *websocket.Conn, socketConnLock *sync.RWMutex, message any) {
	socketConnLock.Lock()
	defer socketConnLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(message)

	gzipCompressor := commonservice.NewGzipService()
	compressedMessage, _ := gzipCompressor.Gzip(messageJsonInBytes)

	conn.WriteMessage(2, compressedMessage)
}
