package postgresgamerepository

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/aggregate"
	gameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
	"github.com/google/uuid"
)

type UnitInDb struct {
	ItemType string `json:"item_type"`
}

type Game struct {
	Id        uuid.UUID    `gorm:"primaryKey;unique;not null"`
	Width     int          `gorm:"not null"`
	Height    int          `gorm:"not null"`
	UnitBlock [][]UnitInDb `gorm:"serializer:json"`
	CreatedAt time.Time    `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime;not null"`
}

func convertUnitBlockInDbToUnitBlock(unitBlockInDb [][]UnitInDb) valueobject.UnitBlock {
	unitMatrix := make([][]valueobject.Unit, 0)
	for colIdx, colOfUnitBlockInDb := range unitBlockInDb {
		unitMatrix = append(unitMatrix, []valueobject.Unit{})
		for _, unit := range colOfUnitBlockInDb {
			unitMatrix[colIdx] = append(unitMatrix[colIdx], valueobject.NewUnit(
				unit.ItemType != string(valueobject.ItemTypeEmpty),
				valueobject.ItemTypeEmpty,
			))
		}
	}
	return valueobject.NewUnitBlock(unitMatrix)
}

func convertUnitBlockToUnitBlockInDb(unitBlock valueobject.UnitBlock) [][]UnitInDb {
	unitBlockInDb := make([][]UnitInDb, 0)
	for unitColIdx, unitCol := range unitBlock.GetUnitMatrix() {
		unitBlockInDb = append(unitBlockInDb, []UnitInDb{})
		for _, unit := range unitCol {
			unitBlockInDb[unitColIdx] = append(unitBlockInDb[unitColIdx], UnitInDb{
				ItemType: string(unit.GetItemType()),
			})
		}
	}
	return unitBlockInDb
}

func NewGameModel(gameAggregate aggregate.Game) Game {
	return Game{
		Id:        gameAggregate.GetId().GetId(),
		Width:     gameAggregate.GetUnitBlockDimension().GetWidth(),
		Height:    gameAggregate.GetUnitBlockDimension().GetHeight(),
		UnitBlock: convertUnitBlockToUnitBlockInDb(gameAggregate.GetUnitBlock()),
	}
}

func (game Game) ToAggregate() aggregate.Game {
	return aggregate.NewGame(gameValueObject.NewGameId(game.Id), convertUnitBlockInDbToUnitBlock(game.UnitBlock))
}
