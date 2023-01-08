package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/messaging/intgreventpublisher"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/samber/lo"
)

type Service interface {
	QueryItems(presenter Presenter)
	SendErroredServerEvent(presenter Presenter, clientMessage string)
	SendInformationUpdatedServerEvent(presenter Presenter, rawMapSize viewmodel.MapSizeViewModel)
	SendObservedMapRangeUpdatedServerEvent(presenter Presenter, mapRange viewmodel.MapRangeViewModel, unitMap viewmodel.UnitMapViewModel)
	SendMapRangeObservedServerEvent(presenter Presenter, mapRange viewmodel.MapRangeViewModel, unitMap viewmodel.UnitMapViewModel)
	RequestToAddPlayer(rawLiveGameId string, rawPlayerId string)
	RequestToObserveMapRange(rawLiveGameId string, rawPlayerId string, rawMapRange viewmodel.MapRangeViewModel)
	RequestToBuildItem(rawLiveGameId string, rawLocation viewmodel.LocationViewModel, rawItemId string)
	RequestToDestroyItem(rawLiveGameId string, rawLocation viewmodel.LocationViewModel)
	RequestToRemovePlayer(rawLiveGameId string, rawPlayerId string)
}

type serve struct {
	intgrEventPublisher intgreventpublisher.Publisher
	itemRepo            itemmodel.Repo
}

func New(intgrEventPublisher intgreventpublisher.Publisher, itemRepo itemmodel.Repo) Service {
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

func (serve *serve) SendObservedMapRangeUpdatedServerEvent(presenter Presenter, mapRange viewmodel.MapRangeViewModel, unitMap viewmodel.UnitMapViewModel) {
	event := ObservedMapRangeUpdatedServerEvent{}
	event.Type = ObservedMapRangeUpdatedServerEventType
	event.Payload.MapRange = mapRange
	event.Payload.UnitMap = unitMap
	presenter.OnSuccess(event)
}

func (serve *serve) SendMapRangeObservedServerEvent(presenter Presenter, mapRange viewmodel.MapRangeViewModel, unitMap viewmodel.UnitMapViewModel) {
	event := MapRangeObservedServerEvent{}
	event.Type = MapRangeObservedServerEventType
	event.Payload.MapRange = mapRange
	event.Payload.UnitMap = unitMap
	presenter.OnSuccess(event)
}

func (serve *serve) RequestToAddPlayer(rawLiveGameId string, rawPlayerId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewAddPlayerRequestedIntgrEvent(rawLiveGameId, rawPlayerId)),
	)
}

func (serve *serve) RequestToObserveMapRange(rawLiveGameId string, rawPlayerId string, rawMapRange viewmodel.MapRangeViewModel) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewObserveMapRangeRequestedEvent(rawLiveGameId, rawPlayerId, rawMapRange)),
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
