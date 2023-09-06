package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

func parseStaticUnitFromUnitModel(unitModel pgmodel.UnitModel) (unit staticunitmodel.StaticUnit, err error) {
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
	return repo.domainEventDispatcher.Dispatch(&staticUnit)
}

func (repo *staticUnitRepo) Get(unitId unitmodel.UnitId) (unit staticunitmodel.StaticUnit, err error) {
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

	return parseStaticUnitFromUnitModel(unitModel)
}

func (repo *staticUnitRepo) Delete(staticUnit staticunitmodel.StaticUnit) error {
	return repo.domainEventDispatcher.Dispatch(&staticUnit)
}
