package gamesocketappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/google/uuid"
)

type PlaceItemCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
	ItemId   uuid.UUID
}

func (command PlaceItemCommand) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, itemId itemmodel.ItemIdVo, err error,
) {
	worldId = worldmodel.NewWorldIdVo(command.WorldId)
	playerId = playermodel.NewPlayerIdVo(command.PlayerId)
	itemId = itemmodel.NewItemIdVo(command.ItemId)
	return worldId, playerId, itemId, nil
}

type DestroyItemCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func (command DestroyItemCommand) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, err error,
) {
	worldId = worldmodel.NewWorldIdVo(command.WorldId)
	playerId = playermodel.NewPlayerIdVo(command.PlayerId)
	return worldId, playerId, nil
}

type AddPlayerCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func (command AddPlayerCommand) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, err error,
) {
	worldId = worldmodel.NewWorldIdVo(command.WorldId)
	playerId = playermodel.NewPlayerIdVo(command.PlayerId)
	return worldId, playerId, nil
}

type MovePlayerCommand struct {
	WorldId   uuid.UUID
	PlayerId  uuid.UUID
	Direction int8
}

func (command MovePlayerCommand) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, direction commonmodel.DirectionVo, err error,
) {
	worldId = worldmodel.NewWorldIdVo(command.WorldId)
	playerId = playermodel.NewPlayerIdVo(command.PlayerId)
	direction = commonmodel.NewDirectionVo(command.Direction)

	return worldId, playerId, direction, nil
}

type RemovePlayerCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func (command RemovePlayerCommand) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, err error,
) {
	worldId = worldmodel.NewWorldIdVo(command.WorldId)
	playerId = playermodel.NewPlayerIdVo(command.PlayerId)
	return worldId, playerId, nil
}
