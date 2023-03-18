package gamesocketappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type GetPlayersQuery struct {
	WorldId  string
	PlayerId string
}

func (query GetPlayersQuery) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, err error,
) {
	worldId, err = worldmodel.ParseWorldIdVo(query.WorldId)
	if err != nil {
		return worldId, playerId, err
	}
	playerId, err = playermodel.ParsePlayerIdVo(query.PlayerId)
	if err != nil {
		return worldId, playerId, err
	}

	return worldId, playerId, nil
}

type GetUnitsVisibleByPlayerQuery struct {
	WorldId  string
	PlayerId string
}

func (query GetUnitsVisibleByPlayerQuery) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, err error,
) {
	worldId, err = worldmodel.ParseWorldIdVo(query.WorldId)
	if err != nil {
		return worldId, playerId, err
	}
	playerId, err = playermodel.ParsePlayerIdVo(query.PlayerId)
	if err != nil {
		return worldId, playerId, err
	}

	return worldId, playerId, nil
}
