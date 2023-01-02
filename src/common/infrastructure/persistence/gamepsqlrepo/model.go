package gamepsqlrepo

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type MapUnitPsqlModel struct {
	ItemId string `json:"item_id"`
}

type GamePsqlModel struct {
	Id        string               `gorm:"primaryKey;unique;not null"`
	Width     int                  `gorm:"not null"`
	Height    int                  `gorm:"not null"`
	UnitBlock [][]MapUnitPsqlModel `gorm:"serializer:json"`
	CreatedAt time.Time            `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time            `gorm:"autoUpdateTime;not null"`
}

func (GamePsqlModel) TableName() string {
	return "games"
}

func convertMapUnitPsqlModelBlockToGameMap(mapUnitModelBlock [][]MapUnitPsqlModel) commonmodel.GameMap {
	mapUnitMatrix := make([][]commonmodel.MapUnit, 0)
	for colIdx, mapUnitModelCol := range mapUnitModelBlock {
		mapUnitMatrix = append(mapUnitMatrix, []commonmodel.MapUnit{})
		for _, mapUnit := range mapUnitModelCol {
			itemId, _ := itemmodel.NewItemId(mapUnit.ItemId)
			mapUnitMatrix[colIdx] = append(mapUnitMatrix[colIdx], commonmodel.NewMapUnit(itemId))
		}
	}
	return commonmodel.NewGameMap(mapUnitMatrix)
}

func convertGameMapToMapUnitPsqlModelBlock(gameMap commonmodel.GameMap) [][]MapUnitPsqlModel {
	mapUnitModelBlock := make([][]MapUnitPsqlModel, 0)
	for mapUnitColIdx, mapUnitCol := range gameMap.GetMapUnitMatrix() {
		mapUnitModelBlock = append(mapUnitModelBlock, []MapUnitPsqlModel{})
		for _, mapUnit := range mapUnitCol {
			mapUnitModelBlock[mapUnitColIdx] = append(mapUnitModelBlock[mapUnitColIdx], MapUnitPsqlModel{
				ItemId: mapUnit.GetItemId().ToString(),
			})
		}
	}
	return mapUnitModelBlock
}

func NewGamePsqlModel(game gamemodel.Game) GamePsqlModel {
	return GamePsqlModel{
		Id:        game.GetId().ToString(),
		Width:     game.GetGameMapMapSize().GetWidth(),
		Height:    game.GetGameMapMapSize().GetHeight(),
		UnitBlock: convertGameMapToMapUnitPsqlModelBlock(game.GetGameMap()),
	}
}

func (gamePostgresModel GamePsqlModel) ToAggregate() gamemodel.Game {
	gameId, _ := gamemodel.NewGameId(gamePostgresModel.Id)
	return gamemodel.NewGame(gameId, convertMapUnitPsqlModelBlockToGameMap(gamePostgresModel.UnitBlock))
}
