package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/samber/lo"
)

type Service interface {
	QueryItems(presenter Presenter)
	SendErroredServerEvent(presenter Presenter, clientMessage string)
	SendInformationUpdatedServerEvent(presenter Presenter, rawMapSize viewmodel.MapSizeViewModel)
	SendObservedExtentUpdatedServerEvent(presenter Presenter, extent viewmodel.ExtentViewModel, unitMap viewmodel.UnitMapViewModel)
	SendExtentObservedServerEvent(presenter Presenter, extent viewmodel.ExtentViewModel, unitMap viewmodel.UnitMapViewModel)
	RequestToAddPlayer(rawLiveGameId string, rawPlayerId string)
	RequestToObserveExtent(rawLiveGameId string, rawPlayerId string, rawExtent viewmodel.ExtentViewModel)
	RequestToBuildItem(rawLiveGameId string, rawLocation viewmodel.LocationViewModel, rawItemId string)
	RequestToDestroyItem(rawLiveGameId string, rawLocation viewmodel.LocationViewModel)
	RequestToRemovePlayer(rawLiveGameId string, rawPlayerId string)
}

type serve struct {
	intgrEventPublisher intgrevent.IntgrEventPublisher
	itemRepo            itemmodel.Repo
}

func New(intgrEventPublisher intgrevent.IntgrEventPublisher, itemRepo itemmodel.Repo) Service {
	return &serve{intgrEventPublisher: intgrEventPublisher, itemRepo: itemRepo}
}

func (serve *serve) SendErroredServerEvent(presenter Presenter, clientMessage string) {
	event := ErroredServerEvent{}
	event.Type = ErroredServerEventType
	event.Payload.ClientMessage = clientMessage
	presenter.OnSuccess(event)
}

func (serve *serve) SendInformationUpdatedServerEvent(presenter Presenter, rawMapSize viewmodel.MapSizeViewModel) {
	event := InformationUpdatedServerEvent{}
	event.Type = InformationUpdatedServerEventType
	event.Payload.MapSize = rawMapSize
	presenter.OnSuccess(event)
}

func (serve *serve) QueryItems(presenter Presenter) {
	items := serve.itemRepo.GetAllItems()
	itemViewModels := lo.Map(items, func(item itemmodel.Item, _ int) viewmodel.ItemViewModel {
		return viewmodel.NewItemViewModel(item)
	})

	event := ItemsUpdatedServerEvent{}
	event.Type = ItemsUpdatedServerEventType
	event.Payload.Items = itemViewModels
	presenter.OnSuccess(event)
}

func (serve *serve) SendObservedExtentUpdatedServerEvent(presenter Presenter, extent viewmodel.ExtentViewModel, unitMap viewmodel.UnitMapViewModel) {
	event := ObservedExtentUpdatedServerEvent{}
	event.Type = ObservedExtentUpdatedServerEventType
	event.Payload.Extent = extent
	event.Payload.UnitMap = unitMap
	presenter.OnSuccess(event)
}

func (serve *serve) SendExtentObservedServerEvent(presenter Presenter, extent viewmodel.ExtentViewModel, unitMap viewmodel.UnitMapViewModel) {
	event := ExtentObservedServerEvent{}
	event.Type = ExtentObservedServerEventType
	event.Payload.Extent = extent
	event.Payload.UnitMap = unitMap
	presenter.OnSuccess(event)
}

func (serve *serve) RequestToAddPlayer(rawLiveGameId string, rawPlayerId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewAddPlayerRequestedIntgrEvent(rawLiveGameId, rawPlayerId)),
	)
}

func (serve *serve) RequestToObserveExtent(rawLiveGameId string, rawPlayerId string, rawExtent viewmodel.ExtentViewModel) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewObserveExtentRequestedEvent(rawLiveGameId, rawPlayerId, rawExtent)),
	)
}

func (serve *serve) RequestToBuildItem(rawLiveGameId string, rawLocation viewmodel.LocationViewModel, rawItemId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewBuildItemRequestedIntgrEvent(rawLiveGameId, rawLocation, rawItemId)),
	)
}

func (serve *serve) RequestToDestroyItem(rawLiveGameId string, rawLocation viewmodel.LocationViewModel) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewDestroyItemRequested(rawLiveGameId, rawLocation)),
	)
}

func (serve *serve) RequestToRemovePlayer(rawLiveGameId string, rawPlayerId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewRemovePlayerRequestedEvent(rawLiveGameId, rawPlayerId)),
	)
}
