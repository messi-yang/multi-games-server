package gamedomainsrv

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
)

type Service interface {
	EnterWorld(commonmodel.WorldId, commonmodel.PlayerId) error
	Move(commonmodel.WorldId, commonmodel.PlayerId, commonmodel.Direction) error
	LeaveWorld(commonmodel.WorldId, commonmodel.PlayerId) error
	ChangeHeldItem(commonmodel.WorldId, commonmodel.PlayerId, commonmodel.ItemId) error
	PlaceItem(commonmodel.WorldId, commonmodel.PlayerId) error
	RemoveItem(commonmodel.WorldId, commonmodel.PlayerId) error
}

type serve struct {
	worldRepo  worldmodel.Repo
	playerRepo playermodel.Repo
	unitRepo   unitmodel.Repo
	itemRepo   itemmodel.Repo
}

func NewService(worldRepo worldmodel.Repo, playerRepo playermodel.Repo, unitRepo unitmodel.Repo, itemRepo itemmodel.Repo) Service {
	return &serve{worldRepo: worldRepo, playerRepo: playerRepo, unitRepo: unitRepo, itemRepo: itemRepo}
}

func (serve *serve) EnterWorld(worldId commonmodel.WorldId, playerId commonmodel.PlayerId) error {
	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	if _, err := serve.worldRepo.Get(worldId); err != nil {
		return err
	}

	firstItem, err := serve.itemRepo.GetFirstItem()
	if err != nil {
		return err
	}
	firstItemId := firstItem.GetId()

	direction := commonmodel.NewDownDirection()
	newPlayer := playermodel.NewPlayer(playerId, worldId, "Hello", commonmodel.NewPosition(0, 0), direction, &firstItemId)

	return serve.playerRepo.Add(newPlayer)
}

func (serve *serve) Move(
	worldId commonmodel.WorldId, playerId commonmodel.PlayerId, direction commonmodel.Direction,
) error {
	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	if _, err := serve.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	if !direction.IsEqual(player.GetDirection()) {
		player.ChangeDirection(direction)
		err = serve.playerRepo.Update(player)
		return err
	}

	player.ChangeDirection(direction)
	newItemPos := player.GetPositionOneStepFoward()

	unit, unitFound, err := serve.unitRepo.FindUnitAt(worldId, newItemPos)
	if err != nil {
		return err
	}

	if unitFound {
		itemId := unit.GetItemId()
		item, err := serve.itemRepo.Get(itemId)
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

	return serve.playerRepo.Update(player)
}

func (serve *serve) LeaveWorld(worldId commonmodel.WorldId, playerId commonmodel.PlayerId) error {
	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	if _, err := serve.worldRepo.Get(worldId); err != nil {
		return err
	}

	return serve.playerRepo.Delete(playerId)
}

func (serve *serve) ChangeHeldItem(worldId commonmodel.WorldId, playerId commonmodel.PlayerId, itemId commonmodel.ItemId) error {
	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	if _, err := serve.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	if _, err = serve.itemRepo.Get(itemId); err != nil {
		return err
	}

	player.ChangeHeldItem(itemId)
	return serve.playerRepo.Update(player)
}

func (serve *serve) PlaceItem(worldId commonmodel.WorldId, playerId commonmodel.PlayerId) error {
	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	if _, err := serve.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	playerHeldItemId := player.GetHeldItemId()
	if playerHeldItemId == nil {
		return nil
	}

	itemId := *playerHeldItemId
	item, err := serve.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	newItemPos := player.GetPositionOneStepFoward()
	if newItemPos.IsEqual(commonmodel.NewPosition(0, 0)) {
		return nil
	}

	_, unitFound, err := serve.unitRepo.FindUnitAt(worldId, newItemPos)
	if err != nil {
		return err
	}
	if unitFound {
		return nil
	}

	_, playerFound, err := serve.playerRepo.FindPlayerAt(worldId, newItemPos)
	if err != nil {
		return err
	}

	if !item.GetTraversable() && playerFound {
		return nil
	}

	newItemDirection := player.GetDirection().Rotate().Rotate()
	return serve.unitRepo.Add(unitmodel.NewUnit(worldId, newItemPos, itemId, newItemDirection))
}

func (serve *serve) RemoveItem(worldId commonmodel.WorldId, playerId commonmodel.PlayerId) error {
	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	if _, err := serve.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	newItemPos := player.GetPositionOneStepFoward()
	return serve.unitRepo.Delete(worldId, newItemPos)
}
