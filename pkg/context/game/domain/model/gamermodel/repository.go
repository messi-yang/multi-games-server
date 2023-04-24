package gamermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type Repo interface {
	FindGamerByUserId(sharedkernelmodel.UserIdVo) (gamer GamerAgg, gamerFound bool, err error)
	Add(GamerAgg) error
	Get(commonmodel.GamerIdVo) (GamerAgg, error)
	GetAll() ([]GamerAgg, error)
}
