package livegamecontroller

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/gameinfoupdatedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/maprangeobservedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/observedmaprangeupdatedintgrevent"
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

	integrationEventSubscriberUnsubscriber := redisintgreventsubscriber.New().Subscribe(
		intgrevent.CreateLiveGameClientChannel(liveGameId, playerId),
		func(message []byte) {
			integrationEvent, err := intgrevent.Parse(message)
			if err != nil {
				return
			}

			switch integrationEvent.Name {
			case observedmaprangeupdatedintgrevent.EVENT_NAME:
				event := observedmaprangeupdatedintgrevent.Deserialize(message)
				controller.liveGameAppService.SendObservedMapRangeUpdatedEventToClient(socketPresenter, event.MapRange, event.GameMap)
			case maprangeobservedintgrevent.EVENT_NAME:
				event := maprangeobservedintgrevent.Deserialize(message)
				controller.liveGameAppService.SendMapRangeObservedEventToClient(socketPresenter, event.MapRange, event.GameMap)
			case gameinfoupdatedintgrevent.EVENT_NAME:
				event := gameinfoupdatedintgrevent.Deserialize(message)
				controller.liveGameAppService.SendInformationUpdatedEventToClient(socketPresenter, event.MapSize)
			}

		})
	defer integrationEventSubscriberUnsubscriber()

	controller.liveGameAppService.RequestToAddPlayer(liveGameId, playerId)

	go func() {
		defer func() {
			closeConnFlag <- true
		}()

		for {
			_, compressedMessage, err := socketConn.ReadMessage()
			if err != nil {
				controller.liveGameAppService.SendErroredEventToClient(socketPresenter, err.Error())
				return
			}

			message, err := gzipCompressor.Ungzip(compressedMessage)
			if err != nil {
				controller.liveGameAppService.SendErroredEventToClient(socketPresenter, err.Error())
				continue
			}

			commandType, err := livegameappservice.ParseCommandType(message)
			if err != nil {
				controller.liveGameAppService.SendErroredEventToClient(socketPresenter, err.Error())
				continue
			}

			switch commandType {
			case livegameappservice.ObserveMapRangeCommanType:
				command, err := livegameappservice.ParseCommand[livegameappservice.ObserveMapRangeCommand](message)
				if err != nil {
					controller.liveGameAppService.SendErroredEventToClient(socketPresenter, err.Error())
					continue
				}

				controller.liveGameAppService.RequestToObserveMapRange(liveGameId, playerId, command.Payload.MapRange)
			case livegameappservice.BuildItemCommanType:
				command, err := livegameappservice.ParseCommand[livegameappservice.BuildItemCommand](message)
				if err != nil {
					controller.liveGameAppService.SendErroredEventToClient(socketPresenter, err.Error())
					continue
				}

				controller.liveGameAppService.RequestToBuildItem(liveGameId, command.Payload.Location, command.Payload.ItemId)
			case livegameappservice.DestroyItemCommanType:
				command, err := livegameappservice.ParseCommand[livegameappservice.DestroyItemCommand](message)
				if err != nil {
					controller.liveGameAppService.SendErroredEventToClient(socketPresenter, err.Error())
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
