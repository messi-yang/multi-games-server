package aggregate

import (
	"errors"
	"time"

	"github.com/DumDumGeniuss/ggol"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

var (
	ErrAreaExceedsUnitMap              = errors.New("area should contain valid from and to coordinates and it should never exceed map size")
	ErrSomeCoordinatesNotIncludedInMap = errors.New("some coordinates are not included in the unit map")
	ErrPlayerNotFound                  = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists             = errors.New("the play with the given id already exists")
)

func ggolNextUnitGenerator(
	coord *ggol.Coordinate,
	cell *valueobject.Unit,
	getAdjacentUnit ggol.AdjacentUnitGetter[valueobject.Unit],
) (nextUnit *valueobject.Unit) {
	var aliveAdjacentCellsCount int = 0
	for i := -1; i < 2; i += 1 {
		for j := -1; j < 2; j += 1 {
			if !(i == 0 && j == 0) {
				adjUnit, _ := getAdjacentUnit(coord, &ggol.Coordinate{X: i, Y: j})
				if adjUnit.GetAlive() {
					aliveAdjacentCellsCount += 1
				}
			}
		}
	}
	alive := cell.GetAlive()
	age := cell.GetAge()
	if alive {
		if aliveAdjacentCellsCount != 2 && aliveAdjacentCellsCount != 3 {
			nextCell := valueobject.NewUnit(false, 0)
			return &nextCell
		} else {
			nextCell := valueobject.NewUnit(alive, age+1)
			return &nextCell
		}
	} else {
		if aliveAdjacentCellsCount == 3 {
			nextCell := valueobject.NewUnit(true, 0)
			return &nextCell
		} else {
			return cell
		}
	}
}

type GameRoom struct {
	game        *entity.Game
	players     map[uuid.UUID]entity.Player
	zoomedAreas map[uuid.UUID]valueobject.Area
}

func NewGameRoom(game entity.Game, players map[uuid.UUID]entity.Player, zoomedAreas map[uuid.UUID]valueobject.Area) GameRoom {
	return GameRoom{
		game:        &game,
		players:     players,
		zoomedAreas: zoomedAreas,
	}
}

func (gr *GameRoom) GetId() uuid.UUID {
	return gr.game.GetId()
}

func (gr *GameRoom) GetMapSize() valueobject.MapSize {
	return gr.game.GetMapSize()
}

func (gr *GameRoom) GetUnitMap() *valueobject.UnitMap {
	return gr.game.GetUnitMap()
}

func (gr *GameRoom) GetUnitMapByArea(area valueobject.Area) (*valueobject.UnitMap, error) {
	if !gr.GetMapSize().IncludesArea(area) {
		return &valueobject.UnitMap{}, ErrAreaExceedsUnitMap
	}
	offsetX := area.GetFrom().GetX()
	offsetY := area.GetFrom().GetY()
	areaWidth := area.GetWidth()
	areaHeight := area.GetHeight()
	unitMatrix := make([][]valueobject.Unit, areaWidth)
	for x := 0; x < areaWidth; x += 1 {
		unitMatrix[x] = make([]valueobject.Unit, areaHeight)
		for y := 0; y < areaHeight; y += 1 {
			coordinate, _ := valueobject.NewCoordinate(x+offsetX, y+offsetY)
			unitMatrix[x][y] = gr.game.GetUnit(coordinate)
		}
	}
	unitMap := valueobject.NewUnitMapFromUnitMatrix(&unitMatrix)

	return unitMap, nil
}

func (gr *GameRoom) GetTickedAt() time.Time {
	return gr.game.GetTickedAt()
}

func (gr *GameRoom) GetZoomedAreas() map[uuid.UUID]valueobject.Area {
	return gr.zoomedAreas
}

func (gr *GameRoom) AddZoomedArea(playerId uuid.UUID, area valueobject.Area) error {
	_, exists := gr.players[playerId]
	if !exists {
		return ErrPlayerNotFound
	}
	gr.zoomedAreas[playerId] = area
	return nil
}

func (gr *GameRoom) RemoveZoomedArea(playerId uuid.UUID) {
	delete(gr.zoomedAreas, playerId)
}

func (gr *GameRoom) GetPlayers() map[uuid.UUID]entity.Player {
	return gr.players
}

func (gr *GameRoom) AddPlayer(newPlayer entity.Player) error {
	_, exists := gr.players[newPlayer.GetId()]
	if exists {
		return ErrPlayerAlreadyExists
	}

	gr.players[newPlayer.GetId()] = newPlayer

	return nil
}

func (gr *GameRoom) RemovePlayer(playerId uuid.UUID) {
	delete(gr.players, playerId)
}

func (gr *GameRoom) ReviveUnits(coordinates []valueobject.Coordinate) error {
	if !gr.GetMapSize().IncludesAllCoordinates(coordinates) {
		return ErrSomeCoordinatesNotIncludedInMap
	}

	for _, coordinate := range coordinates {
		unit := gr.game.GetUnit(coordinate)
		newUnit := unit.SetAlive(true)
		gr.game.SetUnit(coordinate, newUnit)
	}

	return nil
}

func (gr *GameRoom) TickUnitMap() error {
	unitMap := gr.game.GetUnitMap()

	var unitMatrix *[][]valueobject.Unit = unitMap.ToValueObjectMatrix()
	gameOfLiberty, err := ggol.NewGame(unitMatrix)
	if err != nil {
		return err
	}
	gameOfLiberty.SetNextUnitGenerator(ggolNextUnitGenerator)
	nextUnitMatrix := gameOfLiberty.GenerateNextUnits()

	newUnitMap := valueobject.NewUnitMapFromUnitMatrix(nextUnitMatrix)
	gr.game.SetUnitMap(newUnitMap)

	gr.game.SetTickedAt(time.Now())

	return nil
}
