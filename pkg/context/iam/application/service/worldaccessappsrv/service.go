package worldaccessappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/samber/lo"
)

type Service interface {
	AddWorldMember(AddWorldMemberCommand) error
	FindWorldMember(FindWorldMemberQuery) (worldMemberDto dto.WorldMemberDto, found bool, err error)
	GetWorldMembers(GetWorldMembersQuery) ([]dto.WorldMemberDto, error)
}

type serve struct {
	worldMemberRepo          worldaccessmodel.WorldMemberRepo
	worldAccessDomainService service.WorldAccessService
}

func NewService(worldMemberRepo worldaccessmodel.WorldMemberRepo, worldAccessDomainService service.WorldAccessService) Service {
	return &serve{
		worldMemberRepo:          worldMemberRepo,
		worldAccessDomainService: worldAccessDomainService,
	}
}

func (serve *serve) AddWorldMember(command AddWorldMemberCommand) error {
	worldRole, err := sharedkernelmodel.NewWorldRole(string(command.Role))
	if err != nil {
		return err
	}
	return serve.worldAccessDomainService.AddWorldMember(
		sharedkernelmodel.NewWorldId(command.WorldId),
		sharedkernelmodel.NewUserId(command.UserId),
		worldRole,
	)
}

func (serve *serve) FindWorldMember(query FindWorldMemberQuery) (worldMemberDto dto.WorldMemberDto, found bool, err error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	userId := sharedkernelmodel.NewUserId(query.UserId)
	worldMember, worldMemberFound, err := serve.worldMemberRepo.FindUserWorldMember(worldId, userId)
	if err != nil {
		return worldMemberDto, found, err
	}
	if !worldMemberFound {
		return worldMemberDto, false, nil
	}
	return dto.NewWorldMemberDto(worldMember), true, nil
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
