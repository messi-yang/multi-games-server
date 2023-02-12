package appservice

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/library/jsonmarshaller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/library/tool"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/presenter"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/service"
	"github.com/samber/lo"
)

type GameAppService interface {
	SendErroredServerEvent(presenter presenter.SocketPresenter, clientMessage string)
	HandlePlayersUpdatedEvent(presenter presenter.SocketPresenter, intEvent intevent.PlayersUpdatedIntEvent)
	HandleViewUpdatedEvent(presenter presenter.SocketPresenter, intEvent intevent.ViewUpdatedIntEvent)
	LoadGame(gameIdVm string)
	JoinGame(presenter presenter.SocketPresenter, gameIdVm string, playerIdVm string)
	MovePlayer(gameIdVm string, playerIdVm string, directionVm int8)
	LeaveGame(gameIdVm string, playerIdVm string)
	PlaceItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm int16)
	DestroyItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm)
}

type gameAppServe struct {
	IntEventPublisher intevent.IntEventPublisher
	gameRepo          gamemodel.Repo
	unitRepo          unitmodel.Repo
	itemRepo          itemmodel.Repo
	gameService       service.GameService
}

func NewGameAppService(IntEventPublisher intevent.IntEventPublisher, gameRepo gamemodel.Repo, unitRepo unitmodel.Repo, itemRepo itemmodel.Repo) GameAppService {
	return &gameAppServe{
		IntEventPublisher: IntEventPublisher,
		gameRepo:          gameRepo,
		unitRepo:          unitRepo,
		itemRepo:          itemRepo,
		gameService:       service.NewGameService(gameRepo, unitRepo, itemRepo),
	}
}

func (gameAppServe *gameAppServe) publishPlayersUpdatedEvents(
	gameId gamemodel.GameIdVo,
	players []gamemodel.PlayerEntity,
	toPlayerIds []gamemodel.PlayerIdVo,
) {
	playerVms := lo.Map(players, func(player gamemodel.PlayerEntity, _ int) viewmodel.PlayerVm {
		return viewmodel.NewPlayerVm(player)
	})
	lo.ForEach(toPlayerIds, func(playerId gamemodel.PlayerIdVo, _ int) {
		gameAppServe.IntEventPublisher.Publish(
			intevent.CreateGameClientChannel(gameId.ToString(), playerId.ToString()),
			jsonmarshaller.Marshal(intevent.NewPlayersUpdatedIntEvent(
				gameId.ToString(),
				playerVms,
			)))
	})
}

func (gameAppServe *gameAppServe) publishViewUpdatedEvents(gameId gamemodel.GameIdVo, location commonmodel.LocationVo) error {
	game, err := gameAppServe.gameRepo.Get(gameId)
	if err != nil {
		return err
	}

	for _, playerId := range game.GetPlayerIds() {
		if !game.CanPlayerSeeAnyLocations(playerId, []commonmodel.LocationVo{location}) {
			continue
		}

		// Delete this section later
		bound, _ := game.GetPlayerViewBound(playerId)
		units := gameAppServe.unitRepo.GetUnits(gameId, bound)
		view := unitmodel.NewViewVo(bound, units)
		// Delete this section later

		gameAppServe.IntEventPublisher.Publish(
			intevent.CreateGameClientChannel(gameId.ToString(), playerId.ToString()),
			jsonmarshaller.Marshal(intevent.NewViewUpdatedIntEvent(
				gameId.ToString(),
				playerId.ToString(),
				viewmodel.NewViewVm(view),
			)))
	}

	return nil
}

func (gameAppServe *gameAppServe) SendErroredServerEvent(presenter presenter.SocketPresenter, clientMessage string) {
	event := ErroredServerEvent{}
	event.Type = ErroredServerEventType
	event.Payload.ClientMessage = clientMessage
	presenter.OnMessage(event)
}

func (gameAppServe *gameAppServe) HandlePlayersUpdatedEvent(presenter presenter.SocketPresenter, intEvent intevent.PlayersUpdatedIntEvent) {
	event := PlayersUpdatedServerEvent{}
	event.Type = PlayersUpdatedServerEventType
	event.Payload.Players = intEvent.Players
	presenter.OnMessage(event)
}

func (gameAppServe *gameAppServe) HandleViewUpdatedEvent(presenter presenter.SocketPresenter, intEvent intevent.ViewUpdatedIntEvent) {
	event := ViewUpdatedServerEvent{}
	event.Type = ViewUpdatedServerEventType
	event.Payload.View = intEvent.View
	presenter.OnMessage(event)
}

func (gameAppServe *gameAppServe) LoadGame(gameIdVm string) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}

	items := gameAppServe.itemRepo.GetAll()
	tool.RangeMatrix(200, 200, func(x int, y int) {
		randomInt := rand.Intn(17)
		location := commonmodel.NewLocationVo(x, y)
		if randomInt < 2 {
			newUnit := unitmodel.NewUnitAgg(gameId, location, items[randomInt].GetId())
			gameAppServe.unitRepo.UpdateUnit(newUnit)
		}
	})

	newGame := gamemodel.NewGameAgg(gameId)

	gameAppServe.gameRepo.Add(newGame)
}

func (gameAppServe *gameAppServe) JoinGame(presenter presenter.SocketPresenter, gameIdVm string, playerIdVm string) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}

	unlocker := gameAppServe.gameRepo.LockAccess(gameId)
	defer unlocker()

	game, err := gameAppServe.gameRepo.Get(gameId)
	if err != nil {
		return
	}

	err = game.AddPlayer(playerId)
	if err != nil {
		return
	}

	gameAppServe.gameRepo.Update(gameId, game)

	items := gameAppServe.itemRepo.GetAll()
	itemVms := lo.Map(items, func(item itemmodel.ItemAgg, _ int) viewmodel.ItemVm {
		return viewmodel.NewItemVm(item)
	})

	players := game.GetPlayers()
	playerVms := lo.Map(players, func(p gamemodel.PlayerEntity, _ int) viewmodel.PlayerVm {
		return viewmodel.NewPlayerVm(p)
	})

	// Delete this section later
	bound, _ := game.GetPlayerViewBound(playerId)
	units := gameAppServe.unitRepo.GetUnits(gameId, bound)
	view := unitmodel.NewViewVo(bound, units)
	// Delete this section later

	event := GameJoinedServerEvent{}
	event.Type = GameJoinedServerEventType
	event.Payload.Items = itemVms
	event.Payload.PlayerId = playerIdVm
	event.Payload.Players = playerVms
	event.Payload.View = viewmodel.NewViewVm(view)
	presenter.OnMessage(event)

	gameAppServe.publishPlayersUpdatedEvents(gameId, game.GetPlayers(), game.GetPlayerIdsExcept(playerId))
}

func (gameAppServe *gameAppServe) MovePlayer(gameIdVm string, playerIdVm string, directionVm int8) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}
	direction, err := gamemodel.NewDirectionVo(directionVm)
	if err != nil {
		return
	}

	unlocker := gameAppServe.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = gameAppServe.gameService.MovePlayer(gameId, playerId, direction)
	if err != nil {
		return
	}

	game, err := gameAppServe.gameRepo.Get(gameId)
	if err != nil {
		return
	}

	gameAppServe.publishPlayersUpdatedEvents(gameId, game.GetPlayers(), game.GetPlayerIds())

	// Delete this section later
	bound, _ := game.GetPlayerViewBound(playerId)
	units := gameAppServe.unitRepo.GetUnits(gameId, bound)
	view := unitmodel.NewViewVo(bound, units)
	// Delete this section later

	gameAppServe.IntEventPublisher.Publish(
		intevent.CreateGameClientChannel(gameIdVm, playerIdVm),
		jsonmarshaller.Marshal(intevent.NewViewUpdatedIntEvent(gameIdVm, playerIdVm, viewmodel.NewViewVm(view))),
	)
}

func (gameAppServe *gameAppServe) LeaveGame(gameIdVm string, playerIdVm string) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}

	unlocker := gameAppServe.gameRepo.LockAccess(gameId)
	defer unlocker()

	game, err := gameAppServe.gameRepo.Get(gameId)
	if err != nil {
		return
	}

	game.RemovePlayer(playerId)
	gameAppServe.gameRepo.Update(gameId, game)
	gameAppServe.publishPlayersUpdatedEvents(gameId, game.GetPlayers(), game.GetPlayerIdsExcept(playerId))
}

func (gameAppServe *gameAppServe) PlaceItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm int16) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}
	itemId := itemmodel.NewItemIdVo(itemIdVm)
	location := commonmodel.NewLocationVo(locationVm.X, locationVm.Y)

	unlocker := gameAppServe.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = gameAppServe.gameService.PlaceItem(gameId, playerId, itemId, location)
	if err != nil {
		return
	}

	gameAppServe.publishViewUpdatedEvents(gameId, location)
}

func (gameAppServe *gameAppServe) DestroyItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}
	location := commonmodel.NewLocationVo(locationVm.X, locationVm.Y)

	unlocker := gameAppServe.gameRepo.LockAccess(gameId)
	defer unlocker()

	gameAppServe.gameService.DestroyItem(gameId, playerId, location)

	gameAppServe.publishViewUpdatedEvents(gameId, location)
}
