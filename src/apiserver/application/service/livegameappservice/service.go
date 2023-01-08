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
	SendInformationUpdatedServerEvent(presenter Presenter, rawMapSize viewmodel.MapSize)
	SendObservedRangeUpdatedServerEvent(presenter Presenter, rangeVm viewmodel.Range, unitMap viewmodel.UnitMap)
	SendRangeObservedServerEvent(presenter Presenter, rangeVm viewmodel.Range, unitMap viewmodel.UnitMap)
	RequestToAddPlayer(rawLiveGameId string, rawPlayerId string)
	RequestToObserveRange(rawLiveGameId string, rawPlayerId string, rangeVm viewmodel.Range)
	RequestToBuildItem(rawLiveGameId string, rawLocation viewmodel.Location, rawItemId string)
	RequestToDestroyItem(rawLiveGameId string, rawLocation viewmodel.Location)
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

func (serve *serve) SendInformationUpdatedServerEvent(presenter Presenter, rawMapSize viewmodel.MapSize) {
	event := InformationUpdatedServerEvent{}
	event.Type = InformationUpdatedServerEventType
	event.Payload.MapSize = rawMapSize
	presenter.OnSuccess(event)
}

func (serve *serve) QueryItems(presenter Presenter) {
	items := serve.itemRepo.GetAllItems()
	itemViewModels := lo.Map(items, func(item itemmodel.Item, _ int) viewmodel.Item {
		return viewmodel.NewItem(item)
	})

	event := ItemsUpdatedServerEvent{}
	event.Type = ItemsUpdatedServerEventType
	event.Payload.Items = itemViewModels
	presenter.OnSuccess(event)
}

func (serve *serve) SendObservedRangeUpdatedServerEvent(presenter Presenter, rangeVm viewmodel.Range, unitMap viewmodel.UnitMap) {
	event := ObservedRangeUpdatedServerEvent{}
	event.Type = ObservedRangeUpdatedServerEventType
	event.Payload.Range = rangeVm
	event.Payload.UnitMap = unitMap
	presenter.OnSuccess(event)
}

func (serve *serve) SendRangeObservedServerEvent(presenter Presenter, rangeVm viewmodel.Range, unitMap viewmodel.UnitMap) {
	event := RangeObservedServerEvent{}
	event.Type = RangeObservedServerEventType
	event.Payload.Range = rangeVm
	event.Payload.UnitMap = unitMap
	presenter.OnSuccess(event)
}

func (serve *serve) RequestToAddPlayer(rawLiveGameId string, rawPlayerId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewAddPlayerRequestedIntgrEvent(rawLiveGameId, rawPlayerId)),
	)
}

func (serve *serve) RequestToObserveRange(rawLiveGameId string, rawPlayerId string, rangeVm viewmodel.Range) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewObserveRangeRequestedIntgrEvent(rawLiveGameId, rawPlayerId, rangeVm)),
	)
}

func (serve *serve) RequestToBuildItem(rawLiveGameId string, rawLocation viewmodel.Location, rawItemId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewBuildItemRequestedIntgrEvent(rawLiveGameId, rawLocation, rawItemId)),
	)
}

func (serve *serve) RequestToDestroyItem(rawLiveGameId string, rawLocation viewmodel.Location) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewDestroyItemRequestedIntgrEvent(rawLiveGameId, rawLocation)),
	)
}

func (serve *serve) RequestToRemovePlayer(rawLiveGameId string, rawPlayerId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewRemovePlayerRequestedIntgrEvent(rawLiveGameId, rawPlayerId)),
	)
}
