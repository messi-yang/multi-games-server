package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
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
	unitRepo       unitmodel.UnitRepo
	portalUnitRepo portalunitmodel.PortalUnitRepo
	itemRepo       itemmodel.ItemRepo
}

func NewPortalUnitService(
	unitRepo unitmodel.UnitRepo,
	portalUnitRepo portalunitmodel.PortalUnitRepo,
	itemRepo itemmodel.ItemRepo,
) PortalUnitService {
	return &portalUnitServe{
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
	item, err := portalUnitServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.GetCompatibleUnitType().IsPortal() {
		return errItemIsNotForPortalUnit
	}

	if position.IsEqual(worldcommonmodel.NewPosition(0, 0)) {
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
		item.GetDimension(),
		nil,
	)

	hasUnitsInBound, err := portalUnitServe.unitRepo.HasUnitsInBound(worldId, newPortalUnit.GetOccupiedBound())
	if err != nil {
		return err
	}
	if hasUnitsInBound {
		return errBoundAlreadyHasUnit
	}

	if portalUnitWithNoTarget != nil {
		newPortalUnit.UpdateTargetUnitId(commonutil.ToPointer(portalUnitWithNoTarget.GetId()))
		portalUnitWithNoTarget.UpdateTargetUnitId(&id)
		if err = portalUnitServe.portalUnitRepo.Add(newPortalUnit); err != nil {
			return err
		}
		return portalUnitServe.portalUnitRepo.Update(*portalUnitWithNoTarget)
	} else {
		return portalUnitServe.portalUnitRepo.Add(newPortalUnit)
	}

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

	targetUnitId := unit.GetTargetUnitId()
	if targetUnitId != nil {
		targetUnit, err := portalUnitServe.portalUnitRepo.Get(*targetUnitId)
		if err != nil {
			return err
		}
		targetUnit.UpdateTargetUnitId(nil)
		if err = portalUnitServe.portalUnitRepo.Update(targetUnit); err != nil {
			return err
		}
	}

	unit.Delete()
	return portalUnitServe.portalUnitRepo.Delete(unit)
}
