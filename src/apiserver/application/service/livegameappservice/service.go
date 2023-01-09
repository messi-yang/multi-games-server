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
	SendGameJoinedServerEvent(presenter Presenter, rawPlayerId string)
	SendInformationUpdatedServerEvent(presenter Presenter, mapSizeVm viewmodel.MapSizeVm)
	SendObservedRangeUpdatedServerEvent(presenter Presenter, rangeVm viewmodel.RangeVm, mapVm viewmodel.MapVm)
	SendRangeObservedServerEvent(presenter Presenter, rangeVm viewmodel.RangeVm, mapVm viewmodel.MapVm)
	RequestToJoinLiveGame(rawLiveGameId string, rawPlayerId string)
	RequestToObserveRange(rawLiveGameId string, rawPlayerId string, rangeVm viewmodel.RangeVm)
	RequestToBuildItem(rawLiveGameId string, locationVm viewmodel.LocationVm, rawItemId string)
	RequestToDestroyItem(rawLiveGameId string, locationVm viewmodel.LocationVm)
	RequestToLeaveLiveGame(rawLiveGameId string, rawPlayerId string)
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

func (serve *serve) SendGameJoinedServerEvent(presenter Presenter, rawPlayerId string) {
	event := GameJoinedServerEvent{}
	event.Type = GameJoinedServerEventType
	event.Payload.PlayerId = rawPlayerId
	presenter.OnSuccess(event)
}

func (serve *serve) SendInformationUpdatedServerEvent(presenter Presenter, mapSizeVm viewmodel.MapSizeVm) {
	event := InformationUpdatedServerEvent{}
	event.Type = InformationUpdatedServerEventType
	event.Payload.MapSize = mapSizeVm
	presenter.OnSuccess(event)
}

func (serve *serve) QueryItems(presenter Presenter) {
	items := serve.itemRepo.GetAllItems()
	itemVms := lo.Map(items, func(item itemmodel.Item, _ int) viewmodel.ItemVm {
		return viewmodel.NewItemVm(item)
	})

	event := ItemsUpdatedServerEvent{}
	event.Type = ItemsUpdatedServerEventType
	event.Payload.Items = itemVms
	presenter.OnSuccess(event)
}

func (serve *serve) SendObservedRangeUpdatedServerEvent(presenter Presenter, rangeVm viewmodel.RangeVm, mapVm viewmodel.MapVm) {
	event := ObservedRangeUpdatedServerEvent{}
	event.Type = ObservedRangeUpdatedServerEventType
	event.Payload.Range = rangeVm
	event.Payload.Map = mapVm
	presenter.OnSuccess(event)
}

func (serve *serve) SendRangeObservedServerEvent(presenter Presenter, rangeVm viewmodel.RangeVm, mapVm viewmodel.MapVm) {
	event := RangeObservedServerEvent{}
	event.Type = RangeObservedServerEventType
	event.Payload.Range = rangeVm
	event.Payload.Map = mapVm
	presenter.OnSuccess(event)
}

func (serve *serve) RequestToJoinLiveGame(rawLiveGameId string, rawPlayerId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewJoinLiveGameRequestedIntgrEvent(rawLiveGameId, rawPlayerId)),
	)
}

func (serve *serve) RequestToObserveRange(rawLiveGameId string, rawPlayerId string, rangeVm viewmodel.RangeVm) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewObserveRangeRequestedIntgrEvent(rawLiveGameId, rawPlayerId, rangeVm)),
	)
}

func (serve *serve) RequestToBuildItem(rawLiveGameId string, locationVm viewmodel.LocationVm, rawItemId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewBuildItemRequestedIntgrEvent(rawLiveGameId, locationVm, rawItemId)),
	)
}

func (serve *serve) RequestToDestroyItem(rawLiveGameId string, locationVm viewmodel.LocationVm) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewDestroyItemRequestedIntgrEvent(rawLiveGameId, locationVm)),
	)
}

func (serve *serve) RequestToLeaveLiveGame(rawLiveGameId string, rawPlayerId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewLeaveLiveGameRequestedIntgrEvent(rawLiveGameId, rawPlayerId)),
	)
}
