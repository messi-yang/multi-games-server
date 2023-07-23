package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
	"github.com/samber/lo"
)

func newUnitModel(unit unitmodel.Unit) pgmodel.UnitModel {
	return pgmodel.UnitModel{
		WorldId:   unit.GetWorldId().Uuid(),
		PosX:      unit.GetPosition().GetX(),
		PosZ:      unit.GetPosition().GetZ(),
		ItemId:    unit.GetItemId().Uuid(),
		Direction: unit.GetDirection().Int8(),
	}
}

func parseUnitModel(unitModel pgmodel.UnitModel) unitmodel.Unit {
	worldId := globalcommonmodel.NewWorldId(unitModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	return unitmodel.LoadUnit(
		unitmodel.NewUnitId(worldId, pos),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
	)
}

type unitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository unitmodel.UnitRepo) {
	return &unitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *unitRepo) Add(unit unitmodel.Unit) error {
	unitModel := newUnitModel(unit)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&unitModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&unit)
}

func (repo *unitRepo) Get(unitId unitmodel.UnitId) (unit unitmodel.Unit, err error) {
	unitModel := pgmodel.UnitModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ?",
			unitId.GetWorldId().Uuid(),
			unitId.GetPosition().GetX(),
			unitId.GetPosition().GetZ(),
		).First(&unitModel).Error
	}); err != nil {
		return unit, err
	}
	return parseUnitModel(unitModel), nil
}

func (repo *unitRepo) Delete(unit unitmodel.Unit) error {
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ?",
			unit.GetWorldId().Uuid(),
			unit.GetPosition().GetX(),
			unit.GetPosition().GetZ(),
		).Delete(&pgmodel.UnitModel{}).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&unit)
}

func (repo *unitRepo) GetUnitAt(
	worldId globalcommonmodel.WorldId, position worldcommonmodel.Position,
) (*unitmodel.Unit, error) {
	unitModels := []pgmodel.UnitModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ?",
			worldId.Uuid(),
			position.GetX(),
			position.GetZ(),
		).Limit(1).Find(&unitModels).Error
	}); err != nil {
		return nil, err
	}

	if len(unitModels) == 0 {
		return nil, nil
	}

	return commonutil.ToPointer(parseUnitModel(unitModels[0])), nil
}

func (repo *unitRepo) GetUnitsOfWorld(
	worldId globalcommonmodel.WorldId,
) (units []unitmodel.Unit, err error) {
	var unitModels []pgmodel.UnitModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ?",
			worldId.Uuid(),
		).Find(&unitModels, pgmodel.UnitModel{}).Error
	}); err != nil {
		return units, err
	}

	units = lo.Map(unitModels, func(unitModel pgmodel.UnitModel, _ int) unitmodel.Unit {
		return parseUnitModel(unitModel)
	})
	return units, nil
}
