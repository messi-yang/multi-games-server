package gamehandler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/rediseventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/adapter/applicationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/mainserver/application/applicationservice"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/ztrue/tracerr"
)

type clientSession struct {
	liveGameId            uuid.UUID
	playerId              uuid.UUID
	socketSendMessageLock sync.RWMutex
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type HandlerConfiguration struct {
	LiveGameApplicationService *applicationservice.LiveGameApplicationService
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

		gamesIds := configuration.LiveGameApplicationService.GetAllLiveGameIds()
		liveGameId := gamesIds[0]

		clientSession := &clientSession{
			liveGameId:            liveGameId.GetId(),
			playerId:              uuid.New(),
			socketSendMessageLock: sync.RWMutex{},
		}

		gameInfoUpdatedApplicationEventSubscriber := rediseventbus.NewRedisApplicationEventBus(
			rediseventbus.WithRedisInfrastructureService[applicationevent.GameInfoUpdatedApplicationEvent](),
		).Subscribe(
			applicationevent.NewGameInfoUpdatedApplicationEventTopic(clientSession.liveGameId, clientSession.playerId),
			func(event applicationevent.GameInfoUpdatedApplicationEvent) {
				handleGameInfoUpdatedEvent(conn, clientSession, event)
			},
		)
		defer gameInfoUpdatedApplicationEventSubscriber()

		areaZoomedApplicationEventUnsubscriber := rediseventbus.NewRedisApplicationEventBus(
			rediseventbus.WithRedisInfrastructureService[applicationevent.AreaZoomedApplicationEvent](),
		).Subscribe(
			applicationevent.NewAreaZoomedApplicationEventTopic(clientSession.liveGameId, clientSession.playerId),
			func(event applicationevent.AreaZoomedApplicationEvent) {
				handleAreaZoomedEvent(conn, clientSession, event)
			},
		)
		defer areaZoomedApplicationEventUnsubscriber()

		zoomedAreaUpdatedApplicationEventUnsubscriber := rediseventbus.NewRedisApplicationEventBus(
			rediseventbus.WithRedisInfrastructureService[applicationevent.ZoomedAreaUpdatedApplicationEvent](),
		).Subscribe(
			applicationevent.NewZoomedAreaUpdatedApplicationEventTopic(clientSession.liveGameId, clientSession.playerId),
			func(event applicationevent.ZoomedAreaUpdatedApplicationEvent) {
				handleZoomedAreaUpdatedEvent(conn, clientSession, event)
			},
		)
		defer zoomedAreaUpdatedApplicationEventUnsubscriber()

		rediseventbus.NewRedisApplicationEventBus(
			rediseventbus.WithRedisInfrastructureService[applicationevent.AddPlayerRequestedApplicationEvent](),
		).Publish(
			applicationevent.NewAddPlayerRequestedApplicationEventTopic(clientSession.liveGameId),
			applicationevent.NewAddPlayerRequestedApplicationEvent(clientSession.playerId),
		)

		go func() {
			defer func() {
				closeConnFlag <- true
			}()

			for {
				_, compressedMessage, err := conn.ReadMessage()
				if err != nil {
					emitErrorEvent(conn, clientSession, err)
					break
				}

				compressionApplicationService := service.NewCompressionApplicationService()
				message, err := compressionApplicationService.Ungzip(compressedMessage)
				if err != nil {
					emitErrorEvent(conn, clientSession, err)
					break
				}

				eventType, err := gameHandlerPresenter.ExtractEventType(message)
				if err != nil {
					emitErrorEvent(conn, clientSession, err)
					break
				}

				switch eventType {
				case ZoomAreaRequestedEventType:
					handleZoomAreaRequestedEvent(conn, clientSession, message)
				case ReviveUnitsRequestedEventType:
					handleReviveUnitsRequestedEvent(conn, clientSession, message)
				default:
				}
			}
		}()

		for {
			<-closeConnFlag

			rediseventbus.NewRedisApplicationEventBus(
				rediseventbus.WithRedisInfrastructureService[applicationevent.RemovePlayerRequestedApplicationEvent](),
			).Publish(
				applicationevent.NewRemovePlayerRequestedApplicationEventTopic(clientSession.liveGameId),
				applicationevent.NewRemovePlayerRequestedApplicationEvent(clientSession.playerId),
			)

			return
		}
	}
}

func sendJSONMessageToClient(conn *websocket.Conn, clientSession *clientSession, message any) {
	clientSession.socketSendMessageLock.Lock()
	defer clientSession.socketSendMessageLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(message)

	compressionApplicationService := service.NewCompressionApplicationService()
	compressedMessage, err := compressionApplicationService.Gzip(messageJsonInBytes)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	conn.WriteMessage(2, compressedMessage)
}

func emitErrorEvent(conn *websocket.Conn, clientSession *clientSession, err error) {
	tracerr.Print(tracerr.Wrap(err))
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateErroredEvent(err.Error()))
}

func handleGameInfoUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event applicationevent.GameInfoUpdatedApplicationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateInformationUpdatedEvent(event.Dimension))
}

func handleAreaZoomedEvent(conn *websocket.Conn, clientSession *clientSession, event applicationevent.AreaZoomedApplicationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateAreaZoomedEvent(event.Area, event.UnitBlock))
}

func handleZoomedAreaUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event applicationevent.ZoomedAreaUpdatedApplicationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateZoomedAreaUpdatedEvent(event.Area, event.UnitBlock))
}

func handleZoomAreaRequestedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	areaDto, err := gameHandlerPresenter.ExtractZoomAreaRequestedEvent(event)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.ZoomAreaRequestedApplicationEvent](),
	).Publish(
		applicationevent.NewZoomAreaRequestedApplicationEventTopic(clientSession.liveGameId),
		applicationevent.NewZoomAreaRequestedApplicationEvent(clientSession.playerId, areaDto),
	)
}

func handleReviveUnitsRequestedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	coordinateDtos, err := gameHandlerPresenter.ExtractReviveUnitsRequestedEvent(event)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.ReviveUnitsRequestedApplicationEvent](),
	).Publish(
		applicationevent.NewReviveUnitsRequestedApplicationEventTopic(clientSession.liveGameId),
		applicationevent.NewReviveUnitsRequestedApplicationEvent(coordinateDtos),
	)
}
