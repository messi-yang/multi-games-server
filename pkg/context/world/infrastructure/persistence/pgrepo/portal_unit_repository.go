package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/jackc/pgtype"
	"github.com/samber/lo"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

func newModelsFromPortalUnit(portalUnit portalunitmodel.PortalUnit) (pgmodel.PortalUnitInfoModel, pgmodel.UnitModel) {
	targetPosition := portalUnit.GetTargetPosition()
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set(portalUnit.GetInfoSnapshot())

	return pgmodel.PortalUnitInfoModel{
			Id:      portalUnit.GetId().Uuid(),
			WorldId: portalUnit.GetWorldId().Uuid(),
			TargetPosX: lo.TernaryF(
				targetPosition == nil,
				func() *int { return nil },
				func() *int { return commonutil.ToPointer(targetPosition.GetX()) },
			),
			TargetPosZ: lo.TernaryF(
				targetPosition == nil,
				func() *int { return nil },
				func() *int { return commonutil.ToPointer(targetPosition.GetZ()) },
			),
		},
		pgmodel.UnitModel{
			Id:           portalUnit.GetId().Uuid(),
			WorldId:      portalUnit.GetWorldId().Uuid(),
			PosX:         portalUnit.GetPosition().GetX(),
			PosZ:         portalUnit.GetPosition().GetZ(),
			ItemId:       portalUnit.GetItemId().Uuid(),
			Direction:    portalUnit.GetDirection().Int8(),
			Type:         pgmodel.UnitTypeEnumPortal,
			InfoSnapshot: unitInfoSnapshotJsonb,
		}
}

func parseModelsToPortalUnit(unitModel pgmodel.UnitModel, portalUnitInfoModel pgmodel.PortalUnitInfoModel) (unit portalunitmodel.PortalUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(portalUnitInfoModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	targetPosition := lo.TernaryF(
		portalUnitInfoModel.TargetPosX == nil,
		func() *worldcommonmodel.Position {
			return nil
		},
		func() *worldcommonmodel.Position {
			return commonutil.ToPointer(worldcommonmodel.NewPosition(*portalUnitInfoModel.TargetPosX, *portalUnitInfoModel.TargetPosZ))
		},
	)

	return portalunitmodel.LoadPortalUnit(
		portalunitmodel.NewPortalUnitId(portalUnitInfoModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		targetPosition,
	), nil
}

type portalUnitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewPortalUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository portalunitmodel.PortalUnitRepo) {
	return &portalUnitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *portalUnitRepo) Add(portalUnit portalunitmodel.PortalUnit) error {
	portalUnitInfoModel, unitModel := newModelsFromPortalUnit(portalUnit)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Create(&portalUnitInfoModel).Error; err != nil {
			return err
		}
		return transaction.Create(&unitModel).Error
	})
}

func (repo *portalUnitRepo) Get(id portalunitmodel.PortalUnitId) (unit portalunitmodel.PortalUnit, err error) {
	unitModel := pgmodel.UnitModel{}
	portalUnitInfoModel := pgmodel.PortalUnitInfoModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"id = ? AND type = ?",
			id.Uuid(),
			pgmodel.UnitTypeEnumPortal,
		).First(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			unitModel.Id,
		).First(&portalUnitInfoModel).Error
	}); err != nil {
		return unit, err
	}

	return parseModelsToPortalUnit(unitModel, portalUnitInfoModel)
}

func (repo *portalUnitRepo) Find(
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
) (*portalunitmodel.PortalUnit, error) {
	unitModels := []pgmodel.UnitModel{}
	unitModel := pgmodel.UnitModel{}
	portalUnitInfoModel := pgmodel.PortalUnitInfoModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ?",
			worldId.Uuid(),
			position.GetX(),
			position.GetZ(),
		).Limit(1).Find(&unitModels).Error; err != nil {
			return err
		}

		if len(unitModels) == 0 {
			return nil
		}

		unitModel = unitModels[0]

		return transaction.Where(
			"id = ?",
			unitModel.Id,
		).First(&portalUnitInfoModel).Error
	}); err != nil {
		return nil, err
	}

	portalUnit, err := parseModelsToPortalUnit(unitModel, portalUnitInfoModel)
	if err != nil {
		return nil, err
	}

	return &portalUnit, nil
}

func (repo *portalUnitRepo) Update(portalUnit portalunitmodel.PortalUnit) error {
	portalUnitInfoModel, unitModel := newModelsFromPortalUnit(portalUnit)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Model(&pgmodel.UnitModel{}).Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			portalUnit.GetWorldId().Uuid(),
			portalUnit.GetPosition().GetX(),
			portalUnit.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumPortal,
		).Select("*").Updates(unitModel).Error; err != nil {
			return err
		}
		return transaction.Model(&pgmodel.PortalUnitInfoModel{}).Where(
			"id = ?",
			portalUnit.GetId().Uuid(),
		).Select("*").Updates(portalUnitInfoModel).Error
	})
}

func (repo *portalUnitRepo) Delete(portalUnit portalunitmodel.PortalUnit) error {
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			portalUnit.GetWorldId().Uuid(),
			portalUnit.GetPosition().GetX(),
			portalUnit.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumPortal,
		).Delete(&pgmodel.UnitModel{}).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			portalUnit.GetId().Uuid(),
		).Delete(&pgmodel.PortalUnitInfoModel{}).Error
	})
}

func (repo *portalUnitRepo) GetTopLeftMostUnitWithoutTarget(worldId globalcommonmodel.WorldId) (portalUnit *portalunitmodel.PortalUnit, err error) {
	var portalUnitInfoModels []pgmodel.PortalUnitInfoModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Select("units.pos_x, units.pos_z, portal_unit_infos.*").Joins(
			"left join units on units.id = portal_unit_infos.id",
		).Where(
			"portal_unit_infos.world_id = ? AND portal_unit_infos.target_pos_x IS NULL AND portal_unit_infos.target_pos_z IS NULL",
			worldId.Uuid(),
		).Order("units.pos_z asc, units.pos_x asc").Limit(1).Find(&portalUnitInfoModels).Error
	}); err != nil {
		return portalUnit, err
	}

	if len(portalUnitInfoModels) == 0 {
		return nil, nil
	}

	var unitModel pgmodel.UnitModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"id = ?",
			portalUnitInfoModels[0].Id,
		).First(&unitModel).Error
	}); err != nil {
		return portalUnit, err
	}

	firstPortalUnitWithNoTarget, err := parseModelsToPortalUnit(unitModel, portalUnitInfoModels[0])
	if err != nil {
		return nil, err
	}
	return commonutil.ToPointer(firstPortalUnitWithNoTarget), err
}
