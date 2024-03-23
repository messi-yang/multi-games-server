package pgrepo

import (
	"errors"

	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
)

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
	worldMemberModel := pgmodel.NewWorldMemberModel(worldMember)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&worldMemberModel).Error
	})
}

func (repo *worldMemberRepo) Get(worldMemberId worldaccessmodel.WorldMemberId) (worldMember worldaccessmodel.WorldMember, err error) {
	worldMemberModel := pgmodel.WorldMemberModel{Id: worldMemberId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&worldMemberModel).Error
	}); err != nil {
		return worldMember, err
	}
	worldMember, err = pgmodel.ParseWorldMemberModel(worldMemberModel)
	if err != nil {
		return worldMember, err
	}

	return worldMember, nil
}

func (repo *worldMemberRepo) Delete(worldMember worldaccessmodel.WorldMember) error {
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Delete(&pgmodel.WorldMemberModel{}, worldMember.GetId().Uuid()).Error
	})
}

func (repo *worldMemberRepo) GetWorldMemberOfUser(
	worldId globalcommonmodel.WorldId,
	userId globalcommonmodel.UserId,
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
	worldMember, err := pgmodel.ParseWorldMemberModel(worldMemberModel)
	if err != nil {
		return nil, err
	}
	return &worldMember, nil
}

func (repo *worldMemberRepo) GetWorldMembersInWorld(worldId globalcommonmodel.WorldId) (worldMembers []worldaccessmodel.WorldMember, err error) {
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
		return pgmodel.ParseWorldMemberModel(worldMemberModel)
	})
	if err != nil {
		return worldMembers, err
	}
	return worldMembers, nil
}
