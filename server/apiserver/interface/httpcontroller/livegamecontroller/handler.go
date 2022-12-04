package livegamecontroller

import (
	"encoding/json"
	"net/http"
	"sync"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/port/adapter/notification/redis"
	commonapplicationevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
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

type Configuration struct {
	GameApplicationService     service.GameApplicationService
	LiveGameApplicationService service.LiveGameApplicationService
}

func NewController(configuration Configuration) func(c *gin.Context) {
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
			func(event commonapplicationevent.GameInfoUpdatedApplicationEvent) {
				dimension, err := event.GetDimension()
				if err != nil {
					return
				}
				sendJSONMessageToClient(conn, socketConnLock, presenter.PresentInformationUpdatedEvent(dimension))
			},
		)
		defer redisGameInfoUpdatedSubscriberUnsubscriber()

		redisAreaZoomedSubscriber, _ := redis.NewRedisAreaZoomedSubscriber(liveGameId, playerId)
		redisAreaZoomedSubscriberUnsubscriber := redisAreaZoomedSubscriber.Subscribe(
			func(event commonapplicationevent.AreaZoomedApplicationEvent) {
				area, err := event.GetArea()
				if err != nil {
					return
				}
				unitBlock, err := event.GetUnitBlock()
				if err != nil {
					return
				}
				sendJSONMessageToClient(conn, socketConnLock, presenter.PresentAreaZoomedEvent(area, unitBlock))
			},
		)
		defer redisAreaZoomedSubscriberUnsubscriber()

		redisZoomedAreaUpdatedSubscriber, _ := redis.NewRedisZoomedAreaUpdatedSubscriber(liveGameId, playerId)
		redisZoomedAreaUpdatedSubscriberUnsubscriber := redisZoomedAreaUpdatedSubscriber.Subscribe(
			func(event commonapplicationevent.ZoomedAreaUpdatedApplicationEvent) {
				area, err := event.GetArea()
				if err != nil {
					return
				}
				unitBlock, err := event.GetUnitBlock()
				if err != nil {
					return
				}
				sendJSONMessageToClient(conn, socketConnLock, presenter.PresentZoomedAreaUpdatedEvent(area, unitBlock))
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
				case ZoomAreaRequestedEventType:
					area, err := presenter.ParseZoomAreaRequestedEvent(message)
					if err != nil {
						sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
						return
					}

					configuration.LiveGameApplicationService.RequestToZoomArea(liveGameId, playerId, area)
				case ReviveUnitsRequestedEventType:
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
