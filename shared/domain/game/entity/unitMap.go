package entity

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

type UnitMap struct {
	unitMatrix *[][]valueobject.Unit
}

func NewUnitMap(mapSize valueobject.MapSize) *UnitMap {
	unitMap := make([][]valueobject.Unit, mapSize.GetWidth())
	for i := 0; i < mapSize.GetWidth(); i += 1 {
		unitMap[i] = make([]valueobject.Unit, mapSize.GetHeight())
		for j := 0; j < mapSize.GetHeight(); j += 1 {
			unitMap[i][j] = valueobject.NewUnit(false, uuid.Nil)
		}
	}
	return &UnitMap{
		unitMatrix: &unitMap,
	}
}

func NewUnitMapFromUnitMatrix(unitMatrix *[][]valueobject.Unit) *UnitMap {
	return &UnitMap{
		unitMatrix: unitMatrix,
	}
}

func (um UnitMap) ToValueObjectMatrix() *[][]valueobject.Unit {
	return um.unitMatrix
}

func (um UnitMap) GetMapSize() valueobject.MapSize {
	gameMapSize, _ := valueobject.NewMapSize(len(*um.unitMatrix), len((*um.unitMatrix)[0]))
	return gameMapSize
}

func (um UnitMap) GetUnit(coord valueobject.Coordinate) valueobject.Unit {
	return (*um.unitMatrix)[coord.GetX()][coord.GetY()]
}

// TODO - Ideally we shouldn't mutate the valueobject, but the array can be super huge!
func (um UnitMap) SetUnit(coord valueobject.Coordinate, unit valueobject.Unit) {
	(*um.unitMatrix)[coord.GetX()][coord.GetY()] = unit
}
