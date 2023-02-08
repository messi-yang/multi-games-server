package appservice

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/jsonmarshaller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
	"github.com/samber/lo"
)

type GameAppService interface {
	LoadGame(mapSizVm viewmodel.SizeVm, gameIdVm string)
	JoinGame(gameIdVm string, playerIdVm string)
	MovePlayer(gameIdVm string, playerIdVm string, directionVm int8)
	PlaceItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm int16)
	DestroyItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm)
	LeaveGame(gameIdVm string, playerIdVm string)
}

type gameAppServe struct {
	gameRepo          gamemodel.Repo
	unitRepo          unitmodel.Repo
	itemRepo          itemmodel.Repo
	gameService       service.GameService
	IntEventPublisher intevent.IntEventPublisher
}

func NewGameAppService(
	gameRepo gamemodel.Repo,
	unitRepo unitmodel.Repo,
	itemRepo itemmodel.Repo,
	IntEventPublisher intevent.IntEventPublisher,
) GameAppService {
	return &gameAppServe{
		gameRepo:          gameRepo,
		itemRepo:          itemRepo,
		unitRepo:          unitRepo,
		gameService:       service.NewGameService(gameRepo, unitRepo, itemRepo),
		IntEventPublisher: IntEventPublisher,
	}
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

func (gameAppServe *gameAppServe) LoadGame(mapSizeVm viewmodel.SizeVm, gameIdVm string) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}

	mapSize, _ := commonmodel.NewSizeVo(mapSizeVm.Width, mapSizeVm.Height)

	items := gameAppServe.itemRepo.GetAll()
	tool.ForMatrix(mapSize.GetWidth(), mapSize.GetHeight(), func(x int, y int) {
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

func (gameAppServe *gameAppServe) JoinGame(gameIdVm string, playerIdVm string) {
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

	players := game.GetPlayers()
	playerVms := lo.Map(players, func(p gamemodel.PlayerEntity, _ int) viewmodel.PlayerVm {
		return viewmodel.NewPlayerVm(p)
	})

	// Delete this section later
	bound, _ := game.GetPlayerViewBound(playerId)
	units := gameAppServe.unitRepo.GetUnits(gameId, bound)
	view := unitmodel.NewViewVo(bound, units)
	// Delete this section later

	gameAppServe.IntEventPublisher.Publish(
		intevent.CreateGameClientChannel(gameIdVm, playerIdVm),
		jsonmarshaller.Marshal(
			intevent.NewGameJoinedIntEvent(
				gameIdVm,
				playerVms,
				viewmodel.NewSizeVm(game.GetMapSize()),
				viewmodel.NewViewVm(view),
			),
		),
	)
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
