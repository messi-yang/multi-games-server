package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/embedunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

var (
	errItemIsNotForEmbedUnit = fmt.Errorf("item is not for embed unit")
)

type EmbedUnitService interface {
	CreateEmbedUnit(
		embedunitmodel.EmbedUnitId,
		globalcommonmodel.WorldId,
		worldcommonmodel.ItemId,
		worldcommonmodel.Position,
		worldcommonmodel.Direction,
		*string,
		worldcommonmodel.EmbedCode,
	) error
	RotateEmbedUnit(embedunitmodel.EmbedUnitId) error
	RemoveEmbedUnit(embedunitmodel.EmbedUnitId) error
}

type embedUnitServe struct {
	unitRepo      unitmodel.UnitRepo
	embedUnitRepo embedunitmodel.EmbedUnitRepo
	itemRepo      itemmodel.ItemRepo
}

func NewEmbedUnitService(
	unitRepo unitmodel.UnitRepo,
	embedUnitRepo embedunitmodel.EmbedUnitRepo,
	itemRepo itemmodel.ItemRepo,
) EmbedUnitService {
	return &embedUnitServe{
		unitRepo:      unitRepo,
		embedUnitRepo: embedUnitRepo,
		itemRepo:      itemRepo,
	}
}

func (embedUnitServe *embedUnitServe) CreateEmbedUnit(
	id embedunitmodel.EmbedUnitId,
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
	label *string,
	embedCode worldcommonmodel.EmbedCode,
) error {
	item, err := embedUnitServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.GetCompatibleUnitType().IsEmbed() {
		return errItemIsNotForEmbedUnit
	}

	if position.IsEqual(worldcommonmodel.NewPosition(0, 0)) {
		return nil
	}

	unit, err := embedUnitServe.unitRepo.Find(worldId, position)
	if err != nil {
		return err
	}
	if unit != nil {
		return errPositionAlreadyHasUnit
	}

	newEmbedUnit := embedunitmodel.NewEmbedUnit(id, worldId, position, itemId, direction, item.GetDimension(), label, embedCode)
	return embedUnitServe.embedUnitRepo.Add(newEmbedUnit)
}

func (embedUnitServe *embedUnitServe) RotateEmbedUnit(id embedunitmodel.EmbedUnitId) error {
	unit, err := embedUnitServe.embedUnitRepo.Get(id)
	if err != nil {
		return err
	}
	unit.Rotate()

	return embedUnitServe.embedUnitRepo.Update(unit)
}

func (embedUnitServe *embedUnitServe) RemoveEmbedUnit(id embedunitmodel.EmbedUnitId) error {
	unit, err := embedUnitServe.embedUnitRepo.Get(id)
	if err != nil {
		return err
	}

	unit.Delete()
	return embedUnitServe.embedUnitRepo.Delete(unit)
}
