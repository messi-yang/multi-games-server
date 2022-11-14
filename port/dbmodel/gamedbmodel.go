package dbmodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamemodel"
	"github.com/google/uuid"
)

type UnitDbModel struct {
	ItemType string `json:"item_type"`
}

type GameDbModel struct {
	Id        uuid.UUID       `gorm:"primaryKey;unique;not null"`
	Width     int             `gorm:"not null"`
	Height    int             `gorm:"not null"`
	UnitBlock [][]UnitDbModel `gorm:"serializer:json"`
	CreatedAt time.Time       `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime;not null"`
}

func (GameDbModel) TableName() string {
	return "games"
}

func convertUnitDbModelBlockToUnitBlock(unitModelBlock [][]UnitDbModel) gamecommonmodel.UnitBlock {
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

func convertUnitBlockToUnitDbModelBlock(unitBlock gamecommonmodel.UnitBlock) [][]UnitDbModel {
	unitModelBlock := make([][]UnitDbModel, 0)
	for unitColIdx, unitCol := range unitBlock.GetUnitMatrix() {
		unitModelBlock = append(unitModelBlock, []UnitDbModel{})
		for _, unit := range unitCol {
			unitModelBlock[unitColIdx] = append(unitModelBlock[unitColIdx], UnitDbModel{
				ItemType: string(unit.GetItemType()),
			})
		}
	}
	return unitModelBlock
}

func NewGameDbModel(game gamemodel.Game) GameDbModel {
	return GameDbModel{
		Id:        game.GetId().GetId(),
		Width:     game.GetUnitBlockDimension().GetWidth(),
		Height:    game.GetUnitBlockDimension().GetHeight(),
		UnitBlock: convertUnitBlockToUnitDbModelBlock(game.GetUnitBlock()),
	}
}

func (gameDbModel GameDbModel) ToAggregate() gamemodel.Game {
	return gamemodel.NewGame(gamemodel.NewGameId(gameDbModel.Id), convertUnitDbModelBlockToUnitBlock(gameDbModel.UnitBlock))
}
