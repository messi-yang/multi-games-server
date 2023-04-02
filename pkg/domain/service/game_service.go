package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type GameService interface {
	AddPlayer(worldmodel.WorldIdVo, playermodel.PlayerIdVo) error
	MovePlayer(worldmodel.WorldIdVo, playermodel.PlayerIdVo, commonmodel.DirectionVo) error
	RemovePlayer(worldmodel.WorldIdVo, playermodel.PlayerIdVo) error
	ChangeHeldItem(worldmodel.WorldIdVo, playermodel.PlayerIdVo, itemmodel.ItemIdVo) error
	PlaceItem(worldmodel.WorldIdVo, playermodel.PlayerIdVo) error
	DestroyItem(worldmodel.WorldIdVo, playermodel.PlayerIdVo) error
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

	if _, err := serve.worldRepository.Get(worldId); err != nil {
		return err
	}

	firstItem, err := serve.itemRepository.GetFirstItem()
	if err != nil {
		return err
	}
	firstItemId := firstItem.GetId()

	direction := commonmodel.NewDownDirectionVo()
	newPlayer := playermodel.NewPlayerAgg(playerId, worldId, "Hello", commonmodel.NewPositionVo(0, 0), direction, &firstItemId)

	return serve.playerRepository.Add(newPlayer)
}

func (serve *gameServe) MovePlayer(
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, direction commonmodel.DirectionVo,
) error {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	if _, err := serve.worldRepository.Get(worldId); err != nil {
		return err
	}

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	if !direction.IsEqual(player.GetDirection()) {
		player.ChangeDirection(direction)
		err = serve.playerRepository.Update(player)
		return err
	}

	player.ChangeDirection(direction)
	newItemPos := player.GetPositionOneStepFoward()

	unit, unitFound, err := serve.unitRepository.GetUnitAt(worldId, newItemPos)
	if err != nil {
		return err
	}

	if unitFound {
		itemId := unit.GetItemId()
		item, err := serve.itemRepository.Get(itemId)
		if err != nil {
			return err
		}
		if item.GetTraversable() {
			player.ChangePosition(newItemPos)
		}
	} else {
		player.ChangePosition(newItemPos)
	}

	if player.ShallUpdateVisionBound() {
		player.UpdateVisionBound()
	}

	return serve.playerRepository.Update(player)
}

func (serve *gameServe) RemovePlayer(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	if _, err := serve.worldRepository.Get(worldId); err != nil {
		return err
	}

	return serve.playerRepository.Delete(playerId)
}

func (serve *gameServe) ChangeHeldItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, itemId itemmodel.ItemIdVo) error {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	if _, err := serve.worldRepository.Get(worldId); err != nil {
		return err
	}

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	if _, err = serve.itemRepository.Get(itemId); err != nil {
		return err
	}

	player.ChangeHeldItem(itemId)
	return serve.playerRepository.Update(player)
}

func (serve *gameServe) PlaceItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	if _, err := serve.worldRepository.Get(worldId); err != nil {
		return err
	}

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	if !player.HasHeldItem() {
		return nil
	}

	itemId := *player.GetHeldItemId()
	item, err := serve.itemRepository.Get(itemId)
	if err != nil {
		return err
	}

	newItemPos := player.GetPositionOneStepFoward()

	_, unitFound, err := serve.unitRepository.GetUnitAt(worldId, newItemPos)
	if err != nil {
		return err
	}
	if unitFound {
		return nil
	}

	_, playerFound, err := serve.playerRepository.GetPlayerAt(worldId, newItemPos)
	if err != nil {
		return err
	}

	if !item.GetTraversable() && playerFound {
		return nil
	}

	newItemDirection := player.GetDirection().Rotate().Rotate()
	return serve.unitRepository.Add(unitmodel.NewUnitAgg(worldId, newItemPos, itemId, newItemDirection))
}

func (serve *gameServe) DestroyItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	if _, err := serve.worldRepository.Get(worldId); err != nil {
		return err
	}

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	newItemPos := player.GetPositionOneStepFoward()
	return serve.unitRepository.Delete(worldId, newItemPos)
}
