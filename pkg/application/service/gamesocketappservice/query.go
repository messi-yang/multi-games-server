package gamesocketappservice

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"

type GetPlayersQuery struct {
	GameId   string
	PlayerId string
}

func (query GetPlayersQuery) Validate() (gamemodel.GameIdVo, gamemodel.PlayerIdVo, error) {
	gameId, err := gamemodel.NewGameIdVo(query.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(query.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, err
	}

	return gameId, playerId, nil
}

type GetViewQuery struct {
	GameId   string
	PlayerId string
}

func (query GetViewQuery) Validate() (gamemodel.GameIdVo, gamemodel.PlayerIdVo, error) {
	gameId, err := gamemodel.NewGameIdVo(query.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, err
	}
	playerId, err := gamemodel.NewPlayerIdVo(query.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, gamemodel.PlayerIdVo{}, err
	}

	return gameId, playerId, nil
}
