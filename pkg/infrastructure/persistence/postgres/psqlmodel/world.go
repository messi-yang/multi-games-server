package psqlmodel

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
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

func NewWorldModel(world worldmodel.WorldAgg) WorldModel {
	return WorldModel{
		Id:     world.GetId().Uuid(),
		UserId: world.GetUserId().Uuid(),
		Name:   world.GetName(),
	}
}

func (worldPostgresModel WorldModel) ToAggregate() worldmodel.WorldAgg {
	return worldmodel.WorldAgg{}
}
