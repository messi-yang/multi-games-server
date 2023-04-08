package gamermodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"

type GamerAgg struct {
	id     GamerIdVo
	userId commonmodel.UserIdVo
}

func NewGamerAgg(
	id GamerIdVo,
	userId commonmodel.UserIdVo,
) GamerAgg {
	return GamerAgg{id: id, userId: userId}
}

func (agg *GamerAgg) GetId() GamerIdVo {
	return agg.id
}

func (agg *GamerAgg) GetUserId() commonmodel.UserIdVo {
	return agg.userId
}
