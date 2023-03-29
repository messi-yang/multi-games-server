package gameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type FindNearbyPlayersQuery struct {
	WorldId  worldmodel.WorldIdVo
	PlayerId playermodel.PlayerIdVo
}

type FindNearbyUnitsQuery struct {
	WorldId  worldmodel.WorldIdVo
	PlayerId playermodel.PlayerIdVo
}
