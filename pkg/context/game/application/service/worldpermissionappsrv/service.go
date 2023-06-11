package worldpermissionappsrv

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"

type Service interface {
	CanUpdateWorldInfo(CanUpdateWorldInfoQuery) (bool, error)
}

type serve struct {
}

func NewService() Service {
	return &serve{}
}

func (serve *serve) CanUpdateWorldInfo(query CanUpdateWorldInfoQuery) (bool, error) {
	worldRole, err := sharedkernelmodel.NewWorldRole(query.WorldRole)
	if err != nil {
		return false, err
	}
	return worldRole.CanUpdateWorldInfo(), nil
}
