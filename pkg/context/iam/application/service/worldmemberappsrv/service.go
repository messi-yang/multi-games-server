package worldmemberappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/samber/lo"
)

type Service interface {
	AddWorldMember(AddWorldMemberCommand) error
	DeleteAllWorldMembersInWorld(DeleteAllWorldMembersInWorldCommand) error
	GetWorldMemberOfUser(GetWorldMemberOfUserQuery) (worldMemberDto *dto.WorldMemberDto, err error)
	GetWorldMembers(GetWorldMembersQuery) ([]dto.WorldMemberDto, error)
}

type serve struct {
	worldMemberRepo worldaccessmodel.WorldMemberRepo
}

func NewService(worldMemberRepo worldaccessmodel.WorldMemberRepo) Service {
	return &serve{
		worldMemberRepo: worldMemberRepo,
	}
}

func (serve *serve) AddWorldMember(command AddWorldMemberCommand) error {
	worldRole, err := sharedkernelmodel.NewWorldRole(string(command.Role))
	if err != nil {
		return err
	}
	newWorldMember := worldaccessmodel.NewWorldMember(
		sharedkernelmodel.NewWorldId(command.WorldId),
		sharedkernelmodel.NewUserId(command.UserId),
		worldRole,
	)
	return serve.worldMemberRepo.Add(newWorldMember)
}

func (serve *serve) DeleteAllWorldMembersInWorld(command DeleteAllWorldMembersInWorldCommand) error {
	worldId := sharedkernelmodel.NewWorldId(command.WorldId)

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

func (serve *serve) GetWorldMemberOfUser(query GetWorldMemberOfUserQuery) (*dto.WorldMemberDto, error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	userId := sharedkernelmodel.NewUserId(query.UserId)
	worldMember, err := serve.worldMemberRepo.GetWorldMemberOfUser(worldId, userId)
	if err != nil {
		return nil, err
	}
	if worldMember == nil {
		return nil, nil
	}
	worldMemberDto := dto.NewWorldMemberDto(*worldMember)
	return &worldMemberDto, nil
}

func (serve *serve) GetWorldMembers(query GetWorldMembersQuery) (worldMemberDtos []dto.WorldMemberDto, err error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	worldMembers, err := serve.worldMemberRepo.GetWorldMembersInWorld(worldId)
	if err != nil {
		return worldMemberDtos, err
	}

	return lo.Map(worldMembers, func(worldMember worldaccessmodel.WorldMember, _ int) dto.WorldMemberDto {
		return dto.NewWorldMemberDto(worldMember)
	}), nil
}
