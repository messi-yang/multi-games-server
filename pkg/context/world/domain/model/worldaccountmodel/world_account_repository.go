package worldaccountmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type WorldAccountRepo interface {
	Add(WorldAccount) error
	Update(WorldAccount) error
	Get(WorldAccountId) (WorldAccount, error)
	GetAll() ([]WorldAccount, error)
	GetWorldAccountByUserId(globalcommonmodel.UserId) (worldAccount *WorldAccount, err error)
	GetWorldAccountOfUser(globalcommonmodel.UserId) (worldAccount WorldAccount, err error)
}
