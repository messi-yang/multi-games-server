package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
)

var (
	errItemIsNotForPortalUnit = fmt.Errorf("item is not for portal unit")
)

type PortalUnitService interface {
	CreatePortalUnit(
		portalunitmodel.PortalUnitId,
		globalcommonmodel.WorldId,
		worldcommonmodel.ItemId,
		worldcommonmodel.Position,
		worldcommonmodel.Direction,
	) error
	RotatePortalUnit(portalunitmodel.PortalUnitId) error
	RemovePortalUnit(portalunitmodel.PortalUnitId) error
}

type portalUnitServe struct {
	worldRepo      worldmodel.WorldRepo
	unitRepo       unitmodel.UnitRepo
	portalUnitRepo portalunitmodel.PortalUnitRepo
	itemRepo       itemmodel.ItemRepo
}

func NewPortalUnitService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	portalUnitRepo portalunitmodel.PortalUnitRepo,
	itemRepo itemmodel.ItemRepo,
) PortalUnitService {
	return &portalUnitServe{
		worldRepo:      worldRepo,
		unitRepo:       unitRepo,
		portalUnitRepo: portalUnitRepo,
		itemRepo:       itemRepo,
	}
}

func (portalUnitServe *portalUnitServe) CreatePortalUnit(
	id portalunitmodel.PortalUnitId,
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
) error {
	world, err := portalUnitServe.worldRepo.Get(worldId)
	if err != nil {
		return err
	}

	item, err := portalUnitServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.GetCompatibleUnitType().IsPortal() {
		return errItemIsNotForPortalUnit
	}

	if !world.GetBound().CoversPosition(position) {
		return errUnitExceededBoundary
	}

	if position.IsEqual(worldcommonmodel.NewPosition(0, 0)) {
		return nil
	}

	existingUnit, err := portalUnitServe.unitRepo.Find(worldId, position)
	if err != nil {
		return err
	}
	if existingUnit != nil {
		return nil
	}

	portalUnitWithNoTarget, err := portalUnitServe.portalUnitRepo.GetTopLeftMostUnitWithoutTarget(worldId)
	if err != nil {
		return err
	}

	newPortalUnit := portalunitmodel.NewPortalUnit(
		id,
		worldId,
		position,
		itemId,
		direction,
		nil,
	)

	if portalUnitWithNoTarget != nil {
		newPortalUnit.UpdateTargetPosition(commonutil.ToPointer(portalUnitWithNoTarget.GetPosition()))
		portalUnitWithNoTarget.UpdateTargetPosition(&position)
		if err = portalUnitServe.portalUnitRepo.Update(*portalUnitWithNoTarget); err != nil {
			return err
		}
	}

	return portalUnitServe.portalUnitRepo.Add(newPortalUnit)
}

func (portalUnitServe *portalUnitServe) RotatePortalUnit(id portalunitmodel.PortalUnitId) error {
	unit, err := portalUnitServe.portalUnitRepo.Get(id)
	if err != nil {
		return err
	}
	unit.Rotate()

	return portalUnitServe.portalUnitRepo.Update(unit)
}

func (portalUnitServe *portalUnitServe) RemovePortalUnit(id portalunitmodel.PortalUnitId) error {
	unit, err := portalUnitServe.portalUnitRepo.Get(id)
	if err != nil {
		return err
	}

	targetPosition := unit.GetTargetPosition()
	if targetPosition != nil {
		unitAtTargetPosition, err := portalUnitServe.portalUnitRepo.Find(
			unit.GetWorldId(),
			*targetPosition,
		)
		if err != nil {
			return err
		}
		if unitAtTargetPosition != nil {
			unitAtTargetPosition.UpdateTargetPosition(nil)
			if err = portalUnitServe.portalUnitRepo.Update(*unitAtTargetPosition); err != nil {
				return err
			}
		}
	}

	unit.Delete()
	return portalUnitServe.portalUnitRepo.Delete(unit)
}
