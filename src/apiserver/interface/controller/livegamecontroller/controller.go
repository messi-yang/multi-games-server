package livegamecontroller

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
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

type Controller struct {
	gameRepo           gamemodel.GameRepo
	liveGameAppService livegameappservice.Service
}

func NewController(
	gameRepo gamemodel.GameRepo,
	liveGameAppService livegameappservice.Service,
) *Controller {
	return &Controller{
		gameRepo:           gameRepo,
		liveGameAppService: liveGameAppService,
	}
}

func (controller *Controller) HandleLiveGameConnection(c *gin.Context) {
	socketConn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Error(err)
		return
	}
	defer socketConn.Close()
	closeConnFlag := make(chan bool)

	games, err := controller.gameRepo.GetAll()
	if err != nil {
		return
	}
	gameId := games[0].GetId()

	liveGameId := gameId.ToString()
	playerId := uuid.New().String()
	socketPresenter := newSocketPresenter(socketConn, &sync.RWMutex{})
	gzipCompressor := gzipprovider.New()

	go controller.liveGameAppService.QueryItems(socketPresenter)

	intgrEventSubscriberUnsubscriber := redisintgreventsubscriber.New().Subscribe(
		intgrevent.CreateLiveGameClientChannel(liveGameId, playerId),
		func(message []byte) {
			intgrEvent, err := intgrevent.Unmarshal[intgrevent.GenericIntgrEvent](message)
			if err != nil {
				return
			}

			switch intgrEvent.Name {
			case intgrevent.ObservedExtentUpdatedEventName:
				event, err := intgrevent.Unmarshal[intgrevent.ObservedExtentUpdatedEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendObservedExtentUpdatedServerEvent(socketPresenter, event.Extent, event.UnitMap)
			case intgrevent.ExtentObservedEventName:
				event, err := intgrevent.Unmarshal[intgrevent.ExtentObservedEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendExtentObservedServerEvent(socketPresenter, event.Extent, event.UnitMap)
			case intgrevent.GameInfoUpdatedEventName:
				event, err := intgrevent.Unmarshal[intgrevent.GameInfoUpdatedEvent](message)
				if err != nil {
					return
				}
				controller.liveGameAppService.SendInformationUpdatedServerEvent(socketPresenter, event.MapSize)
			}

		})
	defer intgrEventSubscriberUnsubscriber()

	controller.liveGameAppService.RequestToAddPlayer(liveGameId, playerId)

	go func() {
		defer func() {
			closeConnFlag <- true
		}()

		for {
			_, compressedMessage, err := socketConn.ReadMessage()
			if err != nil {
				controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
				return
			}

			message, err := gzipCompressor.Ungzip(compressedMessage)
			if err != nil {
				controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
				continue
			}

			commandType, err := livegameappservice.ParseClientEventType(message)
			if err != nil {
				controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
				continue
			}

			switch commandType {
			case livegameappservice.PingClientEventType:
				continue
			case livegameappservice.ObserveExtentClientEventType:
				command, err := livegameappservice.ParseClientEvent[livegameappservice.ObserveExtentClientEvent](message)
				if err != nil {
					controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
					continue
				}

				controller.liveGameAppService.RequestToObserveExtent(liveGameId, playerId, command.Payload.Extent)
			case livegameappservice.BuildItemClientEventType:
				command, err := livegameappservice.ParseClientEvent[livegameappservice.BuildItemClientEvent](message)
				if err != nil {
					controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
					continue
				}

				controller.liveGameAppService.RequestToBuildItem(liveGameId, command.Payload.Location, command.Payload.ItemId)
			case livegameappservice.DestroyItemClientEventType:
				command, err := livegameappservice.ParseClientEvent[livegameappservice.DestroyItemClientEvent](message)
				if err != nil {
					controller.liveGameAppService.SendErroredServerEvent(socketPresenter, err.Error())
					continue
				}

				controller.liveGameAppService.RequestToDestroyItem(liveGameId, command.Payload.Location)
			default:
			}
		}
	}()

	for {
		<-closeConnFlag

		controller.liveGameAppService.RequestToRemovePlayer(liveGameId, playerId)
		return
	}
}
