package gamepsqlrepo

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type UnitPsqlModel struct {
	ItemId string `json:"item_id"`
}

type GamePsqlModel struct {
	Id        string            `gorm:"primaryKey;unique;not null"`
	Width     int               `gorm:"not null"`
	Height    int               `gorm:"not null"`
	UnitBlock [][]UnitPsqlModel `gorm:"serializer:json"`
	CreatedAt time.Time         `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time         `gorm:"autoUpdateTime;not null"`
}

func (GamePsqlModel) TableName() string {
	return "games"
}

func convertUnitPsqlModelBlockToUnitBlock(unitModelBlock [][]UnitPsqlModel) commonmodel.UnitBlock {
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

func convertUnitBlockToUnitPsqlModelBlock(unitBlock commonmodel.UnitBlock) [][]UnitPsqlModel {
	unitModelBlock := make([][]UnitPsqlModel, 0)
	for unitColIdx, unitCol := range unitBlock.GetUnitMatrix() {
		unitModelBlock = append(unitModelBlock, []UnitPsqlModel{})
		for _, unit := range unitCol {
			unitModelBlock[unitColIdx] = append(unitModelBlock[unitColIdx], UnitPsqlModel{
				ItemId: unit.GetItemId().ToString(),
			})
		}
	}
	return unitModelBlock
}

func NewGamePsqlModel(game gamemodel.Game) GamePsqlModel {
	return GamePsqlModel{
		Id:        game.GetId().ToString(),
		Width:     game.GetUnitBlockDimension().GetWidth(),
		Height:    game.GetUnitBlockDimension().GetHeight(),
		UnitBlock: convertUnitBlockToUnitPsqlModelBlock(game.GetUnitBlock()),
	}
}

func (gamePostgresModel GamePsqlModel) ToAggregate() gamemodel.Game {
	gameId, _ := gamemodel.NewGameId(gamePostgresModel.Id)
	return gamemodel.NewGame(gameId, convertUnitPsqlModelBlockToUnitBlock(gamePostgresModel.UnitBlock))
}
