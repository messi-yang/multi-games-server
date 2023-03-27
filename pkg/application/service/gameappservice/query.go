package gameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/google/uuid"
)

type GetPlayersAroundPlayerQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func (query GetPlayersAroundPlayerQuery) Validate() (
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
