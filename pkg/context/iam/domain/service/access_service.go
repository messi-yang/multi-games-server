package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/accessmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type AccessService interface {
	AssignWorldRoleToUser(sharedkernelmodel.WorldId, sharedkernelmodel.UserId, accessmodel.WorldRole) error
}

type accessServe struct {
	userWorldRoleRepo     accessmodel.UserWorldRoleRepo
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewAccessService(
	userWorldRoleRepo accessmodel.UserWorldRoleRepo,
	domainEventDispatcher domain.DomainEventDispatcher,
) AccessService {
	return &accessServe{
		userWorldRoleRepo:     userWorldRoleRepo,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (accessServe *accessServe) AssignWorldRoleToUser(
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
	worldRole accessmodel.WorldRole,
) error {
	newUserWorldRole := accessmodel.NewUserWorldRole(
		worldId,
		userId,
		worldRole,
	)
	return accessServe.userWorldRoleRepo.Add(newUserWorldRole)
}
