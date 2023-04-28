package gamermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type Repo interface {
	FindGamerByUserId(sharedkernelmodel.UserId) (gamer Gamer, gamerFound bool, err error)
	GetGamerByUserId(sharedkernelmodel.UserId) (gamer Gamer, err error)
	Add(Gamer) error
	Get(commonmodel.GamerId) (Gamer, error)
	GetAll() ([]Gamer, error)
}
