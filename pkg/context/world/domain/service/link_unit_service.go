package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
)

var (
	errItemIsNotForLinkUnit = fmt.Errorf("item is not for link unit")
)

type LinkUnitService interface {
	CreateLinkUnit(
		globalcommonmodel.WorldId,
		worldcommonmodel.ItemId,
		worldcommonmodel.Position,
		worldcommonmodel.Direction,
		globalcommonmodel.Url,
	) error
	RotateLinkUnit(unitmodel.UnitId) error
	RemoveLinkUnit(unitmodel.UnitId) error
}

type linkUnitServe struct {
	worldRepo    worldmodel.WorldRepo
	unitRepo     unitmodel.UnitRepo
	linkUnitRepo linkunitmodel.LinkUnitRepo
	itemRepo     itemmodel.ItemRepo
}

func NewLinkUnitService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	linkUnitRepo linkunitmodel.LinkUnitRepo,
	itemRepo itemmodel.ItemRepo,
) LinkUnitService {
	return &linkUnitServe{
		worldRepo:    worldRepo,
		unitRepo:     unitRepo,
		linkUnitRepo: linkUnitRepo,
		itemRepo:     itemRepo,
	}
}

func (linkUnitServe *linkUnitServe) CreateLinkUnit(
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
	url globalcommonmodel.Url,
) error {
	world, err := linkUnitServe.worldRepo.Get(worldId)
	if err != nil {
		return err
	}

	item, err := linkUnitServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.GetCompatibleUnitType().IsLink() {
		return errItemIsNotForLinkUnit
	}

	if !world.GetBound().CoversPosition(position) {
		return errUnitExceededBoundary
	}

	if position.IsEqual(worldcommonmodel.NewPosition(0, 0)) {
		return nil
	}

	unit, err := linkUnitServe.unitRepo.Find(unitmodel.NewUnitId(worldId, position))
	if err != nil {
		return err
	}
	if unit != nil {
		return nil
	}

	newLinkUnit := linkunitmodel.NewLinkUnit(worldId, position, itemId, direction, url)
	return linkUnitServe.linkUnitRepo.Add(newLinkUnit)
}

func (linkUnitServe *linkUnitServe) RotateLinkUnit(unitId unitmodel.UnitId) error {
	unit, err := linkUnitServe.linkUnitRepo.Get(unitId)
	if err != nil {
		return err
	}
	unit.Rotate()

	return linkUnitServe.linkUnitRepo.Update(unit)
}

func (linkUnitServe *linkUnitServe) RemoveLinkUnit(unitId unitmodel.UnitId) error {
	unit, err := linkUnitServe.linkUnitRepo.Get(unitId)
	if err != nil {
		return err
	}

	unit.Delete()
	return linkUnitServe.linkUnitRepo.Delete(unit)
}
