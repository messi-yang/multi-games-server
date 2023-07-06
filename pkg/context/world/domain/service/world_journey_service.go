package service

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/unitmodel"
	"github.com/google/uuid"
)

type WorldJourneyService interface {
	EnterWorld(sharedkernelmodel.WorldId) (playermodel.PlayerId, error)
	Move(sharedkernelmodel.WorldId, playermodel.PlayerId, commonmodel.Direction) error
	LeaveWorld(sharedkernelmodel.WorldId, playermodel.PlayerId) error
	ChangeHeldItem(sharedkernelmodel.WorldId, playermodel.PlayerId, commonmodel.ItemId) error
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

	player, err := worldJourneyServe.playerRepo.Get(worldId, playerId)
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
	player, err := worldJourneyServe.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}
	return worldJourneyServe.playerRepo.Delete(player)
}

func (worldJourneyServe *worldJourneyServe) ChangeHeldItem(worldId sharedkernelmodel.WorldId, playerId playermodel.PlayerId, itemId commonmodel.ItemId) error {
	player, err := worldJourneyServe.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}

	player.ChangeHeldItem(itemId)
	return worldJourneyServe.playerRepo.Update(player)
}
