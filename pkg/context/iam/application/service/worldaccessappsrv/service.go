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
	GetUserWorldMember(GetUserWorldMemberQuery) (worldMemberDto *dto.WorldMemberDto, err error)
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

func (serve *serve) GetUserWorldMember(query GetUserWorldMemberQuery) (*dto.WorldMemberDto, error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	userId := sharedkernelmodel.NewUserId(query.UserId)
	worldMember, err := serve.worldMemberRepo.GetUserWorldMember(worldId, userId)
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
