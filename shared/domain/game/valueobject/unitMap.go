package valueobject

import "github.com/google/uuid"

type UnitMap struct {
	unitMatrix *[][]Unit
}

func NewUnitMap(mapSize MapSize) *UnitMap {
	unitMap := make([][]Unit, mapSize.GetWidth())
	for i := 0; i < mapSize.GetWidth(); i += 1 {
		unitMap[i] = make([]Unit, mapSize.GetHeight())
		for j := 0; j < mapSize.GetHeight(); j += 1 {
			unitMap[i][j] = NewUnit(false, uuid.Nil)
		}
	}
	return &UnitMap{
		unitMatrix: &unitMap,
	}
}

func NewUnitMapFromUnitMatrix(unitMatrix *[][]Unit) *UnitMap {
	return &UnitMap{
		unitMatrix: unitMatrix,
	}
}

func (um UnitMap) ToValueObjectMatrix() *[][]Unit {
	return um.unitMatrix
}

func (um UnitMap) GetMapSize() MapSize {
	gameMapSize, _ := NewMapSize(len(*um.unitMatrix), len((*um.unitMatrix)[0]))
	return gameMapSize
}

func (um UnitMap) GetUnit(coord Coordinate) Unit {
	return (*um.unitMatrix)[coord.x][coord.y]
}

// TODO - Ideally we shouldn't mutate the valueobject, but the array can be super huge!
func (um UnitMap) SetUnit(coord Coordinate, unit Unit) {
	(*um.unitMatrix)[coord.x][coord.y] = unit
}
