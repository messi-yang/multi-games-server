package service

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type WorldAccessService interface {
	AddWorldMember(sharedkernelmodel.WorldId, sharedkernelmodel.UserId, sharedkernelmodel.WorldRole) error
}

type worldAccessServe struct {
	worldMemberRepo       worldaccessmodel.WorldMemberRepo
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewWorldAccessService(
	worldMemberRepo worldaccessmodel.WorldMemberRepo,
	domainEventDispatcher domain.DomainEventDispatcher,
) WorldAccessService {
	return &worldAccessServe{
		worldMemberRepo:       worldMemberRepo,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (worldAccessServe *worldAccessServe) AddWorldMember(
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
	role sharedkernelmodel.WorldRole,
) error {
	newWorldMember := worldaccessmodel.NewWorldMember(
		worldId,
		userId,
		role,
	)
	return worldAccessServe.worldMemberRepo.Add(newWorldMember)
}
