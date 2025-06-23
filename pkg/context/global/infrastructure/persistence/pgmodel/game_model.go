package pgmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pgmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type GameModel struct {
	Id        uuid.UUID     `gorm:"primaryKey"`
	RoomId    uuid.UUID     `gorm:"not null"`
	Name      string        `gorm:"not null"`
	Started   bool          `gorm:"not null"`
	State     pgmodel.JSONB `gorm:"type:jsonb;not null"`
	CreatedAt time.Time     `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime;not null"`
}

func (GameModel) TableName() string {
	return "games"
}

func NewGameModel(game gamemodel.Game) GameModel {
	return GameModel{
		Id:        game.GetId().Uuid(),
		RoomId:    game.GetRoomId().Uuid(),
		Name:      game.GetName(),
		Started:   game.GetStarted(),
		State:     pgmodel.JSONB(game.GetState()),
		UpdatedAt: game.GetUpdatedAt(),
		CreatedAt: game.GetCreatedAt(),
	}
}

func ParseGameModel(gameModel GameModel) (game gamemodel.Game) {
	return gamemodel.LoadGame(
		gamemodel.NewGameId(gameModel.Id),
		globalcommonmodel.NewRoomId(gameModel.RoomId),
		gameModel.Name,
		gameModel.Started,
		gameModel.State,
		gameModel.CreatedAt,
		gameModel.UpdatedAt,
	)
}
