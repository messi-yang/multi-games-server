package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
)

type MapVm [][]UnitVm

func NewMapVm(mapVo commonmodel.Map) MapVm {
	mapViewModel, _ := tool.MapMatrix(mapVo.GetUnitMatrix(), func(colIdx int, rowIdx int, unit commonmodel.Unit) (UnitVm, error) {
		return NewUnitVm(unit), nil
	})
	return mapViewModel
}

func (dto MapVm) ToValueObject() (commonmodel.Map, error) {
	unitMatrix, _ := tool.MapMatrix(dto, func(colIdx int, rowIdx int, unitVm UnitVm) (commonmodel.Unit, error) {
		unitVo, _ := unitVm.ToValueObject()
		return unitVo, nil
	})
	return commonmodel.NewMap(unitMatrix), nil
}
