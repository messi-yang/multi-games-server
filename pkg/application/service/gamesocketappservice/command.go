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

func (command PlaceItemCommand) Validate() (worldmodel.WorldIdVo, playermodel.PlayerIdVo, itemmodel.ItemIdVo, error) {
	worldId, err := worldmodel.ParseWorldIdVo(command.WorldId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, itemmodel.ItemIdVo{}, err
	}
	playerId, err := playermodel.ParsePlayerIdVo(command.PlayerId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, itemmodel.ItemIdVo{}, err
	}
	itemId, err := itemmodel.ParseItemIdVo(command.ItemId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, itemmodel.ItemIdVo{}, err
	}

	return worldId, playerId, itemId, nil
}

type DestroyItemCommand struct {
	WorldId  string
	PlayerId string
}

func (command DestroyItemCommand) Validate() (worldmodel.WorldIdVo, playermodel.PlayerIdVo, error) {
	worldId, err := worldmodel.ParseWorldIdVo(command.WorldId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, err
	}
	playerId, err := playermodel.ParsePlayerIdVo(command.PlayerId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, err
	}

	return worldId, playerId, nil
}

type AddPlayerCommand struct {
	WorldId  string
	PlayerId string
}

func (command AddPlayerCommand) Validate() (worldmodel.WorldIdVo, playermodel.PlayerIdVo, error) {
	worldId, err := worldmodel.ParseWorldIdVo(command.WorldId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, err
	}
	playerId, err := playermodel.ParsePlayerIdVo(command.PlayerId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, err
	}

	return worldId, playerId, nil
}

type MovePlayerCommand struct {
	WorldId   string
	PlayerId  string
	Direction int8
}

func (command MovePlayerCommand) Validate() (worldmodel.WorldIdVo, playermodel.PlayerIdVo, commonmodel.DirectionVo, error) {
	worldId, err := worldmodel.ParseWorldIdVo(command.WorldId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, 0, err
	}
	playerId, err := playermodel.ParsePlayerIdVo(command.PlayerId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, 0, err
	}
	direction := commonmodel.NewDirectionVo(command.Direction)

	return worldId, playerId, direction, nil
}

type RemovePlayerCommand struct {
	WorldId  string
	PlayerId string
}

func (command RemovePlayerCommand) Validate() (worldmodel.WorldIdVo, playermodel.PlayerIdVo, error) {
	worldId, err := worldmodel.ParseWorldIdVo(command.WorldId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, err
	}
	playerId, err := playermodel.ParsePlayerIdVo(command.PlayerId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, err
	}

	return worldId, playerId, nil
}
