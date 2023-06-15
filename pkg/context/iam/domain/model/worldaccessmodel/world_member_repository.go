package worldaccessmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"

type WorldMemberRepo interface {
	Add(WorldMember) error
	Get(WorldMemberId) (WorldMember, error)
	FindUserWorldMember(sharedkernelmodel.WorldId, sharedkernelmodel.UserId) (worldMember WorldMember, found bool, err error)
	GetWorldMembersInWorld(sharedkernelmodel.WorldId) ([]WorldMember, error)
}
