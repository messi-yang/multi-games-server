package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/signunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

var (
	errItemIsNotForSignUnit = fmt.Errorf("item is not for sign unit")
)

type SignUnitService interface {
	CreateSignUnit(
		signunitmodel.SignUnitId,
		globalcommonmodel.WorldId,
		worldcommonmodel.ItemId,
		worldcommonmodel.Position,
		worldcommonmodel.Direction,
		string,
	) error
	RotateSignUnit(signunitmodel.SignUnitId) error
	RemoveSignUnit(signunitmodel.SignUnitId) error
}

type signUnitServe struct {
	unitRepo     unitmodel.UnitRepo
	signUnitRepo signunitmodel.SignUnitRepo
	itemRepo     itemmodel.ItemRepo
}

func NewSignUnitService(
	unitRepo unitmodel.UnitRepo,
	signUnitRepo signunitmodel.SignUnitRepo,
	itemRepo itemmodel.ItemRepo,
) SignUnitService {
	return &signUnitServe{
		unitRepo:     unitRepo,
		signUnitRepo: signUnitRepo,
		itemRepo:     itemRepo,
	}
}

func (signUnitServe *signUnitServe) CreateSignUnit(
	id signunitmodel.SignUnitId,
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
	label string,
) error {
	item, err := signUnitServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.GetCompatibleUnitType().IsSign() {
		return errItemIsNotForSignUnit
	}

	if position.IsEqual(worldcommonmodel.NewPosition(0, 0)) {
		return nil
	}

	unit, err := signUnitServe.unitRepo.Find(worldId, position)
	if err != nil {
		return err
	}
	if unit != nil {
		return errPositionAlreadyHasUnit
	}

	newSignUnit := signunitmodel.NewSignUnit(id, worldId, position, itemId, direction, item.GetDimension(), label)
	return signUnitServe.signUnitRepo.Add(newSignUnit)
}

func (signUnitServe *signUnitServe) RotateSignUnit(id signunitmodel.SignUnitId) error {
	unit, err := signUnitServe.signUnitRepo.Get(id)
	if err != nil {
		return err
	}
	unit.Rotate()

	return signUnitServe.signUnitRepo.Update(unit)
}

func (signUnitServe *signUnitServe) RemoveSignUnit(id signunitmodel.SignUnitId) error {
	unit, err := signUnitServe.signUnitRepo.Get(id)
	if err != nil {
		return err
	}

	unit.Delete()
	return signUnitServe.signUnitRepo.Delete(unit)
}
