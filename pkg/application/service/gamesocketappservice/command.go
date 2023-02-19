package gamesocketappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
)

type PlaceItemCommand struct {
	GameId   string
	PlayerId string
	ItemId   int16
}

func (command PlaceItemCommand) Validate() (gamemodel.GameIdVo, playermodel.PlayerIdVo, itemmodel.ItemIdVo, error) {
	gameId, err := gamemodel.NewGameIdVo(command.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, itemmodel.ItemIdVo{}, err
	}
	playerId, err := playermodel.NewPlayerIdVo(command.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, itemmodel.ItemIdVo{}, err
	}
	itemId := itemmodel.NewItemIdVo(command.ItemId)

	return gameId, playerId, itemId, nil
}

type DestroyItemCommand struct {
	GameId   string
	PlayerId string
	Location dto.LocationDto
}

func (command DestroyItemCommand) Validate() (gamemodel.GameIdVo, playermodel.PlayerIdVo, commonmodel.LocationVo, error) {
	gameId, err := gamemodel.NewGameIdVo(command.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, commonmodel.LocationVo{}, err
	}
	playerId, err := playermodel.NewPlayerIdVo(command.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, commonmodel.LocationVo{}, err
	}
	location := commonmodel.NewLocationVo(command.Location.X, command.Location.Z)

	return gameId, playerId, location, nil
}

type AddPlayerCommand struct {
	GameId   string
	PlayerId string
}

func (command AddPlayerCommand) Validate() (gamemodel.GameIdVo, playermodel.PlayerIdVo, error) {
	gameId, err := gamemodel.NewGameIdVo(command.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, err
	}
	playerId, err := playermodel.NewPlayerIdVo(command.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, err
	}

	return gameId, playerId, nil
}

type MovePlayerCommand struct {
	GameId    string
	PlayerId  string
	Direction int8
}

func (command MovePlayerCommand) Validate() (gamemodel.GameIdVo, playermodel.PlayerIdVo, commonmodel.DirectionVo, error) {
	gameId, err := gamemodel.NewGameIdVo(command.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, 0, err
	}
	playerId, err := playermodel.NewPlayerIdVo(command.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, 0, err
	}
	direction, err := commonmodel.NewDirectionVo(command.Direction)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, 0, err
	}

	return gameId, playerId, direction, nil
}

type RemovePlayerCommand struct {
	GameId   string
	PlayerId string
}

func (command RemovePlayerCommand) Validate() (gamemodel.GameIdVo, playermodel.PlayerIdVo, error) {
	gameId, err := gamemodel.NewGameIdVo(command.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, err
	}
	playerId, err := playermodel.NewPlayerIdVo(command.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, err
	}

	return gameId, playerId, nil
}
