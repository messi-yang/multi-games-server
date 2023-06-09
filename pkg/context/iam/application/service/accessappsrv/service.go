package accessappsrv

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/accessmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/samber/lo"
)

type Service interface {
	AssignWorldRoleToUser(AssignWorldRoleToUserCommand) error
	GetUserWorldRoles(GetUserWorldRolesQuery) ([]dto.UserWorldRoleDto, error)
}

type serve struct {
	userWorldRoleRepo   accessmodel.UserWorldRoleRepo
	accessDomainService service.AccessService
}

func NewService(userWorldRoleRepo accessmodel.UserWorldRoleRepo, accessDomainService service.AccessService) Service {
	return &serve{
		userWorldRoleRepo:   userWorldRoleRepo,
		accessDomainService: accessDomainService,
	}
}

func (serve *serve) AssignWorldRoleToUser(command AssignWorldRoleToUserCommand) error {
	worldRole, err := accessmodel.NewWorldRole(string(command.WorldRole))
	if err != nil {
		return err
	}
	return serve.accessDomainService.AssignWorldRoleToUser(
		sharedkernelmodel.NewWorldId(command.WorldId),
		sharedkernelmodel.NewUserId(command.UserId),
		worldRole,
	)
}

func (serve *serve) GetUserWorldRoles(query GetUserWorldRolesQuery) (userWorldRoleDtos []dto.UserWorldRoleDto, err error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	userWorldRoles, err := serve.userWorldRoleRepo.GetUserWorldRolesInWorld(worldId)
	if err != nil {
		return userWorldRoleDtos, err
	}

	return lo.Map(userWorldRoles, func(userWorldRole accessmodel.UserWorldRole, _ int) dto.UserWorldRoleDto {
		return dto.NewUserWorldRoleDto(userWorldRole)
	}), nil
}
