package gamesocketappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type PlaceItemCommand struct {
	WorldId  string
	PlayerId string
	ItemId   string
}

func (command PlaceItemCommand) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, itemId itemmodel.ItemIdVo, err error,
) {
	worldId, err = worldmodel.ParseWorldIdVo(command.WorldId)
	if err != nil {
		return worldId, playerId, itemId, err
	}
	playerId, err = playermodel.ParsePlayerIdVo(command.PlayerId)
	if err != nil {
		return worldId, playerId, itemId, err
	}
	itemId, err = itemmodel.ParseItemIdVo(command.ItemId)
	if err != nil {
		return worldId, playerId, itemId, err
	}
	return worldId, playerId, itemId, nil
}

type DestroyItemCommand struct {
	WorldId  string
	PlayerId string
}

func (command DestroyItemCommand) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, err error,
) {
	worldId, err = worldmodel.ParseWorldIdVo(command.WorldId)
	if err != nil {
		return worldId, playerId, err
	}
	playerId, err = playermodel.ParsePlayerIdVo(command.PlayerId)
	if err != nil {
		return worldId, playerId, err
	}
	return worldId, playerId, nil
}

type AddPlayerCommand struct {
	WorldId  string
	PlayerId string
}

func (command AddPlayerCommand) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, err error,
) {
	worldId, err = worldmodel.ParseWorldIdVo(command.WorldId)
	if err != nil {
		return worldId, playerId, err
	}
	playerId, err = playermodel.ParsePlayerIdVo(command.PlayerId)
	if err != nil {
		return worldId, playerId, err
	}
	return worldId, playerId, nil
}

type MovePlayerCommand struct {
	WorldId   string
	PlayerId  string
	Direction int8
}

func (command MovePlayerCommand) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, direction commonmodel.DirectionVo, err error,
) {
	worldId, err = worldmodel.ParseWorldIdVo(command.WorldId)
	if err != nil {
		return worldId, playerId, direction, err
	}
	playerId, err = playermodel.ParsePlayerIdVo(command.PlayerId)
	if err != nil {
		return worldId, playerId, direction, err
	}
	direction = commonmodel.NewDirectionVo(command.Direction)

	return worldId, playerId, direction, nil
}

type RemovePlayerCommand struct {
	WorldId  string
	PlayerId string
}

func (command RemovePlayerCommand) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, err error,
) {
	worldId, err = worldmodel.ParseWorldIdVo(command.WorldId)
	if err != nil {
		return worldId, playerId, err
	}
	playerId, err = playermodel.ParsePlayerIdVo(command.PlayerId)
	if err != nil {
		return worldId, playerId, err
	}
	return worldId, playerId, nil
}
