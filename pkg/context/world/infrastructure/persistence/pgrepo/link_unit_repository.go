package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/jackc/pgtype"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

func newModelsFromLinkUnit(linkUnit linkunitmodel.LinkUnit) (pgmodel.LinkUnitInfoModel, pgmodel.UnitModel) {
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")

	return pgmodel.LinkUnitInfoModel{
			Id:      linkUnit.GetId().Uuid(),
			WorldId: linkUnit.GetWorldId().Uuid(),
			Url:     linkUnit.GetUrl().String(),
		},
		pgmodel.UnitModel{
			WorldId:      linkUnit.GetWorldId().Uuid(),
			PosX:         linkUnit.GetPosition().GetX(),
			PosZ:         linkUnit.GetPosition().GetZ(),
			ItemId:       linkUnit.GetItemId().Uuid(),
			Direction:    linkUnit.GetDirection().Int8(),
			Type:         pgmodel.UnitTypeEnumLink,
			Id:           linkUnit.GetId().Uuid(),
			InfoSnapshot: unitInfoSnapshotJsonb,
		}
}

func parseModelsToLinkUnit(unitModel pgmodel.UnitModel, linkUnitInfoModel pgmodel.LinkUnitInfoModel) (unit linkunitmodel.LinkUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(linkUnitInfoModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	url, err := globalcommonmodel.NewUrl(linkUnitInfoModel.Url)
	if err != nil {
		return unit, err
	}

	return linkunitmodel.LoadLinkUnit(
		linkunitmodel.NewLinkUnitId(linkUnitInfoModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		url,
	), nil
}

type linkUnitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

// Interface Implementation Check
var _ linkunitmodel.LinkUnitRepo = (*linkUnitRepo)(nil)

func NewLinkUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository linkunitmodel.LinkUnitRepo) {
	return &linkUnitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *linkUnitRepo) Add(linkUnit linkunitmodel.LinkUnit) error {
	linkUnitInfoModel, unitModel := newModelsFromLinkUnit(linkUnit)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Create(&linkUnitInfoModel).Error; err != nil {
			return err
		}
		return transaction.Create(&unitModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&linkUnit)
}

func (repo *linkUnitRepo) Get(id linkunitmodel.LinkUnitId) (unit linkunitmodel.LinkUnit, err error) {
	unitModel := pgmodel.UnitModel{}
	linkUnitInfoModel := pgmodel.LinkUnitInfoModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"id = ? AND type = ?",
			id.Uuid(),
			pgmodel.UnitTypeEnumLink,
		).First(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			unitModel.Id,
		).First(&linkUnitInfoModel).Error
	}); err != nil {
		return unit, err
	}

	return parseModelsToLinkUnit(unitModel, linkUnitInfoModel)
}

func (repo *linkUnitRepo) Update(linkUnit linkunitmodel.LinkUnit) error {
	linkUnitInfoModel, unitModel := newModelsFromLinkUnit(linkUnit)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Model(&pgmodel.UnitModel{}).Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			linkUnit.GetWorldId().Uuid(),
			linkUnit.GetPosition().GetX(),
			linkUnit.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumLink,
		).Select("*").Updates(unitModel).Error; err != nil {
			return err
		}
		return transaction.Model(&pgmodel.LinkUnitInfoModel{}).Where(
			"id = ?",
			linkUnit.GetId().Uuid(),
		).Select("*").Updates(linkUnitInfoModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&linkUnit)
}

func (repo *linkUnitRepo) Delete(linkUnit linkunitmodel.LinkUnit) error {
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			linkUnit.GetWorldId().Uuid(),
			linkUnit.GetPosition().GetX(),
			linkUnit.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumLink,
		).Delete(&pgmodel.UnitModel{}).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			linkUnit.GetId().Uuid(),
		).Delete(&pgmodel.LinkUnitInfoModel{}).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&linkUnit)
}
