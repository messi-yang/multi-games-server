package worldaccessappsrv

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/samber/lo"
)

type Service interface {
	AssignWorldRoleToUser(AssignWorldRoleToUserCommand) error
	FindUserWorldRole(FindUserWorldRoleQuery) (userWorldRoleDto dto.UserWorldRoleDto, found bool, err error)
	GetUserWorldRoles(GetUserWorldRolesQuery) ([]dto.UserWorldRoleDto, error)
}

type serve struct {
	userWorldRoleRepo        worldaccessmodel.UserWorldRoleRepo
	worldAccessDomainService service.WorldAccessService
}

func NewService(userWorldRoleRepo worldaccessmodel.UserWorldRoleRepo, worldAccessDomainService service.WorldAccessService) Service {
	return &serve{
		userWorldRoleRepo:        userWorldRoleRepo,
		worldAccessDomainService: worldAccessDomainService,
	}
}

func (serve *serve) AssignWorldRoleToUser(command AssignWorldRoleToUserCommand) error {
	worldRole, err := worldaccessmodel.NewWorldRole(string(command.WorldRole))
	if err != nil {
		return err
	}
	return serve.worldAccessDomainService.AssignWorldRoleToUser(
		sharedkernelmodel.NewWorldId(command.WorldId),
		sharedkernelmodel.NewUserId(command.UserId),
		worldRole,
	)
}

func (serve *serve) FindUserWorldRole(query FindUserWorldRoleQuery) (userWorldRoleDto dto.UserWorldRoleDto, found bool, err error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	userId := sharedkernelmodel.NewUserId(query.UserId)
	userWorldRole, userWorldRoleFound, err := serve.userWorldRoleRepo.FindWorldRoleOfUser(worldId, userId)
	if err != nil {
		return userWorldRoleDto, found, err
	}
	if !userWorldRoleFound {
		return userWorldRoleDto, false, nil
	}
	return dto.NewUserWorldRoleDto(userWorldRole), true, nil
}

func (serve *serve) GetUserWorldRoles(query GetUserWorldRolesQuery) (userWorldRoleDtos []dto.UserWorldRoleDto, err error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	userWorldRoles, err := serve.userWorldRoleRepo.GetUserWorldRolesInWorld(worldId)
	if err != nil {
		return userWorldRoleDtos, err
	}

	return lo.Map(userWorldRoles, func(userWorldRole worldaccessmodel.UserWorldRole, _ int) dto.UserWorldRoleDto {
		return dto.NewUserWorldRoleDto(userWorldRole)
	}), nil
}
