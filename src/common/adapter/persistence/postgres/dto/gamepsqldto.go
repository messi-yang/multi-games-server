package dto

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/itemmodel"
	"github.com/google/uuid"
)

type UnitPostgresDto struct {
	ItemId string `json:"item_id"`
}

type GamePostgresDto struct {
	Id        uuid.UUID           `gorm:"primaryKey;unique;not null"`
	Width     int                 `gorm:"not null"`
	Height    int                 `gorm:"not null"`
	UnitBlock [][]UnitPostgresDto `gorm:"serializer:json"`
	CreatedAt time.Time           `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time           `gorm:"autoUpdateTime;not null"`
}

func (GamePostgresDto) TableName() string {
	return "games"
}

func convertUnitPostgresDtoBlockToUnitBlock(unitModelBlock [][]UnitPostgresDto) commonmodel.UnitBlock {
	unitMatrix := make([][]commonmodel.Unit, 0)
	for colIdx, unitModelCol := range unitModelBlock {
		unitMatrix = append(unitMatrix, []commonmodel.Unit{})
		for _, unit := range unitModelCol {
			itemId, _ := itemmodel.NewItemId(unit.ItemId)
			unitMatrix[colIdx] = append(unitMatrix[colIdx], commonmodel.NewUnit(itemId))
		}
	}
	return commonmodel.NewUnitBlock(unitMatrix)
}

func convertUnitBlockToUnitPostgresDtoBlock(unitBlock commonmodel.UnitBlock) [][]UnitPostgresDto {
	unitModelBlock := make([][]UnitPostgresDto, 0)
	for unitColIdx, unitCol := range unitBlock.GetUnitMatrix() {
		unitModelBlock = append(unitModelBlock, []UnitPostgresDto{})
		for _, unit := range unitCol {
			unitModelBlock[unitColIdx] = append(unitModelBlock[unitColIdx], UnitPostgresDto{
				ItemId: unit.GetItemId().ToString(),
			})
		}
	}
	return unitModelBlock
}

func NewGamePostgresDto(game gamemodel.Game) GamePostgresDto {
	return GamePostgresDto{
		Id:        game.GetId().GetId(),
		Width:     game.GetUnitBlockDimension().GetWidth(),
		Height:    game.GetUnitBlockDimension().GetHeight(),
		UnitBlock: convertUnitBlockToUnitPostgresDtoBlock(game.GetUnitBlock()),
	}
}

func (gamePostgresDto GamePostgresDto) ToAggregate() gamemodel.Game {
	return gamemodel.NewGame(gamemodel.NewGameId(gamePostgresDto.Id), convertUnitPostgresDtoBlockToUnitBlock(gamePostgresDto.UnitBlock))
}
