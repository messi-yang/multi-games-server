package service

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
)

type GameService interface {
	MovePlayer(gameId gamemodel.GameIdVo, playerId gamemodel.PlayerIdVo, direction gamemodel.DirectionVo) error
	PlaceItem(gameId gamemodel.GameIdVo, playerId gamemodel.PlayerIdVo, itemId itemmodel.ItemIdVo, location commonmodel.LocationVo) error
	DestroyItem(gameId gamemodel.GameIdVo, playerId gamemodel.PlayerIdVo, location commonmodel.LocationVo) error
}

type gameServe struct {
	gameRepo gamemodel.Repo
	unitRepo unitmodel.Repo
	itemRepo itemmodel.Repo
}

func NewGameService(gameRepo gamemodel.Repo, unitRepo unitmodel.Repo, itemRepo itemmodel.Repo) GameService {
	return &gameServe{gameRepo: gameRepo, unitRepo: unitRepo, itemRepo: itemRepo}
}

func (serve *gameServe) MovePlayer(gameId gamemodel.GameIdVo, playerId gamemodel.PlayerIdVo, direction gamemodel.DirectionVo) error {
	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return err
	}

	player, err := game.GetPlayer(playerId)
	if err != nil {
		return err
	}

	newLocation := player.GetLocation()
	if direction.IsUp() {
		newLocation = newLocation.Shift(0, -1)
	} else if direction.IsRight() {
		newLocation = newLocation.Shift(1, 0)
	} else if direction.IsDown() {
		newLocation = newLocation.Shift(0, 1)
	} else if direction.IsLeft() {
		newLocation = newLocation.Shift(-1, 0)
	}

	unit, err := serve.unitRepo.GetUnit(gameId, newLocation)
	if err == nil {
		itemId := unit.GetItemId()
		item, _ := serve.itemRepo.Get(itemId)
		if !item.IsTraversable() {
			return errors.New("this item is not traversable")
		}
	}

	player.SetLocation(newLocation)
	game.UpdatePlayer(player)
	serve.gameRepo.Update(gameId, game)

	return nil
}

func (serve *gameServe) PlaceItem(gameId gamemodel.GameIdVo, playerId gamemodel.PlayerIdVo, itemId itemmodel.ItemIdVo, location commonmodel.LocationVo) error {
	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return err
	}

	item, err := serve.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.IsTraversable() && game.DoesLocationHavePlayer(location) {
		return errors.New("cannot place non-traversable item on a location with players")
	}

	serve.unitRepo.UpdateUnit(unitmodel.NewUnitAgg(gameId, location, itemId))
	return nil
}

func (serve *gameServe) DestroyItem(gameId gamemodel.GameIdVo, playerId gamemodel.PlayerIdVo, location commonmodel.LocationVo) error {
	_, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return err
	}

	serve.unitRepo.DeleteUnit(gameId, location)
	return nil
}
