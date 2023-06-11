package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type WorldAccessService interface {
	AssignWorldRoleToUser(sharedkernelmodel.WorldId, sharedkernelmodel.UserId, sharedkernelmodel.WorldRole) error
}

type worldAccessServe struct {
	userWorldRoleRepo     worldaccessmodel.UserWorldRoleRepo
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewWorldAccessService(
	userWorldRoleRepo worldaccessmodel.UserWorldRoleRepo,
	domainEventDispatcher domain.DomainEventDispatcher,
) WorldAccessService {
	return &worldAccessServe{
		userWorldRoleRepo:     userWorldRoleRepo,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (worldAccessServe *worldAccessServe) AssignWorldRoleToUser(
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
	worldRole sharedkernelmodel.WorldRole,
) error {
	newUserWorldRole := worldaccessmodel.NewUserWorldRole(
		worldId,
		userId,
		worldRole,
	)
	return worldAccessServe.userWorldRoleRepo.Add(newUserWorldRole)
}
