package livegamecontroller

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/areazoomedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/gameinfoupdatedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/zoomedareaupdatedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/areaviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/coordinateviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/messaging/redisintgreventsubscriber"
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

		liveGameId, _ := livegamemodel.NewLiveGameId(gameId.ToString())
		playerId, _ := commonmodel.NewPlayerId(uuid.New().String())
		socketConnLock := &sync.RWMutex{}

		integrationEventSubscriberUnsubscriber := redisintgreventsubscriber.New().Subscribe(
			intgrevent.CreateLiveGameClientChannel(liveGameId, playerId),
			func(message []byte) {
				integrationEvent := intgrevent.New(message)

				if integrationEvent.Name == zoomedareaupdatedintgrevent.EVENT_NAME {
					event := zoomedareaupdatedintgrevent.Deserialize(message)
					sendJSONMessageToClient(conn, socketConnLock, presenter.PresentZoomedAreaUpdatedEvent(event.Area, event.UnitBlock))
				} else if integrationEvent.Name == areazoomedintgrevent.EVENT_NAME {
					event := areazoomedintgrevent.Deserialize(message)
					sendJSONMessageToClient(conn, socketConnLock, presenter.PresentAreaZoomedEvent(event.Area, event.UnitBlock))
				} else if integrationEvent.Name == gameinfoupdatedintgrevent.EVENT_NAME {
					event := gameinfoupdatedintgrevent.Deserialize(message)
					sendJSONMessageToClient(conn, socketConnLock, presenter.PresentInformationUpdatedEvent(event.Dimension))
				}
			})
		defer integrationEventSubscriberUnsubscriber()

		liveGameAppService.RequestToAddPlayer(liveGameId.ToString(), playerId.ToString())

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

					liveGameAppService.RequestToZoomArea(liveGameId.ToString(), playerId.ToString(), areaviewmodel.New(area))
				case BuildItemEventType:
					coordinate, itemId, err := presenter.ParseBuildItemEvent(message)
					if err != nil {
						sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
						return
					}

					liveGameAppService.RequestToBuildItem(liveGameId.ToString(), coordinateviewmodel.New(coordinate), itemId.ToString())
				case DestroyItemEventType:
					coordinate, err := presenter.ParseDestroyItemEvent(message)
					if err != nil {
						sendJSONMessageToClient(conn, socketConnLock, presenter.PresentErroredEvent(err.Error()))
						return
					}

					liveGameAppService.RequestToDestroyItem(liveGameId.ToString(), coordinateviewmodel.New(coordinate))
				default:
				}
			}
		}()

		for {
			<-closeConnFlag

			liveGameAppService.RequestToRemovePlayer(liveGameId.ToString(), playerId.ToString())
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
