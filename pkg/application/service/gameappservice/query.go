package gameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type GetNearbyPlayersQuery struct {
	WorldId  worldmodel.WorldIdVo
	PlayerId playermodel.PlayerIdVo
}

type GetNearbyUnitsQuery struct {
	WorldId  worldmodel.WorldIdVo
	PlayerId playermodel.PlayerIdVo
}
