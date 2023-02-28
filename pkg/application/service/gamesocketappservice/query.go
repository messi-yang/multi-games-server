package gamesocketappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type GetPlayersQuery struct {
	WorldId  string
	PlayerId string
}

func (query GetPlayersQuery) Validate() (worldmodel.WorldIdVo, playermodel.PlayerIdVo, error) {
	worldId, err := worldmodel.NewWorldIdVo(query.WorldId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, err
	}
	playerId, err := playermodel.NewPlayerIdVo(query.PlayerId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, err
	}

	return worldId, playerId, nil
}

type GetUnitsVisibleByPlayerQuery struct {
	WorldId  string
	PlayerId string
}

func (query GetUnitsVisibleByPlayerQuery) Validate() (worldmodel.WorldIdVo, playermodel.PlayerIdVo, error) {
	worldId, err := worldmodel.NewWorldIdVo(query.WorldId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, err
	}
	playerId, err := playermodel.NewPlayerIdVo(query.PlayerId)
	if err != nil {
		return worldmodel.WorldIdVo{}, playermodel.PlayerIdVo{}, err
	}

	return worldId, playerId, nil
}
