package gamemodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"

type Repo interface {
	Add(GameAgg) error
	GetByUserId(usermodel.UserIdVo) (*GameAgg, error)
	GetAll() ([]GameAgg, error)

	ReadLockAccess(GameIdVo) (rUnlocker func())
	LockAccess(GameIdVo) (unlocker func())
}
