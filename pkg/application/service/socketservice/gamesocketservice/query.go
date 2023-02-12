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

type GetViewQuery struct {
	GameId   gamemodel.GameIdVo
	PlayerId gamemodel.PlayerIdVo
}

func NewGetViewQuery(gameIdVm string, playerIdVm string) (GetViewQuery, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return GetViewQuery{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return GetViewQuery{}, err
	}

	return GetViewQuery{
		GameId:   gameId,
		PlayerId: playerId,
	}, nil
}
