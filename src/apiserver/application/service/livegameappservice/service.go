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
	SendGameJoinedServerEvent(presenter Presenter, playerIdVm string)
	SendInformationUpdatedServerEvent(presenter Presenter, mapSizeVm viewmodel.MapSizeVm)
	SendObservedRangeUpdatedServerEvent(presenter Presenter, rangeVm viewmodel.RangeVm, mapVm viewmodel.MapVm)
	SendRangeObservedServerEvent(presenter Presenter, rangeVm viewmodel.RangeVm, mapVm viewmodel.MapVm)
	RequestToJoinLiveGame(liveGameIdVm string, playerIdVm string)
	RequestToObserveRange(liveGameIdVm string, playerIdVm string, rangeVm viewmodel.RangeVm)
	RequestToBuildItem(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string)
	RequestToDestroyItem(liveGameIdVm string, locationVm viewmodel.LocationVm)
	RequestToLeaveLiveGame(liveGameIdVm string, playerIdVm string)
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

func (serve *serve) SendGameJoinedServerEvent(presenter Presenter, playerIdVm string) {
	event := GameJoinedServerEvent{}
	event.Type = GameJoinedServerEventType
	event.Payload.PlayerId = playerIdVm
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

func (serve *serve) RequestToJoinLiveGame(liveGameIdVm string, playerIdVm string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewJoinLiveGameRequestedIntgrEvent(liveGameIdVm, playerIdVm)),
	)
}

func (serve *serve) RequestToObserveRange(liveGameIdVm string, playerIdVm string, rangeVm viewmodel.RangeVm) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewObserveRangeRequestedIntgrEvent(liveGameIdVm, playerIdVm, rangeVm)),
	)
}

func (serve *serve) RequestToBuildItem(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewBuildItemRequestedIntgrEvent(liveGameIdVm, locationVm, itemIdVm)),
	)
}

func (serve *serve) RequestToDestroyItem(liveGameIdVm string, locationVm viewmodel.LocationVm) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewDestroyItemRequestedIntgrEvent(liveGameIdVm, locationVm)),
	)
}

func (serve *serve) RequestToLeaveLiveGame(liveGameIdVm string, playerIdVm string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewLeaveLiveGameRequestedIntgrEvent(liveGameIdVm, playerIdVm)),
	)
}
