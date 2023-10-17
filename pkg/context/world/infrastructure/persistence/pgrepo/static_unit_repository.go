package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/jackc/pgtype"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

func newModelFromStaticUnit(staticUnit staticunitmodel.StaticUnit) pgmodel.UnitModel {
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")
	return pgmodel.UnitModel{
		WorldId:      staticUnit.GetWorldId().Uuid(),
		PosX:         staticUnit.GetPosition().GetX(),
		PosZ:         staticUnit.GetPosition().GetZ(),
		ItemId:       staticUnit.GetItemId().Uuid(),
		Direction:    staticUnit.GetDirection().Int8(),
		Type:         pgmodel.UnitTypeEnumStatic,
		InfoId:       nil,
		InfoSnapshot: unitInfoSnapshotJsonb,
	}
}

func parseModelToStaticUnit(unitModel pgmodel.UnitModel) (staticunitmodel.StaticUnit, error) {
	worldId := globalcommonmodel.NewWorldId(unitModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)

	return staticunitmodel.LoadStaticUnit(
		unitmodel.NewUnitId(worldId, pos),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
	), nil
}

type staticUnitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewStaticUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository staticunitmodel.StaticUnitRepo) {
	return &staticUnitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *staticUnitRepo) Add(staticUnit staticunitmodel.StaticUnit) error {
	unitModel := newModelFromStaticUnit(staticUnit)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&unitModel).Error
	}); err != nil {
		return err
	}

	return repo.domainEventDispatcher.Dispatch(&staticUnit)
}

func (repo *staticUnitRepo) Update(staticUnit staticunitmodel.StaticUnit) error {
	unitModel := newModelFromStaticUnit(staticUnit)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Model(&pgmodel.UnitModel{}).Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			unitModel.WorldId,
			unitModel.PosX,
			unitModel.PosZ,
			pgmodel.UnitTypeEnumStatic,
		).Select("*").Updates(unitModel).Error
	}); err != nil {
		return err
	}

	return repo.domainEventDispatcher.Dispatch(&staticUnit)
}

func (repo *staticUnitRepo) Get(unitId unitmodel.UnitId) (unit staticunitmodel.StaticUnit, err error) {
	unitModel := pgmodel.UnitModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			unitId.GetWorldId().Uuid(),
			unitId.GetPosition().GetX(),
			unitId.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumStatic,
		).First(&unitModel).Error
	}); err != nil {
		return unit, err
	}

	return parseModelToStaticUnit(unitModel)
}

func (repo *staticUnitRepo) Delete(staticUnit staticunitmodel.StaticUnit) error {
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			staticUnit.GetWorldId().Uuid(),
			staticUnit.GetPosition().GetX(),
			staticUnit.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumStatic,
		).Delete(&pgmodel.UnitModel{}).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&staticUnit)
}
