package worldaccessmodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"

type WorldMemberRepo interface {
	Add(WorldMember) error
	Get(WorldMemberId) (WorldMember, error)
	Delete(WorldMember) error
	GetWorldMemberOfUser(globalcommonmodel.WorldId, globalcommonmodel.UserId) (worldMember *WorldMember, err error)
	GetWorldMembersInWorld(globalcommonmodel.WorldId) ([]WorldMember, error)
}
