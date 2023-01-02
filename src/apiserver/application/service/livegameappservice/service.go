package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/addplayerrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/buliditemrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/destroyitemrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/removeplayerrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/zoomarearequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/messaging/intgreventpublisher"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/areaviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/coordinateviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/dimensionviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/itemviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/unitblockviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type Service interface {
	QueryItems(presenter Presenter)
	SendErroredEventToClient(presenter Presenter, clientMessage string)
	SendInformationUpdatedEventToClient(presenter Presenter, rawDimension dimensionviewmodel.ViewModel)
	SendZoomedAreaUpdatedEventToClient(presenter Presenter, area areaviewmodel.ViewModel, unitBlock unitblockviewmodel.ViewModel)
	SendAreaZoomedEventToClient(presenter Presenter, area areaviewmodel.ViewModel, unitBlock unitblockviewmodel.ViewModel)
	RequestToAddPlayer(rawLiveGameId string, rawPlayerId string)
	RequestToZoomArea(rawLiveGameId string, rawPlayerId string, rawArea areaviewmodel.ViewModel)
	RequestToBuildItem(rawLiveGameId string, rawCoordinate coordinateviewmodel.ViewModel, rawItemId string)
	RequestToDestroyItem(rawLiveGameId string, rawCoordinate coordinateviewmodel.ViewModel)
	RequestToRemovePlayer(rawLiveGameId string, rawPlayerId string)
}

type serve struct {
	intgrEventPublisher intgreventpublisher.Publisher
	itemRepo            itemmodel.Repo
}

func New(intgrEventPublisher intgreventpublisher.Publisher, itemRepo itemmodel.Repo) Service {
	return &serve{intgrEventPublisher: intgrEventPublisher, itemRepo: itemRepo}
}

func (serve *serve) SendErroredEventToClient(presenter Presenter, clientMessage string) {
	event := ErroredEvent{}
	event.Type = ErrorHappenedEventType
	event.Payload.ClientMessage = clientMessage
	presenter.OnSuccess(event)
}

func (serve *serve) SendInformationUpdatedEventToClient(presenter Presenter, rawDimension dimensionviewmodel.ViewModel) {
	event := InformationUpdatedEvent{}
	event.Type = InformationUpdatedEventType
	event.Payload.Dimension = rawDimension
	presenter.OnSuccess(event)
}

func (serve *serve) QueryItems(presenter Presenter) {
	items := serve.itemRepo.GetAllItems()
	itemViewModels := itemviewmodel.BatchNew(items)

	event := ItemsUpdatedEvent{}
	event.Type = ItemsUpdatedEventType
	event.Payload.Items = itemViewModels
	presenter.OnSuccess(event)
}

func (serve *serve) SendZoomedAreaUpdatedEventToClient(presenter Presenter, area areaviewmodel.ViewModel, unitBlock unitblockviewmodel.ViewModel) {
	event := ZoomedAreaUpdatedEvent{}
	event.Type = ZoomedAreaUpdatedEventType
	event.Payload.Area = area
	event.Payload.UnitBlock = unitBlock
	presenter.OnSuccess(event)
}

func (serve *serve) SendAreaZoomedEventToClient(presenter Presenter, area areaviewmodel.ViewModel, unitBlock unitblockviewmodel.ViewModel) {
	event := AreaZoomedEvent{}
	event.Type = AreaZoomedEventType
	event.Payload.Area = area
	event.Payload.UnitBlock = unitBlock
	presenter.OnSuccess(event)
}

func (serve *serve) RequestToAddPlayer(rawLiveGameId string, rawPlayerId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		addplayerrequestedintgrevent.New(rawLiveGameId, rawPlayerId).Serialize(),
	)
}

func (serve *serve) RequestToZoomArea(rawLiveGameId string, rawPlayerId string, rawArea areaviewmodel.ViewModel) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		zoomarearequestedintgrevent.New(rawLiveGameId, rawPlayerId, rawArea).Serialize(),
	)
}

func (serve *serve) RequestToBuildItem(rawLiveGameId string, rawCoordinate coordinateviewmodel.ViewModel, rawItemId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		buliditemrequestedintgrevent.New(rawLiveGameId, rawCoordinate, rawItemId).Serialize(),
	)
}

func (serve *serve) RequestToDestroyItem(rawLiveGameId string, rawCoordinate coordinateviewmodel.ViewModel) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		destroyitemrequestedintgrevent.New(rawLiveGameId, rawCoordinate).Serialize(),
	)
}

func (serve *serve) RequestToRemovePlayer(rawLiveGameId string, rawPlayerId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		removeplayerrequestedintgrevent.New(rawLiveGameId, rawPlayerId).Serialize(),
	)
}
