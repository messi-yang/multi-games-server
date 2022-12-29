package livegamecontroller

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/areazoomedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/gameinfoupdatedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/zoomedareaupdatedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/integrationevent/redisintegrationeventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/gzipprovider"
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
	GameRepo gamemodel.GameRepo,
	liveGameAppService livegameappservice.Service,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.Error(err)
			return
		}
		defer conn.Close()
		closeConnFlag := make(chan bool)

		games, err := GameRepo.GetAll()
		if err != nil {
			return
		}
		gameId := games[0].GetId()

		liveGameId, _ := livegamemodel.NewLiveGameId(gameId.GetId().String())
		playerId, _ := commonmodel.NewPlayerId(uuid.New().String())
		socketConnLock := &sync.RWMutex{}

		integrationEventSubscriber, _ := redisintegrationeventsubscriber.New()
		integrationEventSubscriberUnsubscriber := integrationEventSubscriber.Subscribe(integrationevent.CreateLiveGameClientChannel(liveGameId, playerId), func(message []byte) {
			integrationEvent := integrationevent.New(message)

			if integrationEvent.Name == zoomedareaupdatedintegrationevent.EVENT_NAME {
				event := zoomedareaupdatedintegrationevent.Deserialize(message)
				area, err := event.GetArea()
				if err != nil {
					return
				}
				unitBlock, err := event.GetUnitBlock()
				if err != nil {
					return
				}
				sendJSONMessageToClient(conn, socketConnLock, presenter.PresentZoomedAreaUpdatedEvent(area, unitBlock))
			} else if integrationEvent.Name == areazoomedintegrationevent.EVENT_NAME {
				event := areazoomedintegrationevent.Deserialize(message)
				area, err := event.GetArea()
				if err != nil {
					return
				}
				unitBlock, err := event.GetUnitBlock()
				if err != nil {
					return
				}
				sendJSONMessageToClient(conn, socketConnLock, presenter.PresentAreaZoomedEvent(area, unitBlock))
			} else if integrationEvent.Name == gameinfoupdatedintegrationevent.EVENT_NAME {
				event := gameinfoupdatedintegrationevent.Deserialize(message)
				dimension, err := event.GetDimension()
				if err != nil {
					return
				}
				sendJSONMessageToClient(conn, socketConnLock, presenter.PresentInformationUpdatedEvent(dimension))
			}
		})
		defer integrationEventSubscriberUnsubscriber()

		liveGameAppService.RequestToAddPlayer(liveGameId, playerId)

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

				gzipCompressor := gzipprovider.New()
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

					liveGameAppService.RequestToZoomArea(liveGameId, playerId, area)
				case BuildItemEventType:
					coordinate, itemId, err := presenter.ParseBuildItemEvent(message)
					if err != nil {
						sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
						return
					}

					liveGameAppService.RequestToBuildItem(liveGameId, coordinate, itemId)
				case DestroyItemEventType:
					coordinate, err := presenter.ParseDestroyItemEvent(message)
					if err != nil {
						sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
						return
					}

					liveGameAppService.RequestToDestroyItem(liveGameId, coordinate)
				default:
				}
			}
		}()

		for {
			<-closeConnFlag

			liveGameAppService.RequestToRemovePlayer(liveGameId, playerId)
			return
		}
	}
}

func sendJSONMessageToClient(conn *websocket.Conn, socketConnLock *sync.RWMutex, message any) {
	socketConnLock.Lock()
	defer socketConnLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(message)

	gzipCompressor := gzipprovider.New()
	compressedMessage, _ := gzipCompressor.Gzip(messageJsonInBytes)

	conn.WriteMessage(2, compressedMessage)
}
