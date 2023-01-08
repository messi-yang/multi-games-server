package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
)

type UnitMap [][]Unit

func NewUnitMap(unitMap commonmodel.UnitMap) UnitMap {
	unitMapViewModel, _ := tool.MapMatrix(unitMap.GetUnitMatrix(), func(colIdx int, rowIdx int, unit commonmodel.Unit) (Unit, error) {
		return NewUnit(unit), nil
	})
	return unitMapViewModel
}

func (dto UnitMap) ToValueObject() (commonmodel.UnitMap, error) {
	unitMatrix, _ := tool.MapMatrix(dto, func(colIdx int, rowIdx int, unitVm Unit) (commonmodel.Unit, error) {
		unitVo, _ := unitVm.ToValueObject()
		return unitVo, nil
	})
	return commonmodel.NewUnitMap(unitMatrix), nil
}
