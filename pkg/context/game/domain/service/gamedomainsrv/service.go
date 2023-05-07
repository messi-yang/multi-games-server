package gamedomainsrv

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
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

func (serve *serve) EnterWorld(worldId commonmodel.WorldId, playerId commonmodel.PlayerId) error {
	if _, err := serve.worldRepo.Get(worldId); err != nil {
		return err
	}

	firstItem, err := serve.itemRepo.GetFirstItem()
	if err != nil {
		return err
	}
	firstItemId := firstItem.GetId()

	direction := commonmodel.NewDownDirection()
	newPlayer := playermodel.NewPlayer(
		playerId, worldId, "Hello", commonmodel.NewPosition(0, 0), direction, &firstItemId,
	)

	err = serve.playerRepo.Add(newPlayer)
	if err != nil {
		return err
	}
	return serve.domainEventDispatcher.Dispatch(&newPlayer)
}

func (serve *serve) Move(
	worldId commonmodel.WorldId, playerId commonmodel.PlayerId, direction commonmodel.Direction,
) error {
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
		if err != nil {
			return err
		}
		return serve.domainEventDispatcher.Dispatch(&player)
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

	err = serve.playerRepo.Update(player)
	if err != nil {
		return err
	}
	return serve.domainEventDispatcher.Dispatch(&player)
}

func (serve *serve) LeaveWorld(worldId commonmodel.WorldId, playerId commonmodel.PlayerId) error {
	if _, err := serve.worldRepo.Get(worldId); err != nil {
		return err
	}

	return serve.playerRepo.Delete(playerId)
}

func (serve *serve) ChangeHeldItem(worldId commonmodel.WorldId, playerId commonmodel.PlayerId, itemId commonmodel.ItemId) error {
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

func (serve *serve) PlaceItem(worldId commonmodel.WorldId, playerId commonmodel.PlayerId) error {
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

	newItemDirection := player.GetDirection().Rotate().Rotate()
	err = serve.unitRepo.Add(unitmodel.NewUnit(worldId, newItemPos, itemId, newItemDirection))
	if err != nil {
		return err
	}
	return serve.domainEventDispatcher.Dispatch(&player)
}

func (serve *serve) RemoveItem(worldId commonmodel.WorldId, playerId commonmodel.PlayerId) error {
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
