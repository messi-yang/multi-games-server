package service

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type GameService interface {
	MovePlayer(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, direction commonmodel.DirectionVo) error
	PlaceItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, itemId itemmodel.ItemIdVo) error
	DestroyItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error
}

type gameServe struct {
	worldRepo  worldmodel.Repo
	playerRepo playermodel.Repo
	unitRepo   unitmodel.Repo
	itemRepo   itemmodel.Repo
}

func NewGameService(worldRepo worldmodel.Repo, playerRepo playermodel.Repo, unitRepo unitmodel.Repo, itemRepo itemmodel.Repo) GameService {
	return &gameServe{worldRepo: worldRepo, playerRepo: playerRepo, unitRepo: unitRepo, itemRepo: itemRepo}
}

func (serve *gameServe) MovePlayer(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, direction commonmodel.DirectionVo) error {
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

	targetPosition := player.GetPosition().MoveToward(direction, 1)

	unit, unitFound, err := serve.unitRepo.GetUnitAt(worldId, targetPosition)
	if err != nil {
		return err
	}

	if unitFound {
		itemId := unit.GetItemId()
		item, _ := serve.itemRepo.Get(itemId)
		if item.IsTraversable() {
			player.SetPosition(targetPosition)
		}
	} else {
		player.SetPosition(targetPosition)
	}

	player.SetDirection(direction)
	err = serve.playerRepo.Update(player)
	if err != nil {
		return err
	}

	return nil
}

func (serve *gameServe) PlaceItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, itemId itemmodel.ItemIdVo) error {
	item, err := serve.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	targetPosition := player.GetPosition().MoveToward(player.GetDirection(), 1)

	_, anyPlayerAtTargetPosition, err := serve.playerRepo.GetPlayerAt(worldId, targetPosition)
	if err != nil {
		return err
	}

	if !item.IsTraversable() && anyPlayerAtTargetPosition {
		return errors.New("cannot place non-traversable item on a position with players")
	}

	_, _, err = serve.unitRepo.GetUnitAt(worldId, targetPosition)
	if err != nil {
		return err
	}

	serve.unitRepo.Add(unitmodel.NewUnitAgg(worldId, targetPosition, itemId))

	return nil
}

func (serve *gameServe) DestroyItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error {
	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	targetPosition := player.GetPosition().MoveToward(player.GetDirection(), 1)
	serve.unitRepo.Delete(worldId, targetPosition)

	return nil
}
