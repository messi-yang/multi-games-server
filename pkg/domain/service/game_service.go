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
		serve.playerRepo.Update(player)
		return nil
	}

	targetLocation := player.GetLocation().MoveToward(direction, 1)

	unit, unitFound := serve.unitRepo.GetUnitAt(gameId, targetLocation)
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
	serve.playerRepo.Update(player)

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

	_, anyPlayerAtTargetLocation := serve.playerRepo.GetPlayerAt(gameId, targetLocation)

	if !item.IsTraversable() && anyPlayerAtTargetLocation {
		return errors.New("cannot place non-traversable item on a location with players")
	}

	serve.unitRepo.Update(unitmodel.NewUnitAgg(gameId, targetLocation, itemId))
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
