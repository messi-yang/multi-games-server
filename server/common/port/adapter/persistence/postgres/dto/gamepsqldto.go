package dto

import (
	"time"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/gamemodel"
	"github.com/google/uuid"
)

type UnitPostgresPresenterDto struct {
	ItemType string `json:"item_type"`
}

type GamePostgresPresenterDto struct {
	Id        uuid.UUID                    `gorm:"primaryKey;unique;not null"`
	Width     int                          `gorm:"not null"`
	Height    int                          `gorm:"not null"`
	UnitBlock [][]UnitPostgresPresenterDto `gorm:"serializer:json"`
	CreatedAt time.Time                    `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time                    `gorm:"autoUpdateTime;not null"`
}

func (GamePostgresPresenterDto) TableName() string {
	return "games"
}

func convertUnitPostgresPresenterDtoBlockToUnitBlock(unitModelBlock [][]UnitPostgresPresenterDto) gamecommonmodel.UnitBlock {
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

func convertUnitBlockToUnitPostgresPresenterDtoBlock(unitBlock gamecommonmodel.UnitBlock) [][]UnitPostgresPresenterDto {
	unitModelBlock := make([][]UnitPostgresPresenterDto, 0)
	for unitColIdx, unitCol := range unitBlock.GetUnitMatrix() {
		unitModelBlock = append(unitModelBlock, []UnitPostgresPresenterDto{})
		for _, unit := range unitCol {
			unitModelBlock[unitColIdx] = append(unitModelBlock[unitColIdx], UnitPostgresPresenterDto{
				ItemType: string(unit.GetItemType()),
			})
		}
	}
	return unitModelBlock
}

func NewGamePostgresPresenterDto(game gamemodel.Game) GamePostgresPresenterDto {
	return GamePostgresPresenterDto{
		Id:        game.GetId().GetId(),
		Width:     game.GetUnitBlockDimension().GetWidth(),
		Height:    game.GetUnitBlockDimension().GetHeight(),
		UnitBlock: convertUnitBlockToUnitPostgresPresenterDtoBlock(game.GetUnitBlock()),
	}
}

func (gamePostgresPresenterDto GamePostgresPresenterDto) ToAggregate() gamemodel.Game {
	return gamemodel.NewGame(gamemodel.NewGameId(gamePostgresPresenterDto.Id), convertUnitPostgresPresenterDtoBlockToUnitBlock(gamePostgresPresenterDto.UnitBlock))
}
