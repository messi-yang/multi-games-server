package worldaccountmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type WorldAccountRepo interface {
	Add(WorldAccount) error
	Update(WorldAccount) error
	Get(WorldAccountId) (WorldAccount, error)
	GetAll() ([]WorldAccount, error)
	FindWorldAccountByUserId(sharedkernelmodel.UserId) (worldAccount WorldAccount, worldAccountFound bool, err error)
	GetWorldAccountOfUser(sharedkernelmodel.UserId) (worldAccount WorldAccount, err error)
}
