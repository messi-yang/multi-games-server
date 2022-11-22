package livegamemodel

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/gamecommonmodel"
)

var (
	ErrAreaExceedsUnitBlock            = errors.New("area should contain valid from and to coordinates and it should never exceed dimension")
	ErrSomeCoordinatesNotIncludedInMap = errors.New("some coordinates are not included in the unit map")
	ErrPlayerNotFound                  = errors.New("the play with the given id does not exist")
	ErrPlayerAlreadyExists             = errors.New("the play with the given id already exists")
)

type LiveGame struct {
	id          LiveGameId
	unitBlock   gamecommonmodel.UnitBlock
	playerIds   map[gamecommonmodel.PlayerId]bool
	zoomedAreas map[gamecommonmodel.PlayerId]gamecommonmodel.Area
}

func NewLiveGame(id LiveGameId, unitBlock gamecommonmodel.UnitBlock) LiveGame {
	return LiveGame{
		id:          id,
		unitBlock:   unitBlock,
		playerIds:   make(map[gamecommonmodel.PlayerId]bool),
		zoomedAreas: make(map[gamecommonmodel.PlayerId]gamecommonmodel.Area),
	}
}

func (liveGame *LiveGame) GetId() LiveGameId {
	return liveGame.id
}

func (liveGame *LiveGame) GetDimension() gamecommonmodel.Dimension {
	return liveGame.unitBlock.GetDimension()
}

func (liveGame *LiveGame) GetUnitBlock() gamecommonmodel.UnitBlock {
	return liveGame.unitBlock
}

func (liveGame *LiveGame) GetUnitBlockByArea(area gamecommonmodel.Area) (gamecommonmodel.UnitBlock, error) {
	if !liveGame.GetDimension().IncludesArea(area) {
		return gamecommonmodel.UnitBlock{}, ErrAreaExceedsUnitBlock
	}
	offsetX := area.GetFrom().GetX()
	offsetY := area.GetFrom().GetY()
	areaWidth := area.GetWidth()
	areaHeight := area.GetHeight()
	unitMatrix := make([][]gamecommonmodel.Unit, areaWidth)
	for x := 0; x < areaWidth; x += 1 {
		unitMatrix[x] = make([]gamecommonmodel.Unit, areaHeight)
		for y := 0; y < areaHeight; y += 1 {
			coordinate, _ := gamecommonmodel.NewCoordinate(x+offsetX, y+offsetY)
			unitMatrix[x][y] = liveGame.unitBlock.GetUnit(coordinate)
		}
	}
	unitBlock := gamecommonmodel.NewUnitBlock(unitMatrix)

	return unitBlock, nil
}

func (liveGame *LiveGame) GetZoomedAreas() map[gamecommonmodel.PlayerId]gamecommonmodel.Area {
	return liveGame.zoomedAreas
}

func (liveGame *LiveGame) AddZoomedArea(playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) error {
	_, exists := liveGame.playerIds[playerId]
	if !exists {
		return ErrPlayerNotFound
	}
	liveGame.zoomedAreas[playerId] = area
	return nil
}

func (liveGame *LiveGame) RemoveZoomedArea(playerId gamecommonmodel.PlayerId) {
	delete(liveGame.zoomedAreas, playerId)
}

func (liveGame *LiveGame) AddPlayer(playerId gamecommonmodel.PlayerId) error {
	_, exists := liveGame.playerIds[playerId]
	if exists {
		return ErrPlayerAlreadyExists
	}

	liveGame.playerIds[playerId] = true

	return nil
}

func (liveGame *LiveGame) RemovePlayer(playerId gamecommonmodel.PlayerId) {
	delete(liveGame.playerIds, playerId)
}

func (liveGame *LiveGame) ReviveUnits(coordinates []gamecommonmodel.Coordinate) error {
	if !liveGame.GetDimension().IncludesAllCoordinates(coordinates) {
		return ErrSomeCoordinatesNotIncludedInMap
	}

	for _, coordinate := range coordinates {
		unit := liveGame.unitBlock.GetUnit(coordinate)
		newUnit := unit.SetAlive(true)
		liveGame.unitBlock.SetUnit(coordinate, newUnit)
	}

	return nil
}
