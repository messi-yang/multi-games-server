package livegamecontroller

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/areazoomedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/gameinfoupdatedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/zoomedareaupdatedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/messaging/redisintgreventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
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
		socketConn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.Error(err)
			return
		}
		defer socketConn.Close()
		closeConnFlag := make(chan bool)

		games, err := GameRepo.GetAll()
		if err != nil {
			return
		}
		gameId := games[0].GetId()

		liveGameId := gameId.ToString()
		playerId := uuid.New().String()
		socketPresenter := newSocketPresenter(socketConn, &sync.RWMutex{})
		gzipCompressor := gzipprovider.New()

		integrationEventSubscriberUnsubscriber := redisintgreventsubscriber.New().Subscribe(
			intgrevent.CreateLiveGameClientChannel(liveGameId, playerId),
			func(message []byte) {
				integrationEvent, err := intgrevent.Parse(message)
				if err != nil {
					return
				}

				if integrationEvent.Name == zoomedareaupdatedintgrevent.EVENT_NAME {
					event := zoomedareaupdatedintgrevent.Deserialize(message)
					liveGameAppService.SendZoomedAreaUpdatedEvent(socketPresenter, event.Area, event.UnitBlock)
				} else if integrationEvent.Name == areazoomedintgrevent.EVENT_NAME {
					event := areazoomedintgrevent.Deserialize(message)
					liveGameAppService.SendAreaZoomedEvent(socketPresenter, event.Area, event.UnitBlock)
				} else if integrationEvent.Name == gameinfoupdatedintgrevent.EVENT_NAME {
					event := gameinfoupdatedintgrevent.Deserialize(message)
					liveGameAppService.SendInformationUpdatedEvent(socketPresenter, event.Dimension)
				}
			})
		defer integrationEventSubscriberUnsubscriber()

		liveGameAppService.RequestToAddPlayer(liveGameId, playerId)

		go func() {
			defer func() {
				closeConnFlag <- true
			}()

			for {
				_, compressedMessage, err := socketConn.ReadMessage()
				if err != nil {
					liveGameAppService.SendErroredEvent(socketPresenter, err.Error())
					return
				}

				message, err := gzipCompressor.Ungzip(compressedMessage)
				if err != nil {
					liveGameAppService.SendErroredEvent(socketPresenter, err.Error())
					continue
				}

				commandType, err := livegameappservice.ParseCommandType(message)
				if err != nil {
					liveGameAppService.SendErroredEvent(socketPresenter, err.Error())
					continue
				}

				switch commandType {
				case livegameappservice.ZoomAreaCommanType:
					command, err := livegameappservice.ParseCommand[livegameappservice.ZoomAreaCommand](message)
					if err != nil {
						liveGameAppService.SendErroredEvent(socketPresenter, err.Error())
						continue
					}

					liveGameAppService.RequestToZoomArea(liveGameId, playerId, command.Payload.Area)
				case livegameappservice.BuildItemCommanType:
					command, err := livegameappservice.ParseCommand[livegameappservice.BuildItemCommand](message)
					if err != nil {
						liveGameAppService.SendErroredEvent(socketPresenter, err.Error())
						continue
					}

					liveGameAppService.RequestToBuildItem(liveGameId, command.Payload.Coordinate, command.Payload.ItemId)
				case livegameappservice.DestroyItemCommanType:
					command, err := livegameappservice.ParseCommand[livegameappservice.DestroyItemCommand](message)
					if err != nil {
						liveGameAppService.SendErroredEvent(socketPresenter, err.Error())
						continue
					}

					liveGameAppService.RequestToDestroyItem(liveGameId, command.Payload.Coordinate)
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
