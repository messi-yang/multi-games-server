package gamedomainsrv

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

type Service interface {
	EnterWorld(sharedkernelmodel.WorldId) (commonmodel.PlayerId, error)
	Move(sharedkernelmodel.WorldId, commonmodel.PlayerId, commonmodel.Direction) error
	LeaveWorld(sharedkernelmodel.WorldId, commonmodel.PlayerId) error
	ChangeHeldItem(sharedkernelmodel.WorldId, commonmodel.PlayerId, commonmodel.ItemId) error
	PlaceItem(sharedkernelmodel.WorldId, commonmodel.PlayerId) error
	RemoveItem(sharedkernelmodel.WorldId, commonmodel.PlayerId) error
}

type serve struct {
	worldRepo             worldmodel.Repo
	playerRepo            playermodel.Repo
	unitRepo              unitmodel.Repo
	itemRepo              itemmodel.Repo
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewService(
	worldRepo worldmodel.Repo,
	playerRepo playermodel.Repo,
	unitRepo unitmodel.Repo,
	itemRepo itemmodel.Repo,
	domainEventDispatcher domain.DomainEventDispatcher,
) Service {
	return &serve{worldRepo: worldRepo, playerRepo: playerRepo, unitRepo: unitRepo, itemRepo: itemRepo, domainEventDispatcher: domainEventDispatcher}
}

func (serve *serve) EnterWorld(worldId sharedkernelmodel.WorldId) (playerId commonmodel.PlayerId, err error) {
	if _, err := serve.worldRepo.Get(worldId); err != nil {
		return playerId, err
	}

	firstItem, err := serve.itemRepo.GetFirstItem()
	if err != nil {
		return playerId, err
	}
	firstItemId := firstItem.GetId()

	direction := commonmodel.NewDownDirection()
	newPlayer := playermodel.NewPlayer(
		commonmodel.NewPlayerId(uuid.New()), worldId, "Hello", commonmodel.NewPosition(0, 0), direction, &firstItemId,
	)

	if err = serve.playerRepo.Add(newPlayer); err != nil {
		return playerId, err
	}
	if err = serve.domainEventDispatcher.Dispatch(&newPlayer); err != nil {
		return playerId, err
	}
	return newPlayer.GetId(), nil
}

func (serve *serve) Move(
	worldId sharedkernelmodel.WorldId, playerId commonmodel.PlayerId, direction commonmodel.Direction,
) error {
	if _, err := serve.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	if !direction.IsEqual(player.GetDirection()) {
		player.Move(player.GetPosition(), direction)
		err = serve.playerRepo.Update(player)
		if err != nil {
			return err
		}
		return serve.domainEventDispatcher.Dispatch(&player)
	}

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
			player.Move(newItemPos, direction)
		}
	} else {
		player.Move(newItemPos, direction)
	}

	if err = serve.playerRepo.Update(player); err != nil {
		return err
	}
	return serve.domainEventDispatcher.Dispatch(&player)
}

func (serve *serve) LeaveWorld(worldId sharedkernelmodel.WorldId, playerId commonmodel.PlayerId) error {
	if _, err := serve.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	if err = serve.playerRepo.Delete(player); err != nil {
		return err
	}
	return serve.domainEventDispatcher.Dispatch(&player)
}

func (serve *serve) ChangeHeldItem(worldId sharedkernelmodel.WorldId, playerId commonmodel.PlayerId, itemId commonmodel.ItemId) error {
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
	err = serve.playerRepo.Update(player)
	if err != nil {
		return err
	}
	return serve.domainEventDispatcher.Dispatch(&player)
}

func (serve *serve) PlaceItem(worldId sharedkernelmodel.WorldId, playerId commonmodel.PlayerId) error {
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

	_, playerFound, err := serve.playerRepo.FindPlayersAt(worldId, newItemPos)
	if err != nil {
		return err
	}

	if !item.GetTraversable() && playerFound {
		return nil
	}

	newUnitDirection := player.GetDirection().Rotate().Rotate()
	newUnit := unitmodel.NewUnit(commonmodel.NewUnitId(worldId, newItemPos), worldId, newItemPos, itemId, newUnitDirection)
	if err = serve.unitRepo.Add(newUnit); err != nil {
		return err
	}
	return serve.domainEventDispatcher.Dispatch(&newUnit)
}

func (serve *serve) RemoveItem(worldId sharedkernelmodel.WorldId, playerId commonmodel.PlayerId) error {
	if _, err := serve.worldRepo.Get(worldId); err != nil {
		return err
	}

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	targetUnitPos := player.GetPositionOneStepFoward()

	unit, unitFound, err := serve.unitRepo.FindUnitAt(worldId, targetUnitPos)
	if err != nil {
		return err
	}
	if !unitFound {
		return nil
	}
	unit.Delete()
	if err = serve.unitRepo.Delete(unit); err != nil {
		return err
	}
	return serve.domainEventDispatcher.Dispatch(&unit)
}
