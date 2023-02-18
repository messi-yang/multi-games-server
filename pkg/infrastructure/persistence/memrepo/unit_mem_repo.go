package memrepo

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/tool"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
)

var (
	ErrUnitNotFound = errors.New("the unit in the location was not found")
)

type unitMemRepo struct {
	units map[gamemodel.GameIdVo]map[commonmodel.LocationVo]unitmodel.UnitAgg
}

var unitMemRepoSingleton *unitMemRepo

func NewUnitMemRepo() unitmodel.Repo {
	if unitMemRepoSingleton == nil {
		unitMemRepoSingleton = &unitMemRepo{
			units: make(map[gamemodel.GameIdVo]map[commonmodel.LocationVo]unitmodel.UnitAgg),
		}
		return unitMemRepoSingleton
	}
	return unitMemRepoSingleton
}

func (m *unitMemRepo) GetUnitAt(gameId gamemodel.GameIdVo, location commonmodel.LocationVo) (unitmodel.UnitAgg, bool) {
	unit, exists := m.units[gameId][location]
	if !exists {
		return unitmodel.UnitAgg{}, false
	}
	return unit, true
}

func (m *unitMemRepo) GetUnitsInBound(gameId gamemodel.GameIdVo, bound commonmodel.BoundVo) []unitmodel.UnitAgg {
	units := make([]unitmodel.UnitAgg, 0)
	tool.RangeMatrix(bound.GetWidth(), bound.GetHeight(), func(x int, z int) {
		location := commonmodel.NewLocationVo(x+bound.GetFrom().GetX(), z+bound.GetFrom().GetZ())
		unit, exists := m.units[gameId][location]
		if exists {
			units = append(units, unit)
		}
	})
	return units
}

func (m *unitMemRepo) Update(unit unitmodel.UnitAgg) {
	_, exists := m.units[unit.GetGameId()]
	if !exists {
		m.units[unit.GetGameId()] = make(map[commonmodel.LocationVo]unitmodel.UnitAgg)
	}
	m.units[unit.GetGameId()][unit.GetLocation()] = unit
}

func (m *unitMemRepo) Delete(gameId gamemodel.GameIdVo, location commonmodel.LocationVo) {
	_, exists := m.units[gameId]
	if !exists {
		return
	}
	delete(m.units[gameId], location)
}
