package gameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/google/uuid"
)

type FindNearbyPlayersQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func (query FindNearbyPlayersQuery) Parse() (worldmodel.WorldIdVo, playermodel.PlayerIdVo) {
	return worldmodel.NewWorldIdVo(query.WorldId), playermodel.NewPlayerIdVo(query.PlayerId)
}

type FindUnitsQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func (query FindUnitsQuery) Parse() (worldmodel.WorldIdVo, playermodel.PlayerIdVo) {
	return worldmodel.NewWorldIdVo(query.WorldId), playermodel.NewPlayerIdVo(query.PlayerId)
}
