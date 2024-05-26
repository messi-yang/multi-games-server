package worldmemberappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
)

type Service interface {
	DeleteAllWorldMembersInWorld(DeleteAllWorldMembersInWorldCommand) error
}

type serve struct {
	worldMemberRepo worldaccessmodel.WorldMemberRepo
	userRepo        usermodel.UserRepo
}

func NewService(worldMemberRepo worldaccessmodel.WorldMemberRepo, userRepo usermodel.UserRepo) Service {
	return &serve{
		worldMemberRepo: worldMemberRepo,
		userRepo:        userRepo,
	}
}

func (serve *serve) DeleteAllWorldMembersInWorld(command DeleteAllWorldMembersInWorldCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)

	worldMembersInWorld, err := serve.worldMemberRepo.GetWorldMembersInWorld(worldId)
	if err != nil {
		return err
	}

	for _, worldMember := range worldMembersInWorld {
		if err = serve.worldMemberRepo.Delete(worldMember); err != nil {
			return err
		}
	}
	return nil
}
