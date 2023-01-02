package gamepsqlrepo

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type GameMapUnitPsqlModel struct {
	ItemId string `json:"item_id"`
}

type GamePsqlModel struct {
	Id        string                   `gorm:"primaryKey;unique;not null"`
	Width     int                      `gorm:"not null"`
	Height    int                      `gorm:"not null"`
	UnitBlock [][]GameMapUnitPsqlModel `gorm:"serializer:json"`
	CreatedAt time.Time                `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time                `gorm:"autoUpdateTime;not null"`
}

func (GamePsqlModel) TableName() string {
	return "games"
}

func convertGameMapUnitPsqlModelBlockToGameMap(gameMapUnitModelBlock [][]GameMapUnitPsqlModel) commonmodel.GameMap {
	gameMapUnitMatrix := make([][]commonmodel.GameMapUnit, 0)
	for colIdx, gameMapUnitModelCol := range gameMapUnitModelBlock {
		gameMapUnitMatrix = append(gameMapUnitMatrix, []commonmodel.GameMapUnit{})
		for _, gameMapUnit := range gameMapUnitModelCol {
			itemId, _ := itemmodel.NewItemId(gameMapUnit.ItemId)
			gameMapUnitMatrix[colIdx] = append(gameMapUnitMatrix[colIdx], commonmodel.NewGameMapUnit(itemId))
		}
	}
	return commonmodel.NewGameMap(gameMapUnitMatrix)
}

func convertGameMapToGameMapUnitPsqlModelBlock(gameMap commonmodel.GameMap) [][]GameMapUnitPsqlModel {
	gameMapUnitModelBlock := make([][]GameMapUnitPsqlModel, 0)
	for gameMapUnitColIdx, gameMapUnitCol := range gameMap.GetGameMapUnitMatrix() {
		gameMapUnitModelBlock = append(gameMapUnitModelBlock, []GameMapUnitPsqlModel{})
		for _, gameMapUnit := range gameMapUnitCol {
			gameMapUnitModelBlock[gameMapUnitColIdx] = append(gameMapUnitModelBlock[gameMapUnitColIdx], GameMapUnitPsqlModel{
				ItemId: gameMapUnit.GetItemId().ToString(),
			})
		}
	}
	return gameMapUnitModelBlock
}

func NewGamePsqlModel(game gamemodel.Game) GamePsqlModel {
	return GamePsqlModel{
		Id:        game.GetId().ToString(),
		Width:     game.GetGameMapMapSize().GetWidth(),
		Height:    game.GetGameMapMapSize().GetHeight(),
		UnitBlock: convertGameMapToGameMapUnitPsqlModelBlock(game.GetGameMap()),
	}
}

func (gamePostgresModel GamePsqlModel) ToAggregate() gamemodel.Game {
	gameId, _ := gamemodel.NewGameId(gamePostgresModel.Id)
	return gamemodel.NewGame(gameId, convertGameMapUnitPsqlModelBlockToGameMap(gamePostgresModel.UnitBlock))
}
