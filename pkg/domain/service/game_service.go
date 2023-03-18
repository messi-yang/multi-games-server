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
	AddPlayer(worldmodel.WorldIdVo, playermodel.PlayerIdVo) error
	MovePlayer(worldmodel.WorldIdVo, playermodel.PlayerIdVo, commonmodel.DirectionVo) (isVisionBoundUpdated bool, err error)
	RemovePlayer(worldmodel.WorldIdVo, playermodel.PlayerIdVo) error
	PlaceItem(worldmodel.WorldIdVo, playermodel.PlayerIdVo, itemmodel.ItemIdVo) error
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

func (serve *gameServe) AddPlayer(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) (err error) {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	if _, err = serve.worldRepository.Get(worldId); err != nil {
		return err
	}

	direction := commonmodel.NewDownDirectionVo()
	newPlayer := playermodel.NewPlayerAgg(playerId, worldId, "Hello", commonmodel.NewPositionVo(0, 0), direction)

	return serve.playerRepository.Add(newPlayer)
}

func (serve *gameServe) MovePlayer(
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, direction commonmodel.DirectionVo,
) (isVisionBoundUpdated bool, err error) {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	if _, err = serve.worldRepository.Get(worldId); err != nil {
		return isVisionBoundUpdated, err
	}

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return isVisionBoundUpdated, err
	}

	if !direction.IsEqual(player.GetDirection()) {
		player.ChangeDirection(direction)
		err = serve.playerRepository.Update(player)
		return isVisionBoundUpdated, err
	}

	player.ChangeDirection(direction)
	positionOneStepFoward := player.GetPositionOneStepFoward()

	unit, unitFound, err := serve.unitRepository.GetUnitAt(worldId, positionOneStepFoward)
	if err != nil {
		return isVisionBoundUpdated, err
	}

	if unitFound {
		itemId := unit.GetItemId()
		item, err := serve.itemRepository.Get(itemId)
		if err != nil {
			return isVisionBoundUpdated, err
		}
		if item.GetTraversable() {
			player.ChangePosition(positionOneStepFoward)
		}
	} else {
		player.ChangePosition(positionOneStepFoward)
	}

	if player.ShallUpdateVisionBound() {
		player.UpdateVisionBound()
		isVisionBoundUpdated = true
	}

	err = serve.playerRepository.Update(player)
	return isVisionBoundUpdated, err
}

func (serve *gameServe) RemovePlayer(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) (err error) {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	if _, err = serve.worldRepository.Get(worldId); err != nil {
		return err
	}

	return serve.playerRepository.Delete(playerId)
}

func (serve *gameServe) PlaceItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, itemId itemmodel.ItemIdVo) (err error) {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	if _, err = serve.worldRepository.Get(worldId); err != nil {
		return err
	}

	item, err := serve.itemRepository.Get(itemId)
	if err != nil {
		return err
	}

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	positionOneStepFoward := player.GetPositionOneStepFoward()

	_, playerFound, err := serve.playerRepository.GetPlayerAt(worldId, positionOneStepFoward)
	if err != nil {
		return err
	}

	if !item.GetTraversable() && playerFound {
		return errors.New("cannot place non-traversable item on a position with players")
	}

	itemDirection := player.GetDirection().Rotate().Rotate()
	return serve.unitRepository.Add(unitmodel.NewUnitAgg(worldId, positionOneStepFoward, itemId, itemDirection))
}

func (serve *gameServe) DestroyItem(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) (err error) {
	unlocker := serve.worldRepository.LockAccess(worldId)
	defer unlocker()

	if _, err := serve.worldRepository.Get(worldId); err != nil {
		return err
	}

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	positionOneStepFoward := player.GetPositionOneStepFoward()
	return serve.unitRepository.Delete(worldId, positionOneStepFoward)
}
