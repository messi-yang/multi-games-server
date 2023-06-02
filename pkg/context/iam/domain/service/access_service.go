package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/accessmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type AccessService interface {
	AssignUserToWorldRole(sharedkernelmodel.WorldId, sharedkernelmodel.UserId, accessmodel.WorldRoleName) error
}

type accessServe struct {
	worldRoleRepo         accessmodel.WorldRoleRepo
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewAccessService(
	worldRoleRepo accessmodel.WorldRoleRepo,
	domainEventDispatcher domain.DomainEventDispatcher,
) AccessService {
	return &accessServe{
		worldRoleRepo:         worldRoleRepo,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (accessServe *accessServe) AssignUserToWorldRole(
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
	worldRoleName accessmodel.WorldRoleName,
) error {
	newWorldRole := accessmodel.NewWorldRole(
		worldId,
		userId,
		worldRoleName,
	)
	return accessServe.worldRoleRepo.Add(newWorldRole)
}
