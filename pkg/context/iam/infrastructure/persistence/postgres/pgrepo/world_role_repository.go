package pgrepo

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/commonutil"
)

func newUserWorldRoleModel(userWorldRole worldaccessmodel.UserWorldRole) pgmodel.UserWorldRoleModel {
	return pgmodel.UserWorldRoleModel{
		Id:        userWorldRole.GetId().Uuid(),
		WorldId:   userWorldRole.GeWorldId().Uuid(),
		UserId:    userWorldRole.GeUserId().Uuid(),
		WorldRole: pgmodel.WorldRole(userWorldRole.GetWorldRole().String()),
		CreatedAt: userWorldRole.GetCreatedAt(),
		UpdatedAt: userWorldRole.GetUpdatedAt(),
	}
}

func parseUserWorldRoleModel(userWorldRoleModel pgmodel.UserWorldRoleModel) (userWorldRole worldaccessmodel.UserWorldRole, err error) {
	worldRole, err := sharedkernelmodel.NewWorldRole(string(userWorldRoleModel.WorldRole))
	if err != nil {
		return userWorldRole, err
	}
	return worldaccessmodel.LoadWorldRole(
		worldaccessmodel.NewUserWorldRoleId(userWorldRoleModel.Id),
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

func NewUserWorldRoleRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository worldaccessmodel.UserWorldRoleRepo) {
	return &userWorldRoleRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *userWorldRoleRepo) Add(userWorldRole worldaccessmodel.UserWorldRole) error {
	userWorldRoleModel := newUserWorldRoleModel(userWorldRole)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&userWorldRoleModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&userWorldRole)
}

func (repo *userWorldRoleRepo) Get(userWorldRoleId worldaccessmodel.UserWorldRoleId) (userWorldRole worldaccessmodel.UserWorldRole, err error) {
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

func (repo *userWorldRoleRepo) FindWorldRoleOfUser(
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
) (userWorldRole worldaccessmodel.UserWorldRole, found bool, err error) {
	userWorldRoleModels := []pgmodel.UserWorldRoleModel{}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND user_id = ?",
			worldId.Uuid(),
			userId.Uuid(),
		).Find(&userWorldRoleModels).Error
	}); err != nil {
		return userWorldRole, found, err
	}

	found = len(userWorldRoleModels) >= 1
	if !found {
		return userWorldRole, false, nil
	} else {
		userWorldRole, err = parseUserWorldRoleModel(userWorldRoleModels[0])
		if err != nil {
			return userWorldRole, true, err
		}
		return userWorldRole, true, nil
	}
}

func (repo *userWorldRoleRepo) GetUserWorldRolesInWorld(worldId sharedkernelmodel.WorldId) (userWorldRoles []worldaccessmodel.UserWorldRole, err error) {
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

	userWorldRoles, err = commonutil.MapWithError(userWorldRoleModels, func(_ int, userWorldRoleModel pgmodel.UserWorldRoleModel) (userWorldRole worldaccessmodel.UserWorldRole, err error) {
		return parseUserWorldRoleModel(userWorldRoleModel)
	})
	if err != nil {
		return userWorldRoles, err
	}
	return userWorldRoles, nil
}
