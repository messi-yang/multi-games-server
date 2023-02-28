package psqlmodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/google/uuid"
)

type WorldModel struct {
	Id        uuid.UUID `gorm:"primaryKey;unique"`
	UserId    uuid.UUID `gorm:"unique;not null"`
	User      UserModel `gorm:"foreignKey:UserId;references:Id"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
}

func (WorldModel) TableName() string {
	return "worlds"
}

func NewWorldModel(game gamemodel.GameAgg) WorldModel {
	return WorldModel{
		Id:     game.GetId().Uuid(),
		UserId: game.GetUserId().Uuid(),
		Name:   game.GetName(),
	}
}

func (gamePostgresModel WorldModel) ToAggregate() gamemodel.GameAgg {
	return gamemodel.GameAgg{}
}
