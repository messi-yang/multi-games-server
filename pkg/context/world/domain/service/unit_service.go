package service

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type UnitService interface {
	MoveUnit(unitId unitmodel.UnitId, position worldcommonmodel.Position) error
}

type unitServe struct {
	unitRepo unitmodel.UnitRepo
}

func NewUnitService(unitRepo unitmodel.UnitRepo) UnitService {
	return &unitServe{unitRepo: unitRepo}
}

func (u *unitServe) MoveUnit(unitId unitmodel.UnitId, position worldcommonmodel.Position) error {
	unit, err := u.unitRepo.Get(unitId)
	if err != nil {
		return err
	}

	unit.Move(position)
	return u.unitRepo.Update(unit)
}
