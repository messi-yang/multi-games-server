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

func (query GetPlayersAroundPlayerQuery) Parse() (worldmodel.WorldIdVo, playermodel.PlayerIdVo) {
	return worldmodel.NewWorldIdVo(query.WorldId), playermodel.NewPlayerIdVo(query.PlayerId)
}

type GetUnitsVisibleByPlayerQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func (query GetUnitsVisibleByPlayerQuery) Parse() (worldmodel.WorldIdVo, playermodel.PlayerIdVo) {
	return worldmodel.NewWorldIdVo(query.WorldId), playermodel.NewPlayerIdVo(query.PlayerId)
}
