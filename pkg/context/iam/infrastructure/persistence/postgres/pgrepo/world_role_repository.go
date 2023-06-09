package pgrepo

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/accessmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/commonutil"
)

func newUserWorldRoleModel(userWorldRole accessmodel.UserWorldRole) pgmodel.UserWorldRoleModel {
	return pgmodel.UserWorldRoleModel{
		Id:        userWorldRole.GetId().Uuid(),
		WorldId:   userWorldRole.GeWorldId().Uuid(),
		UserId:    userWorldRole.GeUserId().Uuid(),
		WorldRole: pgmodel.WorldRole(userWorldRole.GetWorldRole().String()),
		CreatedAt: userWorldRole.GetCreatedAt(),
		UpdatedAt: userWorldRole.GetUpdatedAt(),
	}
}

func parseUserWorldRoleModel(userWorldRoleModel pgmodel.UserWorldRoleModel) (userWorldRole accessmodel.UserWorldRole, err error) {
	worldRole, err := accessmodel.NewWorldRole(string(userWorldRoleModel.WorldRole))
	if err != nil {
		return userWorldRole, err
	}
	return accessmodel.LoadWorldRole(
		accessmodel.NewUserWorldRoleId(userWorldRoleModel.Id),
		sharedkernelmodel.NewWorldId(userWorldRoleModel.WorldId),
		sharedkernelmodel.NewUserId(userWorldRoleModel.UserId),
		worldRole,
		userWorldRoleModel.CreatedAt,
		userWorldRoleModel.UpdatedAt,
	), nil
}

type userWorldRoleRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewUserWorldRoleRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository accessmodel.UserWorldRoleRepo) {
	return &userWorldRoleRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *userWorldRoleRepo) Add(userWorldRole accessmodel.UserWorldRole) error {
	userWorldRoleModel := newUserWorldRoleModel(userWorldRole)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&userWorldRoleModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&userWorldRole)
}

func (repo *userWorldRoleRepo) Get(userWorldRoleId accessmodel.UserWorldRoleId) (userWorldRole accessmodel.UserWorldRole, err error) {
	userWorldRoleModel := pgmodel.UserWorldRoleModel{Id: userWorldRoleId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&userWorldRoleModel).Error
	}); err != nil {
		return userWorldRole, err
	}
	userWorldRole, err = parseUserWorldRoleModel(userWorldRoleModel)
	if err != nil {
		return userWorldRole, err
	}

	return userWorldRole, nil
}

func (repo *userWorldRoleRepo) GetUserWorldRolesInWorld(worldId sharedkernelmodel.WorldId) (userWorldRoles []accessmodel.UserWorldRole, err error) {
	userWorldRoleModels := []pgmodel.UserWorldRoleModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ?",
			worldId.Uuid(),
		).Find(&userWorldRoleModels, pgmodel.UserWorldRoleModel{}).Error
	}); err != nil {
		return userWorldRoles, err
	}
	fmt.Println(worldId.Uuid(), userWorldRoleModels)

	userWorldRoles, err = commonutil.MapWithError(userWorldRoleModels, func(_ int, userWorldRoleModel pgmodel.UserWorldRoleModel) (userWorldRole accessmodel.UserWorldRole, err error) {
		return parseUserWorldRoleModel(userWorldRoleModel)
	})
	if err != nil {
		return userWorldRoles, err
	}
	return userWorldRoles, nil
}
