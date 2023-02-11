package psqlrepo

// import (
// 	"time"

// 	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
// 	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
// 	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
// 	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/library/tool"
// )

// type UnitPsqlModel struct {
// 	ItemId string `json:"item_id"`
// }

// type GamePsqlModel struct {
// 	Id        string            `gorm:"primaryKey;unique;not null"`
// 	Width     int               `gorm:"not null"`
// 	Height    int               `gorm:"not null"`
// 	UnitMap   [][]UnitPsqlModel `gorm:"serializer:json"`
// 	CreatedAt time.Time         `gorm:"autoCreateTime;not null"`
// 	UpdatedAt time.Time         `gorm:"autoUpdateTime;not null"`
// }

// func (GamePsqlModel) TableName() string {
// 	return "games"
// }

// func NewGamePsqlModel(game gamemodel.GameAgg) GamePsqlModel {
// 	unitPsqlModelMatrix, _ := tool.MapMatrix(game.GetMap().GetUnitMatrix(), func(_ int, _ int, unit commonmodel.UnitVo) (UnitPsqlModel, error) {
// 		return UnitPsqlModel{
// 			ItemId: unit.GetItemId().ToString(),
// 		}, nil
// 	})

// 	return GamePsqlModel{
// 		Id:      game.GetId().ToString(),
// 		Width:   game.GetMapSize().GetWidth(),
// 		Height:  game.GetMapSize().GetHeight(),
// 		UnitMap: unitPsqlModelMatrix,
// 	}
// }

// func (gamePostgresModel GamePsqlModel) ToAggregate() gamemodel.GameAgg {
// 	gameId, _ := gamemodel.NewGameIdVo(gamePostgresModel.Id)
// 	unitMatrix, _ := tool.MapMatrix(gamePostgresModel.UnitMap, func(_ int, _ int, unitPsqlModel UnitPsqlModel) (commonmodel.UnitVo, error) {
// 		itemId, _ := itemmodel.NewItemIdVo(unitPsqlModel.ItemId)
// 		return commonmodel.NewUnitVo(itemId), nil
// 	})
// 	return gamemodel.NewGameAgg(gameId, gamemodel.NewMapVo(unitMatrix))
// }
