package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/unitmodel"
	"github.com/google/uuid"
)

var (
	errPlayerExceededBoundary          = fmt.Errorf("player exceeded the boundary of the world")
	errPositionAlreadyHasUnitOrPlayers = fmt.Errorf("the position already has an unit or players")
	errPositionDoesNotHaveUnit         = fmt.Errorf("the position does not have an unit")
)

type WorldJourneyService interface {
	EnterWorld(sharedkernelmodel.WorldId) (playermodel.PlayerId, error)
	Move(sharedkernelmodel.WorldId, playermodel.PlayerId, commonmodel.Direction) error
	LeaveWorld(sharedkernelmodel.WorldId, playermodel.PlayerId) error
	ChangeHeldItem(sharedkernelmodel.WorldId, playermodel.PlayerId, commonmodel.ItemId) error
	PlaceUnit(sharedkernelmodel.WorldId, playermodel.PlayerId) error
	RemoveUnit(sharedkernelmodel.WorldId, playermodel.PlayerId) error
}

type worldJourneyServe struct {
	worldRepo  worldmodel.WorldRepo
	playerRepo playermodel.PlayerRepo
	unitRepo   unitmodel.UnitRepo
	itemRepo   itemmodel.ItemRepo
}

func NewWorldJourneyService(
	worldRepo worldmodel.WorldRepo,
	playerRepo playermodel.PlayerRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
) WorldJourneyService {
	return &worldJourneyServe{
		worldRepo:  worldRepo,
		playerRepo: playerRepo,
		unitRepo:   unitRepo,
		itemRepo:   itemRepo,
	}
}

func (worldJourneyServe *worldJourneyServe) EnterWorld(worldId sharedkernelmodel.WorldId) (playerId playermodel.PlayerId, err error) {
	if _, err := worldJourneyServe.worldRepo.Get(worldId); err != nil {
		return playerId, err
	}

	firstItem, err := worldJourneyServe.itemRepo.GetFirstItem()
	if err != nil {
		return playerId, err
	}
	firstItemId := firstItem.GetId()

	direction := commonmodel.NewDownDirection()
	newPlayer := playermodel.NewPlayer(
		playermodel.NewPlayerId(uuid.New()), worldId, "Hello", commonmodel.NewPosition(0, 0), direction, &firstItemId,
	)

	if err = worldJourneyServe.playerRepo.Add(newPlayer); err != nil {
		return playerId, err
	}
	return newPlayer.GetId(), nil
}

func (worldJourneyServe *worldJourneyServe) Move(
	worldId sharedkernelmodel.WorldId, playerId playermodel.PlayerId, direction commonmodel.Direction,
) error {
	world, err := worldJourneyServe.worldRepo.Get(worldId)
	if err != nil {
		return err
	}

	player, err := worldJourneyServe.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	if !direction.IsEqual(player.GetDirection()) {
		player.Move(player.GetPosition(), direction)
		return worldJourneyServe.playerRepo.Update(player)
	}

	newItemPos := player.GetPositionOneStepFoward()

	unit, err := worldJourneyServe.unitRepo.GetUnitAt(worldId, newItemPos)
	if err != nil {
		return err
	}

	if unit != nil {
		itemId := unit.GetItemId()
		item, err := worldJourneyServe.itemRepo.Get(itemId)
		if err != nil {
			return err
		}
		if item.GetTraversable() {
			player.Move(newItemPos, direction)
		}
	} else {
		player.Move(newItemPos, direction)
	}

	if !world.GetBound().CoversPosition(player.GetPosition()) {
		return errPlayerExceededBoundary
	}

	return worldJourneyServe.playerRepo.Update(player)
}

func (worldJourneyServe *worldJourneyServe) LeaveWorld(worldId sharedkernelmodel.WorldId, playerId playermodel.PlayerId) error {
	if _, err := worldJourneyServe.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := worldJourneyServe.playerRepo.Get(playerId)
	if err != nil {
		return err
	}
	player.Delete()
	return worldJourneyServe.playerRepo.Delete(player)
}

func (worldJourneyServe *worldJourneyServe) ChangeHeldItem(worldId sharedkernelmodel.WorldId, playerId playermodel.PlayerId, itemId commonmodel.ItemId) error {
	if _, err := worldJourneyServe.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := worldJourneyServe.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	if _, err = worldJourneyServe.itemRepo.Get(itemId); err != nil {
		return err
	}

	player.ChangeHeldItem(itemId)
	return worldJourneyServe.playerRepo.Update(player)
}

func (worldJourneyServe *worldJourneyServe) PlaceUnit(worldId sharedkernelmodel.WorldId, playerId playermodel.PlayerId) error {
	if _, err := worldJourneyServe.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := worldJourneyServe.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	playerHeldItemId := player.GetHeldItemId()
	if playerHeldItemId == nil {
		return nil
	}

	itemId := *playerHeldItemId
	item, err := worldJourneyServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	newItemPos := player.GetPositionOneStepFoward()
	if newItemPos.IsEqual(commonmodel.NewPosition(0, 0)) {
		return nil
	}

	unit, err := worldJourneyServe.unitRepo.GetUnitAt(worldId, newItemPos)
	if err != nil {
		return err
	}
	if unit != nil {
		return errPositionAlreadyHasUnitOrPlayers
	}

	players, err := worldJourneyServe.playerRepo.GetPlayersAt(worldId, newItemPos)
	if err != nil {
		return err
	}

	if !item.GetTraversable() && len(players) > 0 {
		return errPositionAlreadyHasUnitOrPlayers
	}

	newUnitDirection := player.GetDirection().Rotate().Rotate()
	newUnit := unitmodel.NewUnit(unitmodel.NewUnitId(worldId, newItemPos), worldId, newItemPos, itemId, newUnitDirection)
	return worldJourneyServe.unitRepo.Add(newUnit)
}

func (worldJourneyServe *worldJourneyServe) RemoveUnit(worldId sharedkernelmodel.WorldId, playerId playermodel.PlayerId) error {
	if _, err := worldJourneyServe.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := worldJourneyServe.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	targetUnitPos := player.GetPositionOneStepFoward()

	unit, err := worldJourneyServe.unitRepo.GetUnitAt(worldId, targetUnitPos)
	if err != nil {
		return err
	}
	if unit == nil {
		return errPositionDoesNotHaveUnit
	}
	unit.Delete()
	return worldJourneyServe.unitRepo.Delete(*unit)
}
