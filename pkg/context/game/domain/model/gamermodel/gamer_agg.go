package gamermodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"

type GamerAgg struct {
	id     commonmodel.GamerIdVo
	userId commonmodel.UserIdVo
}

func NewGamerAgg(
	id commonmodel.GamerIdVo,
	userId commonmodel.UserIdVo,
) GamerAgg {
	return GamerAgg{id: id, userId: userId}
}

func (agg *GamerAgg) GetId() commonmodel.GamerIdVo {
	return agg.id
}

func (agg *GamerAgg) GetUserId() commonmodel.UserIdVo {
	return agg.userId
}
