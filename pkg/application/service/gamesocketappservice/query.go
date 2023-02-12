package gamesocketappservice

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"

type GetPlayersQuery struct {
	GameId   gamemodel.GameIdVo
	PlayerId gamemodel.PlayerIdVo
}

func NewGetPlayersQuery(gameIdDto string, playerIdDto string) (GetPlayersQuery, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdDto)
	if err != nil {
		return GetPlayersQuery{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdDto)
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

func NewGetViewQuery(gameIdDto string, playerIdDto string) (GetViewQuery, error) {
	gameId, err := gamemodel.NewGameIdVo(gameIdDto)
	if err != nil {
		return GetViewQuery{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdDto)
	if err != nil {
		return GetViewQuery{}, err
	}

	return GetViewQuery{
		GameId:   gameId,
		PlayerId: playerId,
	}, nil
}
