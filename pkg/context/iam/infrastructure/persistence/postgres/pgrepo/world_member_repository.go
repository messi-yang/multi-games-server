package pgrepo

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
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

func (repo *worldMemberRepo) FindUserWorldMember(
	worldId sharedkernelmodel.WorldId,
	userId sharedkernelmodel.UserId,
) (worldMember worldaccessmodel.WorldMember, found bool, err error) {
	worldMemberModels := []pgmodel.WorldMemberModel{}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND user_id = ?",
			worldId.Uuid(),
			userId.Uuid(),
		).Find(&worldMemberModels).Error
	}); err != nil {
		return worldMember, found, err
	}

	found = len(worldMemberModels) >= 1
	if !found {
		return worldMember, false, nil
	} else {
		worldMember, err = parseWorldMemberModel(worldMemberModels[0])
		if err != nil {
			return worldMember, true, err
		}
		return worldMember, true, nil
	}
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
	fmt.Println(worldId.Uuid(), worldMemberModels)

	worldMembers, err = commonutil.MapWithError(worldMemberModels, func(_ int, worldMemberModel pgmodel.WorldMemberModel) (worldMember worldaccessmodel.WorldMember, err error) {
		return parseWorldMemberModel(worldMemberModel)
	})
	if err != nil {
		return worldMembers, err
	}
	return worldMembers, nil
}
