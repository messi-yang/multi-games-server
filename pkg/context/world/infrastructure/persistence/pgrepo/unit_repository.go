package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

func parseModelToUnit(unitModel pgmodel.UnitModel) (unit unitmodel.Unit, err error) {
	worldId := globalcommonmodel.NewWorldId(unitModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	unitType, err := worldcommonmodel.NewUnitType(string(unitModel.Type))
	if err != nil {
		return unit, err
	}
	return unitmodel.LoadUnit(
		unitmodel.NewUnitId(unitModel.InfoId),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		unitType,
		unitModel.InfoSnapshot,
		unitModel.InfoId,
	), nil
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

func (repo *unitRepo) Get(id unitmodel.UnitId) (unit unitmodel.Unit, err error) {
	unitModel := pgmodel.UnitModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"info_id = ?",
			id.Uuid(),
		).First(&unitModel).Error
	}); err != nil {
		return unit, err
	}
	return parseModelToUnit(unitModel)
}

func (repo *unitRepo) Find(
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
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

	unit, err := parseModelToUnit(unitModels[0])
	if err != nil {
		return nil, err
	}

	return commonutil.ToPointer(unit), nil
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

	return commonutil.MapWithError(unitModels, func(_ int, unitModel pgmodel.UnitModel) (unitmodel.Unit, error) {
		return parseModelToUnit(unitModel)
	})
}
