package gameappservice

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

func (command PlaceItemCommand) Parse() (worldmodel.WorldIdVo, playermodel.PlayerIdVo, itemmodel.ItemIdVo) {
	return worldmodel.NewWorldIdVo(command.WorldId), playermodel.NewPlayerIdVo(command.PlayerId), itemmodel.NewItemIdVo(command.ItemId)
}

type DestroyItemCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func (command DestroyItemCommand) Parse() (worldmodel.WorldIdVo, playermodel.PlayerIdVo) {
	return worldmodel.NewWorldIdVo(command.WorldId), playermodel.NewPlayerIdVo(command.PlayerId)
}

type AddPlayerCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func (command AddPlayerCommand) Parse() (worldmodel.WorldIdVo, playermodel.PlayerIdVo) {
	return worldmodel.NewWorldIdVo(command.WorldId), playermodel.NewPlayerIdVo(command.PlayerId)
}

type MovePlayerCommand struct {
	WorldId   uuid.UUID
	PlayerId  uuid.UUID
	Direction int8
}

func (command MovePlayerCommand) Parse() (worldmodel.WorldIdVo, playermodel.PlayerIdVo, commonmodel.DirectionVo) {
	return worldmodel.NewWorldIdVo(command.WorldId), playermodel.NewPlayerIdVo(command.PlayerId), commonmodel.NewDirectionVo(command.Direction)
}

type RemovePlayerCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func (command RemovePlayerCommand) Parse() (worldmodel.WorldIdVo, playermodel.PlayerIdVo) {
	return worldmodel.NewWorldIdVo(command.WorldId), playermodel.NewPlayerIdVo(command.PlayerId)
}
