package gamesocketappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
)

type PlaceItemCommand struct {
	GameId   gamemodel.GameIdVo
	PlayerId gamemodel.PlayerIdVo
	ItemId   itemmodel.ItemIdVo
	Location commonmodel.LocationVo
}

func NewPlaceItemCommand(gameIdDto string, playerIdDto string, locationDto dto.LocationDto, itemIdDto int16) (PlaceItemCommand, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdDto)
	if err != nil {
		return PlaceItemCommand{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdDto)
	if err != nil {
		return PlaceItemCommand{}, err
	}
	itemId := itemmodel.NewItemIdVo(itemIdDto)
	location := commonmodel.NewLocationVo(locationDto.X, locationDto.Y)

	return PlaceItemCommand{
		GameId:   gameId,
		PlayerId: playerId,
		ItemId:   itemId,
		Location: location,
	}, nil
}

type DestroyItemCommand struct {
	GameId   gamemodel.GameIdVo
	PlayerId gamemodel.PlayerIdVo
	Location commonmodel.LocationVo
}

func NewDestroyItemCommand(gameIdDto string, playerIdDto string, locationDto dto.LocationDto) (DestroyItemCommand, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdDto)
	if err != nil {
		return DestroyItemCommand{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdDto)
	if err != nil {
		return DestroyItemCommand{}, err
	}
	location := commonmodel.NewLocationVo(locationDto.X, locationDto.Y)

	return DestroyItemCommand{
		GameId:   gameId,
		PlayerId: playerId,
		Location: location,
	}, nil
}

type AddPlayerCommand struct {
	GameId   gamemodel.GameIdVo
	PlayerId gamemodel.PlayerIdVo
}

func NewAddPlayerCommand(gameIdDto string, playerIdDto string) (AddPlayerCommand, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdDto)
	if err != nil {
		return AddPlayerCommand{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdDto)
	if err != nil {
		return AddPlayerCommand{}, err
	}

	return AddPlayerCommand{
		GameId:   gameId,
		PlayerId: playerId,
	}, nil
}

type MovePlayerCommand struct {
	GameId    gamemodel.GameIdVo
	PlayerId  gamemodel.PlayerIdVo
	Direction gamemodel.DirectionVo
}

func NewMovePlayerCommand(gameIdDto string, playerIdDto string, directionDto int8) (MovePlayerCommand, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdDto)
	if err != nil {
		return MovePlayerCommand{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdDto)
	if err != nil {
		return MovePlayerCommand{}, err
	}
	direction, err := gamemodel.NewDirectionVo(directionDto)
	if err != nil {
		return MovePlayerCommand{}, err
	}

	return MovePlayerCommand{
		GameId:    gameId,
		PlayerId:  playerId,
		Direction: direction,
	}, nil
}

type RemovePlayerCommand struct {
	GameId   gamemodel.GameIdVo
	PlayerId gamemodel.PlayerIdVo
}

func NewRemovePlayerCommand(gameIdDto string, playerIdDto string) (RemovePlayerCommand, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdDto)
	if err != nil {
		return RemovePlayerCommand{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdDto)
	if err != nil {
		return RemovePlayerCommand{}, err
	}

	return RemovePlayerCommand{
		GameId:   gameId,
		PlayerId: playerId,
	}, nil
}
