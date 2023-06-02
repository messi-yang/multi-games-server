package gamermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type GamerRepo interface {
	Add(Gamer) error
	Update(Gamer) error
	Get(GamerId) (Gamer, error)
	GetAll() ([]Gamer, error)
	FindGamerByUserId(sharedkernelmodel.UserId) (gamer Gamer, gamerFound bool, err error)
	GetGamerByUserId(sharedkernelmodel.UserId) (gamer Gamer, err error)
}
