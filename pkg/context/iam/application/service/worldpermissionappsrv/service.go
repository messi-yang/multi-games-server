package worldpermissionappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type Service interface {
	CanGetWorldMembers(CanGetWorldMembersQuery) (bool, error)
	CanUpdateWorld(CanUpdateWorldQuery) (bool, error)
	CanDeleteWorld(CanDeleteWorldQuery) (bool, error)
}

type serve struct {
	worldMemberRepo worldaccessmodel.WorldMemberRepo
}

func NewService(worldMemberRepo worldaccessmodel.WorldMemberRepo) Service {
	return &serve{
		worldMemberRepo: worldMemberRepo,
	}
}

func (serve *serve) CanGetWorldMembers(query CanGetWorldMembersQuery) (bool, error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	userId := sharedkernelmodel.NewUserId(query.UserId)
	worldMember, err := serve.worldMemberRepo.GetWorldMemberOfUser(worldId, userId)
	if err != nil {
		return false, err
	}

	if worldMember == nil {
		return false, nil
	}

	worldPermission := worldaccessmodel.NewWorldPermission(worldMember.GetRole())
	return worldPermission.CanGetWorldMembers(), nil
}

func (serve *serve) CanUpdateWorld(query CanUpdateWorldQuery) (bool, error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	userId := sharedkernelmodel.NewUserId(query.UserId)
	worldMember, err := serve.worldMemberRepo.GetWorldMemberOfUser(worldId, userId)
	if err != nil {
		return false, err
	}

	if worldMember == nil {
		return false, nil
	}

	worldPermission := worldaccessmodel.NewWorldPermission(worldMember.GetRole())
	return worldPermission.CanUpdateWorld(), nil
}

func (serve *serve) CanDeleteWorld(query CanDeleteWorldQuery) (bool, error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	userId := sharedkernelmodel.NewUserId(query.UserId)
	worldMember, err := serve.worldMemberRepo.GetWorldMemberOfUser(worldId, userId)
	if err != nil {
		return false, err
	}

	if worldMember == nil {
		return false, nil
	}

	worldPermission := worldaccessmodel.NewWorldPermission(worldMember.GetRole())
	return worldPermission.CanDeleteWorld(), nil
}
