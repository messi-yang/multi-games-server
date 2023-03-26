package gamesocketappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/google/uuid"
)

type GetPlayersQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func (query GetPlayersQuery) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, err error,
) {
	worldId = worldmodel.NewWorldIdVo(query.WorldId)
	playerId = playermodel.NewPlayerIdVo(query.PlayerId)

	return worldId, playerId, nil
}

type GetUnitsVisibleByPlayerQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func (query GetUnitsVisibleByPlayerQuery) Validate() (
	worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo, err error,
) {
	worldId = worldmodel.NewWorldIdVo(query.WorldId)
	playerId = playermodel.NewPlayerIdVo(query.PlayerId)

	return worldId, playerId, nil
}
