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
	AddPlayer(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error
	MovePlayer(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, direction commonmodel.DirectionVo) error
	RemovePlayer(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error
	PlaceItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, itemId itemmodel.ItemIdVo) error
	DestroyItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error
}

type gameServe struct {
	worldRepository  worldmodel.Repository
	playerRepository playermodel.Repository
	unitRepository   unitmodel.Repository
	itemRepository   itemmodel.Repository
}

func NewGameService(worldRepository worldmodel.Repository, playerRepository playermodel.Repository, unitRepository unitmodel.Repository, itemRepository itemmodel.Repository) GameService {
	return &gameServe{worldRepository: worldRepository, playerRepository: playerRepository, unitRepository: unitRepository, itemRepository: itemRepository}
}

func (serve *gameServe) AddPlayer(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	direction, _ := commonmodel.NewDirectionVo(2)
	newPlayer := playermodel.NewPlayerAgg(playerId, worldId, "Hello", commonmodel.NewPositionVo(0, 0), direction)

	err := serve.playerRepository.Add(newPlayer)
	if err != nil {
		return err
	}

	return nil
}

func (serve *gameServe) MovePlayer(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, direction commonmodel.DirectionVo) error {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	if !direction.IsEqual(player.GetDirection()) {
		player.ChangeDirection(direction)
		err = serve.playerRepository.Update(player)
		if err != nil {
			return err
		}
		return nil
	}

	targetPosition := player.GetPosition().MoveToward(direction, 1)

	unit, unitFound, err := serve.unitRepository.GetUnitAt(worldId, targetPosition)
	if err != nil {
		return err
	}

	if unitFound {
		itemId := unit.GetItemId()
		item, _ := serve.itemRepository.Get(itemId)
		if item.GetTraversable() {
			player.ChangePosition(targetPosition)
		}
	} else {
		player.ChangePosition(targetPosition)
	}

	player.ChangeDirection(direction)
	err = serve.playerRepository.Update(player)
	if err != nil {
		return err
	}

	return nil
}

func (serve *gameServe) RemovePlayer(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	err := serve.playerRepository.Delete(playerId)
	if err != nil {
		return err
	}
	return nil
}

func (serve *gameServe) PlaceItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, itemId itemmodel.ItemIdVo) error {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	item, err := serve.itemRepository.Get(itemId)
	if err != nil {
		return err
	}

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	targetPosition := player.GetPosition().MoveToward(player.GetDirection(), 1)

	_, playerFound, err := serve.playerRepository.GetPlayerAt(worldId, targetPosition)
	if err != nil {
		return err
	}

	if !item.GetTraversable() && playerFound {
		return errors.New("cannot place non-traversable item on a position with players")
	}

	serve.unitRepository.Add(unitmodel.NewUnitAgg(worldId, targetPosition, itemId))

	return nil
}

func (serve *gameServe) DestroyItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	targetPosition := player.GetPosition().MoveToward(player.GetDirection(), 1)
	serve.unitRepository.Delete(worldId, targetPosition)

	return nil
}
