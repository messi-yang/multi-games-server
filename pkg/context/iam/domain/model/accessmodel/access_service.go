package accessmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type AccessService interface {
	AssignUserToWorldRole(sharedkernelmodel.WorldId, sharedkernelmodel.UserId, WorldRoleName) error
}

type accessServe struct {
	worldRoleRepo         WorldRoleRepo
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewAccessService(
	worldRoleRepo WorldRoleRepo,
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
	worldRoleName WorldRoleName,
) error {
	newWorldRole := NewWorldRole(
		worldId,
		userId,
		worldRoleName,
	)
	return accessServe.worldRoleRepo.Add(newWorldRole)
}
