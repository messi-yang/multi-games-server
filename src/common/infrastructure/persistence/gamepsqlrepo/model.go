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

func convertMapPsqlModelToMap(mapPsqlModel [][]UnitPsqlModel) commonmodel.Map {
	unitMatrix := make([][]commonmodel.Unit, 0)
	for colIdx, unitModelCol := range mapPsqlModel {
		unitMatrix = append(unitMatrix, []commonmodel.Unit{})
		for _, unit := range unitModelCol {
			itemId, _ := itemmodel.NewItemId(unit.ItemId)
			unitMatrix[colIdx] = append(unitMatrix[colIdx], commonmodel.NewUnit(itemId))
		}
	}
	return commonmodel.NewMap(unitMatrix)
}

func convertMapToMapPsqlModel(map_ commonmodel.Map) [][]UnitPsqlModel {
	mapPsqlModel := make([][]UnitPsqlModel, 0)
	for unitColIdx, unitCol := range map_.GetUnitMatrix() {
		mapPsqlModel = append(mapPsqlModel, []UnitPsqlModel{})
		for _, unit := range unitCol {
			mapPsqlModel[unitColIdx] = append(mapPsqlModel[unitColIdx], UnitPsqlModel{
				ItemId: unit.GetItemId().ToString(),
			})
		}
	}
	return mapPsqlModel
}

func NewGamePsqlModel(game gamemodel.Game) GamePsqlModel {
	return GamePsqlModel{
		Id:      game.GetId().ToString(),
		Width:   game.GetSize().GetWidth(),
		Height:  game.GetSize().GetHeight(),
		UnitMap: convertMapToMapPsqlModel(game.GetMap()),
	}
}

func (gamePostgresModel GamePsqlModel) ToAggregate() gamemodel.Game {
	gameId, _ := gamemodel.NewGameId(gamePostgresModel.Id)
	return gamemodel.NewGame(gameId, convertMapPsqlModelToMap(gamePostgresModel.UnitMap))
}
