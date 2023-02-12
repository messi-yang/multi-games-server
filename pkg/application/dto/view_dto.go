package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/samber/lo"
)

type ViewDto struct {
	Bound BoundDto  `json:"bound"`
	Units []UnitDto `json:"units"`
}

func NewViewDto(view unitmodel.ViewVo) ViewDto {
	unitDtos := lo.Map(view.GetUnits(), func(unit unitmodel.UnitAgg, _ int) UnitDto {
		return NewUnitDto(unit)
	})
	return ViewDto{
		Bound: NewBoundDto(view.GetBound()),
		Units: unitDtos,
	}
}
