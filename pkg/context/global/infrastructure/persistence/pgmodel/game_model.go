package pgmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pgmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type GameModel struct {
	Id        uuid.UUID      `gorm:"primaryKey"`
	RoomId    uuid.UUID      `gorm:"not null"`
	Name      string         `gorm:"not null"`
	Started   bool           `gorm:"not null"`
	State     *pgmodel.JSONB `gorm:"type:jsonb"`
	CreatedAt time.Time      `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;not null"`
}

func (GameModel) TableName() string {
	return "games"
}

func NewGameModel(game gamemodel.Game) GameModel {
	state := game.GetState()

	return GameModel{
		Id:      game.GetId().Uuid(),
		RoomId:  game.GetRoomId().Uuid(),
		Name:    game.GetName(),
		Started: game.GetStarted(),
		State: lo.TernaryF[*pgmodel.JSONB](
			state == nil,
			func() *pgmodel.JSONB { return nil },
			func() *pgmodel.JSONB {
				return commonutil.ToPointer(pgmodel.JSONB(*state))
			}),
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
		lo.TernaryF(
			gameModel.State == nil,
			func() *map[string]interface{} { return nil },
			func() *map[string]interface{} {
				var _state map[string]interface{} = *gameModel.State
				return commonutil.ToPointer(_state)
			}),
		gameModel.CreatedAt,
		gameModel.UpdatedAt,
	)
}
