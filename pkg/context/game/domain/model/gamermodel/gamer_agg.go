package gamermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type GamerAgg struct {
	id     commonmodel.GamerIdVo
	userId sharedkernelmodel.UserIdVo
}

func NewGamerAgg(
	id commonmodel.GamerIdVo,
	userId sharedkernelmodel.UserIdVo,
) GamerAgg {
	return GamerAgg{id: id, userId: userId}
}

func (agg *GamerAgg) GetId() commonmodel.GamerIdVo {
	return agg.id
}

func (agg *GamerAgg) GetUserId() sharedkernelmodel.UserIdVo {
	return agg.userId
}
