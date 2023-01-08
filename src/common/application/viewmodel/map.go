package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
)

type Map [][]Unit

func NewMap(mapVo commonmodel.Map) Map {
	mapViewModel, _ := tool.MapMatrix(mapVo.GetUnitMatrix(), func(colIdx int, rowIdx int, unit commonmodel.Unit) (Unit, error) {
		return NewUnit(unit), nil
	})
	return mapViewModel
}

func (dto Map) ToValueObject() (commonmodel.Map, error) {
	unitMatrix, _ := tool.MapMatrix(dto, func(colIdx int, rowIdx int, unitVm Unit) (commonmodel.Unit, error) {
		unitVo, _ := unitVm.ToValueObject()
		return unitVo, nil
	})
	return commonmodel.NewMap(unitMatrix), nil
}
