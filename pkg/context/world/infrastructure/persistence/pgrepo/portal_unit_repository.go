package pgrepo

import (
	"math/rand"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/samber/lo"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

func newPortalUnitModel(portalUnit portalunitmodel.PortalUnit) pgmodel.PortalUnitModel {
	targetPosition := portalUnit.GetTargetPosition()
	return pgmodel.PortalUnitModel{
		WorldId:   portalUnit.GetWorldId().Uuid(),
		PosX:      portalUnit.GetPosition().GetX(),
		PosZ:      portalUnit.GetPosition().GetZ(),
		ItemId:    portalUnit.GetItemId().Uuid(),
		Direction: portalUnit.GetDirection().Int8(),
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

func parsePortalUnitModel(portalUnitModel pgmodel.PortalUnitModel) (unit portalunitmodel.PortalUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(portalUnitModel.WorldId)
	pos := worldcommonmodel.NewPosition(portalUnitModel.PosX, portalUnitModel.PosZ)
	targetPosition := lo.TernaryF(
		portalUnitModel.TargetPosX == nil,
		func() *worldcommonmodel.Position {
			return nil
		},
		func() *worldcommonmodel.Position {
			return commonutil.ToPointer(worldcommonmodel.NewPosition(*portalUnitModel.TargetPosX, *portalUnitModel.TargetPosZ))
		},
	)

	return portalunitmodel.LoadPortalUnit(
		unitmodel.NewUnitId(worldId, pos),
		worldId,
		pos,
		worldcommonmodel.NewItemId(portalUnitModel.ItemId),
		worldcommonmodel.NewDirection(portalUnitModel.Direction),
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
	portalUnitModel := newPortalUnitModel(portalUnit)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&portalUnitModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&portalUnit)
}

func (repo *portalUnitRepo) Get(unitId unitmodel.UnitId) (unit portalunitmodel.PortalUnit, err error) {
	portalUnitModel := pgmodel.PortalUnitModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ?",
			unitId.GetWorldId().Uuid(),
			unitId.GetPosition().GetX(),
			unitId.GetPosition().GetZ(),
		).First(&portalUnitModel).Error
	}); err != nil {
		return unit, err
	}

	return parsePortalUnitModel(portalUnitModel)
}

func (repo *portalUnitRepo) Update(portalUnit portalunitmodel.PortalUnit) error {
	portalUnitModel := newPortalUnitModel(portalUnit)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Model(&pgmodel.PortalUnitModel{}).Where(
			"world_id = ? AND pos_x = ? AND pos_z = ?",
			portalUnit.GetWorldId().Uuid(),
			portalUnit.GetPosition().GetX(),
			portalUnit.GetPosition().GetZ(),
		).Select("*").Updates(portalUnitModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&portalUnit)
}

func (repo *portalUnitRepo) Delete(portalUnit portalunitmodel.PortalUnit) error {
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ?",
			portalUnit.GetWorldId().Uuid(),
			portalUnit.GetPosition().GetX(),
			portalUnit.GetPosition().GetZ(),
		).Delete(&pgmodel.PortalUnitModel{}).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&portalUnit)
}

func (repo *portalUnitRepo) GetFirstPortalUnitWithNoTarget(worldId globalcommonmodel.WorldId) (portalUnit *portalunitmodel.PortalUnit, err error) {
	var portalUnitModels []pgmodel.PortalUnitModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND target_pos_x IS NULL AND target_pos_z IS NULL",
			worldId.Uuid(),
		).Find(&portalUnitModels, pgmodel.PortalUnitModel{}).Error
	}); err != nil {
		return portalUnit, err
	}

	if len(portalUnitModels) == 0 {
		return nil, nil
	}

	randomPortalUnitIndex := rand.Intn(len(portalUnitModels))
	randomPortalUnit, err := parsePortalUnitModel(portalUnitModels[randomPortalUnitIndex])
	if err != nil {
		return nil, err
	}
	return commonutil.ToPointer(randomPortalUnit), err
}
