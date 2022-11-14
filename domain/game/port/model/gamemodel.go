package model

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamemodel"
	"github.com/google/uuid"
)

type UnitModel struct {
	ItemType string `json:"item_type"`
}

type GameModel struct {
	Id        uuid.UUID     `gorm:"primaryKey;unique;not null"`
	Width     int           `gorm:"not null"`
	Height    int           `gorm:"not null"`
	UnitBlock [][]UnitModel `gorm:"serializer:json"`
	CreatedAt time.Time     `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime;not null"`
}

func (GameModel) TableName() string {
	return "games"
}

func convertUnitModelBlockToUnitBlock(unitModelBlock [][]UnitModel) valueobject.UnitBlock {
	unitMatrix := make([][]valueobject.Unit, 0)
	for colIdx, unitModelCol := range unitModelBlock {
		unitMatrix = append(unitMatrix, []valueobject.Unit{})
		for _, unit := range unitModelCol {
			unitMatrix[colIdx] = append(unitMatrix[colIdx], valueobject.NewUnit(
				unit.ItemType != string(valueobject.ItemTypeEmpty),
				valueobject.ItemTypeEmpty,
			))
		}
	}
	return valueobject.NewUnitBlock(unitMatrix)
}

func convertUnitBlockToUnitModelBlock(unitBlock valueobject.UnitBlock) [][]UnitModel {
	unitModelBlock := make([][]UnitModel, 0)
	for unitColIdx, unitCol := range unitBlock.GetUnitMatrix() {
		unitModelBlock = append(unitModelBlock, []UnitModel{})
		for _, unit := range unitCol {
			unitModelBlock[unitColIdx] = append(unitModelBlock[unitColIdx], UnitModel{
				ItemType: string(unit.GetItemType()),
			})
		}
	}
	return unitModelBlock
}

func NewGameModel(gameAggregate gamemodel.Game) GameModel {
	return GameModel{
		Id:        gameAggregate.GetId().GetId(),
		Width:     gameAggregate.GetUnitBlockDimension().GetWidth(),
		Height:    gameAggregate.GetUnitBlockDimension().GetHeight(),
		UnitBlock: convertUnitBlockToUnitModelBlock(gameAggregate.GetUnitBlock()),
	}
}

func (gameModel GameModel) ToAggregate() gamemodel.Game {
	return gamemodel.NewGame(gamemodel.NewGameId(gameModel.Id), convertUnitModelBlockToUnitBlock(gameModel.UnitBlock))
}
