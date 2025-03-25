package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

var (
	errItemIsNotForLinkUnit = fmt.Errorf("item is not for link unit")
)

type LinkUnitService interface {
	CreateLinkUnit(
		linkunitmodel.LinkUnitId,
		globalcommonmodel.WorldId,
		worldcommonmodel.ItemId,
		worldcommonmodel.Position,
		worldcommonmodel.Direction,
		*string,
		globalcommonmodel.Url,
	) error
	RotateLinkUnit(linkunitmodel.LinkUnitId) error
	RemoveLinkUnit(linkunitmodel.LinkUnitId) error
}

type linkUnitServe struct {
	unitRepo     unitmodel.UnitRepo
	linkUnitRepo linkunitmodel.LinkUnitRepo
	itemRepo     itemmodel.ItemRepo
}

func NewLinkUnitService(
	unitRepo unitmodel.UnitRepo,
	linkUnitRepo linkunitmodel.LinkUnitRepo,
	itemRepo itemmodel.ItemRepo,
) LinkUnitService {
	return &linkUnitServe{
		unitRepo:     unitRepo,
		linkUnitRepo: linkUnitRepo,
		itemRepo:     itemRepo,
	}
}

func (linkUnitServe *linkUnitServe) CreateLinkUnit(
	id linkunitmodel.LinkUnitId,
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
	label *string,
	url globalcommonmodel.Url,
) error {
	item, err := linkUnitServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.GetCompatibleUnitType().IsLink() {
		return errItemIsNotForLinkUnit
	}

	if position.IsEqual(worldcommonmodel.NewPosition(0, 0)) {
		return nil
	}

	newLinkUnit := linkunitmodel.NewLinkUnit(id, worldId, position, itemId, direction, item.GetDimension(), label, url)

	hasUnitsInBound, err := linkUnitServe.unitRepo.HasUnitsInBound(worldId, newLinkUnit.GetOccupiedBound())
	if err != nil {
		return err
	}
	if hasUnitsInBound {
		return errBoundAlreadyHasUnit
	}

	return linkUnitServe.linkUnitRepo.Add(newLinkUnit)
}

func (linkUnitServe *linkUnitServe) RotateLinkUnit(id linkunitmodel.LinkUnitId) error {
	unit, err := linkUnitServe.linkUnitRepo.Get(id)
	if err != nil {
		return err
	}
	unit.Rotate()

	return linkUnitServe.linkUnitRepo.Update(unit)
}

func (linkUnitServe *linkUnitServe) RemoveLinkUnit(id linkunitmodel.LinkUnitId) error {
	unit, err := linkUnitServe.linkUnitRepo.Get(id)
	if err != nil {
		return err
	}

	unit.Delete()
	return linkUnitServe.linkUnitRepo.Delete(unit)
}
