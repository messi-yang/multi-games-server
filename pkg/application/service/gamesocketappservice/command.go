package gamesocketappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
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

func NewPlaceItemCommand(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm int16) (PlaceItemCommand, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return PlaceItemCommand{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return PlaceItemCommand{}, err
	}
	itemId := itemmodel.NewItemIdVo(itemIdVm)
	location := commonmodel.NewLocationVo(locationVm.X, locationVm.Y)

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

func NewDestroyItemCommand(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm) (DestroyItemCommand, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return DestroyItemCommand{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return DestroyItemCommand{}, err
	}
	location := commonmodel.NewLocationVo(locationVm.X, locationVm.Y)

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

func NewAddPlayerCommand(gameIdVm string, playerIdVm string) (AddPlayerCommand, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return AddPlayerCommand{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
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

func NewMovePlayerCommand(gameIdVm string, playerIdVm string, directionVm int8) (MovePlayerCommand, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return MovePlayerCommand{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return MovePlayerCommand{}, err
	}
	direction, err := gamemodel.NewDirectionVo(directionVm)
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

func NewRemovePlayerCommand(gameIdVm string, playerIdVm string) (RemovePlayerCommand, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return RemovePlayerCommand{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return RemovePlayerCommand{}, err
	}

	return RemovePlayerCommand{
		GameId:   gameId,
		PlayerId: playerId,
	}, nil
}
