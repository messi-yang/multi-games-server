package livegamecontroller

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/adapter/notification/redis"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/application/appservice"
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/service"
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

func NewController(
	GameRepository gamemodel.GameRepository,
	LiveGameAppService appservice.LiveGameAppService,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.Error(err)
			return
		}
		defer conn.Close()
		closeConnFlag := make(chan bool)

		games, err := GameRepository.GetAll()
		if err != nil {
			return
		}
		gameId := games[0].GetId()

		liveGameId, _ := livegamemodel.NewLiveGameId(gameId.GetId().String())
		playerId, _ := commonmodel.NewPlayerId(uuid.New().String())
		socketConnLock := &sync.RWMutex{}

		redisGameInfoUpdatedSubscriber, _ := redis.NewRedisGameInfoUpdatedSubscriber(liveGameId, playerId)
		redisGameInfoUpdatedSubscriberUnsubscriber := redisGameInfoUpdatedSubscriber.Subscribe(
			func(event *commonappevent.GameInfoUpdatedAppEvent) {
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
			func(event *commonappevent.AreaZoomedAppEvent) {
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
			func(event *commonappevent.ZoomedAreaUpdatedAppEvent) {
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

		LiveGameAppService.RequestToAddPlayer(liveGameId, playerId)

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

				gzipCompressor := service.NewGzipService()
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
				case ZoomAreaEventType:
					area, err := presenter.ParseZoomAreaEvent(message)
					if err != nil {
						sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
						return
					}

					LiveGameAppService.RequestToZoomArea(liveGameId, playerId, area)
				case BuildItemEventType:
					coordinate, itemId, err := presenter.ParseBuildItemEvent(message)
					if err != nil {
						sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
						return
					}

					LiveGameAppService.RequestToBuildItem(liveGameId, coordinate, itemId)
				case DestroyItemEventType:
					coordinate, err := presenter.ParseDestroyItemEvent(message)
					if err != nil {
						sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
						return
					}

					LiveGameAppService.RequestToDestroyItem(liveGameId, coordinate)
				default:
				}
			}
		}()

		for {
			<-closeConnFlag

			LiveGameAppService.RequestToRemovePlayer(liveGameId, playerId)
			return
		}
	}
}

func sendJSONMessageToClient(conn *websocket.Conn, socketConnLock *sync.RWMutex, message any) {
	socketConnLock.Lock()
	defer socketConnLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(message)

	gzipCompressor := service.NewGzipService()
	compressedMessage, _ := gzipCompressor.Gzip(messageJsonInBytes)

	conn.WriteMessage(2, compressedMessage)
}
