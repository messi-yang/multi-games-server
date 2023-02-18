package service

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
)

type GameService interface {
	MovePlayer(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo, direction gamemodel.DirectionVo) error
	PlaceItem(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo, itemId itemmodel.ItemIdVo, location commonmodel.LocationVo) error
	DestroyItem(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo, location commonmodel.LocationVo) error
}

type gameServe struct {
	gameRepo   gamemodel.Repo
	playerRepo playermodel.Repo
	unitRepo   unitmodel.Repo
	itemRepo   itemmodel.Repo
}

func NewGameService(gameRepo gamemodel.Repo, playerRepo playermodel.Repo, unitRepo unitmodel.Repo, itemRepo itemmodel.Repo) GameService {
	return &gameServe{gameRepo: gameRepo, playerRepo: playerRepo, unitRepo: unitRepo, itemRepo: itemRepo}
}

func (serve *gameServe) MovePlayer(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo, direction gamemodel.DirectionVo) error {
	player, err := serve.playerRepo.Get(playerId)
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
	serve.playerRepo.Update(player)

	return nil
}

func (serve *gameServe) PlaceItem(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo, itemId itemmodel.ItemIdVo, location commonmodel.LocationVo) error {
	item, err := serve.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	_, playerFound := serve.playerRepo.GetPlayerAt(gameId, location)

	if !item.IsTraversable() && playerFound {
		return errors.New("cannot place non-traversable item on a location with players")
	}

	serve.unitRepo.UpdateUnit(unitmodel.NewUnitAgg(gameId, location, itemId))
	return nil
}

func (serve *gameServe) DestroyItem(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo, location commonmodel.LocationVo) error {
	serve.unitRepo.DeleteUnit(gameId, location)
	return nil
}
