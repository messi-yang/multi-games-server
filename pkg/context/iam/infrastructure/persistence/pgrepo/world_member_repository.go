package pgrepo

import (
	"errors"

	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
)

func newWorldMemberModel(worldMember worldaccessmodel.WorldMember) pgmodel.WorldMemberModel {
	return pgmodel.WorldMemberModel{
		Id:        worldMember.GetId().Uuid(),
		WorldId:   worldMember.GeWorldId().Uuid(),
		UserId:    worldMember.GeUserId().Uuid(),
		Role:      pgmodel.WorldRole(worldMember.GetRole().String()),
		CreatedAt: worldMember.GetCreatedAt(),
		UpdatedAt: worldMember.GetUpdatedAt(),
	}
}

func parseWorldMemberModel(worldMemberModel pgmodel.WorldMemberModel) (worldMember worldaccessmodel.WorldMember, err error) {
	worldRole, err := sharedkernelmodel.NewWorldRole(string(worldMemberModel.Role))
	if err != nil {
		return worldMember, err
	}
	return worldaccessmodel.LoadWorldMember(
		worldaccessmodel.NewWorldMemberId(worldMemberModel.Id),
		sharedkernelmodel.NewWorldId(worldMemberModel.WorldId),
		sharedkernelmodel.NewUserId(worldMemberModel.UserId),
		worldRole,
		worldMemberModel.CreatedAt,
		worldMemberModel.UpdatedAt,
	), nil
}

type worldMemberRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewWorldMemberRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository worldaccessmodel.WorldMemberRepo) {
	return &worldMemberRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *worldMemberRepo) Add(worldMember worldaccessmodel.WorldMember) error {
	worldMemberModel := newWorldMemberModel(worldMember)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&worldMemberModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&worldMember)
}

func (repo *worldMemberRepo) Get(worldMemberId worldaccessmodel.WorldMemberId) (worldMember worldaccessmodel.WorldMember, err error) {
	worldMemberModel := pgmodel.WorldMemberModel{Id: worldMemberId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&worldMemberModel).Error
	}); err != nil {
		return worldMember, err
	}
	worldMember, err = parseWorldMemberModel(worldMemberModel)
	if err != nil {
		return worldMember, err
	}

	return worldMember, nil
}

func (repo *worldMemberRepo) Delete(worldMember worldaccessmodel.WorldMember) error {
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Delete(&pgmodel.WorldMemberModel{}, worldMember.GetId().Uuid()).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&worldMember)
}

func (repo *worldMemberRepo) GetWorldMemberOfUser(
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
) (*worldaccessmodel.WorldMember, error) {
	worldMemberModel := pgmodel.WorldMemberModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND user_id = ?",
			worldId.Uuid(),
			userId.Uuid(),
		).First(&worldMemberModel).Error
	}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	worldMember, err := parseWorldMemberModel(worldMemberModel)
	if err != nil {
		return nil, err
	}
	return &worldMember, nil
}

func (repo *worldMemberRepo) GetWorldMembersInWorld(worldId sharedkernelmodel.WorldId) (worldMembers []worldaccessmodel.WorldMember, err error) {
	worldMemberModels := []pgmodel.WorldMemberModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ?",
			worldId.Uuid(),
		).Find(&worldMemberModels, pgmodel.WorldMemberModel{}).Error
	}); err != nil {
		return worldMembers, err
	}

	worldMembers, err = commonutil.MapWithError(worldMemberModels, func(_ int, worldMemberModel pgmodel.WorldMemberModel) (worldMember worldaccessmodel.WorldMember, err error) {
		return parseWorldMemberModel(worldMemberModel)
	})
	if err != nil {
		return worldMembers, err
	}
	return worldMembers, nil
}
