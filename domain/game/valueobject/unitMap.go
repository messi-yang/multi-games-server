package valueobject

import "math/rand"

type UnitMap [][]Unit

func NewUnitMap(mapSize MapSize) UnitMap {
	unitMap := make(UnitMap, mapSize.GetWidth())
	for i := 0; i < mapSize.GetWidth(); i += 1 {
		unitMap[i] = make([]Unit, mapSize.GetHeight())
		for j := 0; j < mapSize.GetHeight(); j += 1 {
			unitMap[i][j] = NewUnit(rand.Intn(2) == 0, 0)
		}
	}
	return unitMap
}
