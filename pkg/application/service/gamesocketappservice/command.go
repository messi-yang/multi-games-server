package gamesocketappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
)

type PlaceItemCommand struct {
	GameId   string
	PlayerId string
	Location dto.LocationDto
	ItemId   int16
}

func (command PlaceItemCommand) Validate() (gamemodel.GameIdVo, gamemodel.PlayerIdVo, itemmodel.ItemIdVo, commonmodel.LocationVo, error) {
	gameId, err := gamemodel.NewGameIdVo(command.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, itemmodel.ItemIdVo{}, commonmodel.LocationVo{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(command.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, itemmodel.ItemIdVo{}, commonmodel.LocationVo{}, err
	}
	itemId := itemmodel.NewItemIdVo(command.ItemId)
	location := commonmodel.NewLocationVo(command.Location.X, command.Location.Z)

	return gameId, playerId, itemId, location, nil
}

type DestroyItemCommand struct {
	GameId   string
	PlayerId string
	Location dto.LocationDto
}

func (command DestroyItemCommand) Validate() (gamemodel.GameIdVo, gamemodel.PlayerIdVo, commonmodel.LocationVo, error) {
	gameId, err := gamemodel.NewGameIdVo(command.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, commonmodel.LocationVo{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(command.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, commonmodel.LocationVo{}, err
	}
	location := commonmodel.NewLocationVo(command.Location.X, command.Location.Z)

	return gameId, playerId, location, nil
}

type AddPlayerCommand struct {
	GameId   string
	PlayerId string
}

func (command AddPlayerCommand) Validate() (gamemodel.GameIdVo, gamemodel.PlayerIdVo, error) {
	gameId, err := gamemodel.NewGameIdVo(command.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(command.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, err
	}

	return gameId, playerId, nil
}

type MovePlayerCommand struct {
	GameId    string
	PlayerId  string
	Direction int8
}

func (command MovePlayerCommand) Validate() (gamemodel.GameIdVo, gamemodel.PlayerIdVo, gamemodel.DirectionVo, error) {
	gameId, err := gamemodel.NewGameIdVo(command.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, 0, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(command.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, 0, err
	}
	direction, err := gamemodel.NewDirectionVo(command.Direction)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, 0, err
	}

	return gameId, playerId, direction, nil
}

type RemovePlayerCommand struct {
	GameId   string
	PlayerId string
}

func (command RemovePlayerCommand) Validate() (gamemodel.GameIdVo, gamemodel.PlayerIdVo, error) {
	gameId, err := gamemodel.NewGameIdVo(command.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(command.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, err
	}

	return gameId, playerId, nil
}
