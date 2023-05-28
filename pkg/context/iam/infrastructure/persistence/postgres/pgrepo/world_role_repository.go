package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/worldrolemodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func newWorldRoleModel(worldRole worldrolemodel.WorldRole) pgmodel.WorldRoleModel {
	return pgmodel.WorldRoleModel{
		Id:        worldRole.GetId().Uuid(),
		WorldId:   worldRole.GeWorldId().Uuid(),
		UserId:    worldRole.GeUserId().Uuid(),
		Name:      pgmodel.WorldRoleName(worldRole.GetName().String()),
		CreatedAt: worldRole.GetCreatedAt(),
		UpdatedAt: worldRole.GetUpdatedAt(),
	}
}

func parseWorldRoleModel(worldRoleModel pgmodel.WorldRoleModel) (worldRole worldrolemodel.WorldRole, err error) {
	worldRoleName, err := worldrolemodel.NewWorldRoleName(string(worldRoleModel.Name))
	if err != nil {
		return worldRole, err
	}
	return worldrolemodel.LoadWorldRole(
		worldrolemodel.NewWorldRoleId(worldRoleModel.Id),
		sharedkernelmodel.NewWorldId(worldRoleModel.WorldId),
		sharedkernelmodel.NewUserId(worldRoleModel.UserId),
		worldRoleName,
		worldRoleModel.CreatedAt,
		worldRoleModel.UpdatedAt,
	), nil
}

type worldRoleRepo struct {
	uow pguow.Uow
}

func NewWorldRoleRepo(uow pguow.Uow) (repository worldrolemodel.Repo) {
	return &worldRoleRepo{uow: uow}
}

func (repo *worldRoleRepo) Add(worldRole worldrolemodel.WorldRole) error {
	worldRoleModel := newWorldRoleModel(worldRole)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&worldRoleModel).Error
	})
}

func (repo *worldRoleRepo) Get(worldRoleId worldrolemodel.WorldRoleId) (worldRole worldrolemodel.WorldRole, err error) {
	worldRoleModel := pgmodel.WorldRoleModel{Id: worldRoleId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&worldRoleModel).Error
	}); err != nil {
		return worldRole, err
	}
	worldRole, err = parseWorldRoleModel(worldRoleModel)
	if err != nil {
		return worldRole, err
	}

	return worldRole, nil
}
