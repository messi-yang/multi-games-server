package worldpermissionappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
)

type Service interface {
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

func (serve *serve) CanUpdateWorld(query CanUpdateWorldQuery) (bool, error) {
	worldId := globalcommonmodel.NewWorldId(query.WorldId)
	userId := globalcommonmodel.NewUserId(query.UserId)
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
	worldId := globalcommonmodel.NewWorldId(query.WorldId)
	userId := globalcommonmodel.NewUserId(query.UserId)
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
