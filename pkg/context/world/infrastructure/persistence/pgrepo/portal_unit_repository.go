package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/samber/lo"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

func newPortalUnitModel(portalUnit unitmodel.PortalUnit) pgmodel.PortalUnitModel {
	targetPosition := portalUnit.GetTargetPosition()
	return pgmodel.PortalUnitModel{
		Id: portalUnit.GetId().Uuid(),
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
	}
}

func parsePortalUnitModel(portalUnitModel pgmodel.PortalUnitModel) (unit unitmodel.PortalUnit, err error) {
	portalUnitId := unitmodel.NewPortalUnitId(portalUnitModel.Id)
	targetPosition := lo.TernaryF(
		portalUnitModel.TargetPosX == nil,
		func() *worldcommonmodel.Position {
			return nil
		},
		func() *worldcommonmodel.Position {
			return commonutil.ToPointer(worldcommonmodel.NewPosition(*portalUnitModel.TargetPosX, *portalUnitModel.TargetPosZ))
		},
	)

	return unitmodel.LoadPortalUnit(
		portalUnitId,
		targetPosition,
	), nil
}

type portalUnitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewPortalUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository unitmodel.PortalUnitRepo) {
	return &portalUnitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *portalUnitRepo) Add(unit unitmodel.PortalUnit) error {
	portalUnitModel := newPortalUnitModel(unit)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&portalUnitModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&unit)
}

func (repo *portalUnitRepo) Get(unitId unitmodel.PortalUnitId) (unit unitmodel.PortalUnit, err error) {
	portalUnitModel := pgmodel.PortalUnitModel{Id: unitId.Uuid()}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&portalUnitModel).Error
	}); err != nil {
		return unit, err
	}
	return parsePortalUnitModel(portalUnitModel)
}

func (repo *portalUnitRepo) Delete(unit unitmodel.PortalUnit) error {
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Delete(&pgmodel.PortalUnitModel{}, unit.GetId().Uuid()).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&unit)
}
