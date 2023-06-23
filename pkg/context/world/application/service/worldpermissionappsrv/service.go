package worldpermissionappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
)

type Service interface {
	CanUpdateWorldInfo(CanUpdateWorldInfoQuery) (bool, error)
}

type serve struct {
}

func NewService() Service {
	return &serve{}
}

func (serve *serve) CanUpdateWorldInfo(query CanUpdateWorldInfoQuery) (bool, error) {
	if query.Role == nil {
		return false, nil
	}
	worldRole, err := sharedkernelmodel.NewWorldRole(*query.Role)
	if err != nil {
		return false, err
	}
	worldPermission := worldmodel.NewWorldPermission(worldRole)
	return worldPermission.CanUpdateWorldInfo(), nil
}
