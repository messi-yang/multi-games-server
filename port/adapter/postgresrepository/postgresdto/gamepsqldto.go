package postgresdto

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamemodel"
	"github.com/google/uuid"
)

type UnitPostgresUiDto struct {
	ItemType string `json:"item_type"`
}

type GamePostgresUiDto struct {
	Id        uuid.UUID             `gorm:"primaryKey;unique;not null"`
	Width     int                   `gorm:"not null"`
	Height    int                   `gorm:"not null"`
	UnitBlock [][]UnitPostgresUiDto `gorm:"serializer:json"`
	CreatedAt time.Time             `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time             `gorm:"autoUpdateTime;not null"`
}

func (GamePostgresUiDto) TableName() string {
	return "games"
}

func convertUnitPostgresUiDtoBlockToUnitBlock(unitModelBlock [][]UnitPostgresUiDto) gamecommonmodel.UnitBlock {
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

func convertUnitBlockToUnitPostgresUiDtoBlock(unitBlock gamecommonmodel.UnitBlock) [][]UnitPostgresUiDto {
	unitModelBlock := make([][]UnitPostgresUiDto, 0)
	for unitColIdx, unitCol := range unitBlock.GetUnitMatrix() {
		unitModelBlock = append(unitModelBlock, []UnitPostgresUiDto{})
		for _, unit := range unitCol {
			unitModelBlock[unitColIdx] = append(unitModelBlock[unitColIdx], UnitPostgresUiDto{
				ItemType: string(unit.GetItemType()),
			})
		}
	}
	return unitModelBlock
}

func NewGamePostgresUiDto(game gamemodel.Game) GamePostgresUiDto {
	return GamePostgresUiDto{
		Id:        game.GetId().GetId(),
		Width:     game.GetUnitBlockDimension().GetWidth(),
		Height:    game.GetUnitBlockDimension().GetHeight(),
		UnitBlock: convertUnitBlockToUnitPostgresUiDtoBlock(game.GetUnitBlock()),
	}
}

func (gamePostgresUiDto GamePostgresUiDto) ToAggregate() gamemodel.Game {
	return gamemodel.NewGame(gamemodel.NewGameId(gamePostgresUiDto.Id), convertUnitPostgresUiDtoBlockToUnitBlock(gamePostgresUiDto.UnitBlock))
}
