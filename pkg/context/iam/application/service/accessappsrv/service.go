package accessappsrv

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/worldrolemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/service/accessdomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type Service interface {
	AssignUserToWorldRole(AssignUserToWorldRoleCommand) error
}

type serve struct {
	accessDomainService accessdomainsrv.Service
}

func NewService(accessDomainService accessdomainsrv.Service) Service {
	return &serve{
		accessDomainService: accessDomainService,
	}
}

func (serve *serve) AssignUserToWorldRole(command AssignUserToWorldRoleCommand) error {
	worldRoleName, err := worldrolemodel.NewWorldRoleName(string(command.WorldRoleName))
	if err != nil {
		return err
	}
	return serve.accessDomainService.AssignUserToWorldRole(
		sharedkernelmodel.NewWorldId(command.WorldId),
		sharedkernelmodel.NewUserId(command.UserId),
		worldRoleName,
	)
}
