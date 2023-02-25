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
	MovePlayer(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo, direction commonmodel.DirectionVo) error
	PlaceItem(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo, itemId itemmodel.ItemIdVo) error
	DestroyItem(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo) error
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

func (serve *gameServe) MovePlayer(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo, direction commonmodel.DirectionVo) error {
	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	if !direction.IsEqual(player.GetDirection()) {
		player.SetDirection(direction)
		err = serve.playerRepo.Update(player)
		if err != nil {
			return err
		}
		return nil
	}

	targetLocation := player.GetLocation().MoveToward(direction, 1)

	unit, unitFound, err := serve.unitRepo.GetUnitAt(gameId, targetLocation)
	if err != nil {
		return err
	}

	if unitFound {
		itemId := unit.GetItemId()
		item, _ := serve.itemRepo.Get(itemId)
		if item.IsTraversable() {
			player.SetLocation(targetLocation)
		}
	} else {
		player.SetLocation(targetLocation)
	}

	player.SetDirection(direction)
	err = serve.playerRepo.Update(player)
	if err != nil {
		return err
	}

	return nil
}

func (serve *gameServe) PlaceItem(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo, itemId itemmodel.ItemIdVo) error {
	item, err := serve.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	targetLocation := player.GetLocation().MoveToward(player.GetDirection(), 1)

	_, anyPlayerAtTargetLocation, err := serve.playerRepo.GetPlayerAt(gameId, targetLocation)
	if err != nil {
		return err
	}

	if !item.IsTraversable() && anyPlayerAtTargetLocation {
		return errors.New("cannot place non-traversable item on a location with players")
	}

	_, _, err = serve.unitRepo.GetUnitAt(gameId, targetLocation)
	if err != nil {
		return err
	}

	serve.unitRepo.Add(unitmodel.NewUnitAgg(gameId, targetLocation, itemId))

	return nil
}

func (serve *gameServe) DestroyItem(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo) error {
	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	targetLocation := player.GetLocation().MoveToward(player.GetDirection(), 1)
	serve.unitRepo.Delete(gameId, targetLocation)

	return nil
}
