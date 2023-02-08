package memrepo

import (
	"errors"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
)

var (
	ErrUnitNotFound = errors.New("the unit in the location was not found")
)

type unitMemRepo struct {
	unitRecords   map[gamemodel.GameIdVo]map[commonmodel.LocationVo]unitmodel.UnitAgg
	recordLockers map[gamemodel.GameIdVo]*sync.RWMutex
}

var unitMemRepoSingleton *unitMemRepo

func NewUnitMemRepo() unitmodel.Repo {
	if unitMemRepoSingleton == nil {
		unitMemRepoSingleton = &unitMemRepo{
			unitRecords:   make(map[gamemodel.GameIdVo]map[commonmodel.LocationVo]unitmodel.UnitAgg),
			recordLockers: make(map[gamemodel.GameIdVo]*sync.RWMutex),
		}
		return unitMemRepoSingleton
	}
	return unitMemRepoSingleton
}

func (m *unitMemRepo) GetUnit(gameId gamemodel.GameIdVo, location commonmodel.LocationVo) (unitmodel.UnitAgg, error) {
	unit, exists := m.unitRecords[gameId][location]
	if !exists {
		return unitmodel.UnitAgg{}, ErrUnitNotFound
	}
	return unit, nil
}

func (m *unitMemRepo) GetUnits(gameId gamemodel.GameIdVo, bound commonmodel.BoundVo) []unitmodel.UnitAgg {
	units := make([]unitmodel.UnitAgg, 0)
	tool.ForMatrix(bound.GetWidth(), bound.GetHeight(), func(x int, y int) {
		location := commonmodel.NewLocationVo(x+bound.GetFrom().GetX(), y+bound.GetFrom().GetY())
		unit, exists := m.unitRecords[gameId][location]
		if exists {
			units = append(units, unit)
		}
	})
	return units
}

func (m *unitMemRepo) UpdateUnit(unit unitmodel.UnitAgg) {
	_, exists := m.unitRecords[unit.GetGameId()]
	if !exists {
		m.unitRecords[unit.GetGameId()] = make(map[commonmodel.LocationVo]unitmodel.UnitAgg)
	}
	m.unitRecords[unit.GetGameId()][unit.GetLocation()] = unit
}

func (m *unitMemRepo) DeleteUnit(gameId gamemodel.GameIdVo, location commonmodel.LocationVo) {
	_, exists := m.unitRecords[gameId]
	if !exists {
		return
	}
	delete(m.unitRecords[gameId], location)
}
