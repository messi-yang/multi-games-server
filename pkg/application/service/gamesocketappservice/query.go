package gamesocketappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
)

type GetPlayersQuery struct {
	GameId   string
	PlayerId string
}

func (query GetPlayersQuery) Validate() (gamemodel.GameIdVo, playermodel.PlayerIdVo, error) {
	gameId, err := gamemodel.NewGameIdVo(query.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, err
	}
	playerId, err := playermodel.NewPlayerIdVo(query.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, err
	}

	return gameId, playerId, nil
}

type GetUnitsInBoundAroundPlayerQuery struct {
	GameId   string
	PlayerId string
}

func (query GetUnitsInBoundAroundPlayerQuery) Validate() (gamemodel.GameIdVo, playermodel.PlayerIdVo, error) {
	gameId, err := gamemodel.NewGameIdVo(query.GameId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, err
	}
	playerId, err := playermodel.NewPlayerIdVo(query.PlayerId)
	if err != nil {
		return gamemodel.GameIdVo{}, playermodel.PlayerIdVo{}, err
	}

	return gameId, playerId, nil
}
