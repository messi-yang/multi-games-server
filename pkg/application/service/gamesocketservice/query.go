package gamesocketservice

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"

type GetPlayersQuery struct {
	GameId   gamemodel.GameIdVo
	PlayerId gamemodel.PlayerIdVo
}

func NewGetPlayersQuery(gameIdVm string, playerIdVm string) (GetPlayersQuery, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return GetPlayersQuery{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return GetPlayersQuery{}, err
	}

	return GetPlayersQuery{
		GameId:   gameId,
		PlayerId: playerId,
	}, nil
}

type GetPlayerViewQuery struct {
	GameId   gamemodel.GameIdVo
	PlayerId gamemodel.PlayerIdVo
}

func NewGetPlayerViewQuery(gameIdVm string, playerIdVm string) (GetPlayerViewQuery, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return GetPlayerViewQuery{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return GetPlayerViewQuery{}, err
	}

	return GetPlayerViewQuery{
		GameId:   gameId,
		PlayerId: playerId,
	}, nil
}
