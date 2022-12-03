package dto

import (
	"time"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/gamemodel"
	"github.com/google/uuid"
)

type UnitPostgresJsonDto struct {
	ItemType string `json:"item_type"`
}

type GamePostgresJsonDto struct {
	Id        uuid.UUID               `gorm:"primaryKey;unique;not null"`
	Width     int                     `gorm:"not null"`
	Height    int                     `gorm:"not null"`
	UnitBlock [][]UnitPostgresJsonDto `gorm:"serializer:json"`
	CreatedAt time.Time               `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time               `gorm:"autoUpdateTime;not null"`
}

func (GamePostgresJsonDto) TableName() string {
	return "games"
}

func convertUnitPostgresJsonDtoBlockToUnitBlock(unitModelBlock [][]UnitPostgresJsonDto) gamecommonmodel.UnitBlock {
	unitMatrix := make([][]gamecommonmodel.Unit, 0)
	for colIdx, unitModelCol := range unitModelBlock {
		unitMatrix = append(unitMatrix, []gamecommonmodel.Unit{})
		for _, unit := range unitModelCol {
			unitMatrix[colIdx] = append(unitMatrix[colIdx], gamecommonmodel.NewUnit(
				unit.ItemType != string(gamecommonmodel.ItemTypeEmpty),
				gamecommonmodel.ItemTypeEmpty,
			))
		}
	}
	return gamecommonmodel.NewUnitBlock(unitMatrix)
}

func convertUnitBlockToUnitPostgresJsonDtoBlock(unitBlock gamecommonmodel.UnitBlock) [][]UnitPostgresJsonDto {
	unitModelBlock := make([][]UnitPostgresJsonDto, 0)
	for unitColIdx, unitCol := range unitBlock.GetUnitMatrix() {
		unitModelBlock = append(unitModelBlock, []UnitPostgresJsonDto{})
		for _, unit := range unitCol {
			unitModelBlock[unitColIdx] = append(unitModelBlock[unitColIdx], UnitPostgresJsonDto{
				ItemType: string(unit.GetItemType()),
			})
		}
	}
	return unitModelBlock
}

func NewGamePostgresJsonDto(game gamemodel.Game) GamePostgresJsonDto {
	return GamePostgresJsonDto{
		Id:        game.GetId().GetId(),
		Width:     game.GetUnitBlockDimension().GetWidth(),
		Height:    game.GetUnitBlockDimension().GetHeight(),
		UnitBlock: convertUnitBlockToUnitPostgresJsonDtoBlock(game.GetUnitBlock()),
	}
}

func (gamePostgresJsonDto GamePostgresJsonDto) ToAggregate() gamemodel.Game {
	return gamemodel.NewGame(gamemodel.NewGameId(gamePostgresJsonDto.Id), convertUnitPostgresJsonDtoBlockToUnitBlock(gamePostgresJsonDto.UnitBlock))
}
