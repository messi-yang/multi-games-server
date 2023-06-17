package service

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/worldmodel/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

type GameService interface {
	EnterWorld(sharedkernelmodel.WorldId) (playermodel.PlayerId, error)
	Move(sharedkernelmodel.WorldId, playermodel.PlayerId, commonmodel.Direction) error
	LeaveWorld(sharedkernelmodel.WorldId, playermodel.PlayerId) error
	ChangeHeldItem(sharedkernelmodel.WorldId, playermodel.PlayerId, commonmodel.ItemId) error
	PlaceItem(sharedkernelmodel.WorldId, playermodel.PlayerId) error
	RemoveItem(sharedkernelmodel.WorldId, playermodel.PlayerId) error
}

type gamerServe struct {
	worldRepo  worldmodel.WorldRepo
	playerRepo playermodel.PlayerRepo
	unitRepo   unitmodel.UnitRepo
	itemRepo   itemmodel.ItemRepo
}

func NewGameService(
	worldRepo worldmodel.WorldRepo,
	playerRepo playermodel.PlayerRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
) GameService {
	return &gamerServe{
		worldRepo:  worldRepo,
		playerRepo: playerRepo,
		unitRepo:   unitRepo,
		itemRepo:   itemRepo,
	}
}

func (gamerServe *gamerServe) EnterWorld(worldId sharedkernelmodel.WorldId) (playerId playermodel.PlayerId, err error) {
	if _, err := gamerServe.worldRepo.Get(worldId); err != nil {
		return playerId, err
	}

	firstItem, err := gamerServe.itemRepo.GetFirstItem()
	if err != nil {
		return playerId, err
	}
	firstItemId := firstItem.GetId()

	direction := commonmodel.NewDownDirection()
	newPlayer := playermodel.NewPlayer(
		playermodel.NewPlayerId(uuid.New()), worldId, "Hello", commonmodel.NewPosition(0, 0), direction, &firstItemId,
	)

	if err = gamerServe.playerRepo.Add(newPlayer); err != nil {
		return playerId, err
	}
	return newPlayer.GetId(), nil
}

func (gamerServe *gamerServe) Move(
	worldId sharedkernelmodel.WorldId, playerId playermodel.PlayerId, direction commonmodel.Direction,
) error {
	if _, err := gamerServe.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := gamerServe.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	if !direction.IsEqual(player.GetDirection()) {
		player.Move(player.GetPosition(), direction)
		return gamerServe.playerRepo.Update(player)
	}

	newItemPos := player.GetPositionOneStepFoward()

	unit, unitFound, err := gamerServe.unitRepo.FindUnitAt(worldId, newItemPos)
	if err != nil {
		return err
	}

	if unitFound {
		itemId := unit.GetItemId()
		item, err := gamerServe.itemRepo.Get(itemId)
		if err != nil {
			return err
		}
		if item.GetTraversable() {
			player.Move(newItemPos, direction)
		}
	} else {
		player.Move(newItemPos, direction)
	}

	return gamerServe.playerRepo.Update(player)
}

func (gamerServe *gamerServe) LeaveWorld(worldId sharedkernelmodel.WorldId, playerId playermodel.PlayerId) error {
	if _, err := gamerServe.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := gamerServe.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	return gamerServe.playerRepo.Delete(player)
}

func (gamerServe *gamerServe) ChangeHeldItem(worldId sharedkernelmodel.WorldId, playerId playermodel.PlayerId, itemId commonmodel.ItemId) error {
	if _, err := gamerServe.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := gamerServe.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	if _, err = gamerServe.itemRepo.Get(itemId); err != nil {
		return err
	}

	player.ChangeHeldItem(itemId)
	return gamerServe.playerRepo.Update(player)
}

func (gamerServe *gamerServe) PlaceItem(worldId sharedkernelmodel.WorldId, playerId playermodel.PlayerId) error {
	if _, err := gamerServe.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := gamerServe.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	playerHeldItemId := player.GetHeldItemId()
	if playerHeldItemId == nil {
		return nil
	}

	itemId := *playerHeldItemId
	item, err := gamerServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	newItemPos := player.GetPositionOneStepFoward()
	if newItemPos.IsEqual(commonmodel.NewPosition(0, 0)) {
		return nil
	}

	_, unitFound, err := gamerServe.unitRepo.FindUnitAt(worldId, newItemPos)
	if err != nil {
		return err
	}
	if unitFound {
		return nil
	}

	_, playerFound, err := gamerServe.playerRepo.FindPlayersAt(worldId, newItemPos)
	if err != nil {
		return err
	}

	if !item.GetTraversable() && playerFound {
		return nil
	}

	newUnitDirection := player.GetDirection().Rotate().Rotate()
	newUnit := unitmodel.NewUnit(unitmodel.NewUnitId(worldId, newItemPos), worldId, newItemPos, itemId, newUnitDirection)
	return gamerServe.unitRepo.Add(newUnit)
}

func (gamerServe *gamerServe) RemoveItem(worldId sharedkernelmodel.WorldId, playerId playermodel.PlayerId) error {
	if _, err := gamerServe.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := gamerServe.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	targetUnitPos := player.GetPositionOneStepFoward()

	unit, unitFound, err := gamerServe.unitRepo.FindUnitAt(worldId, targetUnitPos)
	if err != nil {
		return err
	}
	if !unitFound {
		return nil
	}
	unit.Delete()
	return gamerServe.unitRepo.Delete(unit)
}
