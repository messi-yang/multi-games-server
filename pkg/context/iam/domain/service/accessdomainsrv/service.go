package accessdomainsrv

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/worldrolemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type Service interface {
	AssignWorldRole(sharedkernelmodel.WorldId, sharedkernelmodel.UserId, worldrolemodel.WorldRoleName) error
}

type serve struct {
	worldRoleRepo         worldrolemodel.Repo
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewService(
	worldRoleRepo worldrolemodel.Repo,
	domainEventDispatcher domain.DomainEventDispatcher,
) Service {
	return &serve{
		worldRoleRepo:         worldRoleRepo,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (serve *serve) AssignWorldRole(
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
	worldRoleName worldrolemodel.WorldRoleName,
) error {
	newWorldRole := worldrolemodel.NewWorldRole(
		worldId,
		userId,
		worldRoleName,
	)
	if err := serve.worldRoleRepo.Add(newWorldRole); err != nil {
		return err
	}
	if err := serve.domainEventDispatcher.Dispatch(&newWorldRole); err != nil {
		return err
	}
	return nil
}
