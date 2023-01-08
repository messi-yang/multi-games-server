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
	UnitMap   [][]UnitPsqlModel `gorm:"serializer:json"`
	CreatedAt time.Time         `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time         `gorm:"autoUpdateTime;not null"`
}

func (GamePsqlModel) TableName() string {
	return "games"
}

func convertUnitMapPsqlModelToUnitMap(unitMapPsqlModel [][]UnitPsqlModel) commonmodel.UnitMap {
	unitMatrix := make([][]commonmodel.Unit, 0)
	for colIdx, unitModelCol := range unitMapPsqlModel {
		unitMatrix = append(unitMatrix, []commonmodel.Unit{})
		for _, unit := range unitModelCol {
			itemId, _ := itemmodel.NewItemId(unit.ItemId)
			unitMatrix[colIdx] = append(unitMatrix[colIdx], commonmodel.NewUnit(itemId))
		}
	}
	return commonmodel.NewUnitMap(unitMatrix)
}

func convertUnitMapToUnitMapPsqlModel(unitMap commonmodel.UnitMap) [][]UnitPsqlModel {
	unitMapPsqlModel := make([][]UnitPsqlModel, 0)
	for unitColIdx, unitCol := range unitMap.GetUnitMatrix() {
		unitMapPsqlModel = append(unitMapPsqlModel, []UnitPsqlModel{})
		for _, unit := range unitCol {
			unitMapPsqlModel[unitColIdx] = append(unitMapPsqlModel[unitColIdx], UnitPsqlModel{
				ItemId: unit.GetItemId().ToString(),
			})
		}
	}
	return unitMapPsqlModel
}

func NewGamePsqlModel(game gamemodel.Game) GamePsqlModel {
	return GamePsqlModel{
		Id:      game.GetId().ToString(),
		Width:   game.GetMapSize().GetWidth(),
		Height:  game.GetMapSize().GetHeight(),
		UnitMap: convertUnitMapToUnitMapPsqlModel(game.GetUnitMap()),
	}
}

func (gamePostgresModel GamePsqlModel) ToAggregate() gamemodel.Game {
	gameId, _ := gamemodel.NewGameId(gamePostgresModel.Id)
	return gamemodel.NewGame(gameId, convertUnitMapPsqlModelToUnitMap(gamePostgresModel.UnitMap))
}
