package worldaccessmodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"

type WorldMemberRepo interface {
	Add(WorldMember) error
	Get(WorldMemberId) (WorldMember, error)
	GetUserWorldMember(sharedkernelmodel.WorldId, sharedkernelmodel.UserId) (worldMember *WorldMember, err error)
	GetWorldMembersInWorld(sharedkernelmodel.WorldId) ([]WorldMember, error)
}
