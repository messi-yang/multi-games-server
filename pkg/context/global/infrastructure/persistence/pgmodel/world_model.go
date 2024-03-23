package pgmodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/google/uuid"
)

type WorldModel struct {
	Id         uuid.UUID `gorm:"primaryKey"`
	UserId     uuid.UUID `gorm:"unique;not null"`
	Name       string    `gorm:"not null"`
	BoundFromX int       `gorm:"not null"`
	BoundFromZ int       `gorm:"not null"`
	BoundToX   int       `gorm:"not null"`
	BoundToZ   int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime;not null"`
}

func (WorldModel) TableName() string {
	return "worlds"
}

func NewWorldModel(world worldmodel.World) WorldModel {
	return WorldModel{
		Id:         world.GetId().Uuid(),
		UserId:     world.GetUserId().Uuid(),
		Name:       world.GetName(),
		BoundFromX: world.GetBound().GetFrom().GetX(),
		BoundFromZ: world.GetBound().GetFrom().GetZ(),
		BoundToX:   world.GetBound().GetTo().GetX(),
		BoundToZ:   world.GetBound().GetTo().GetZ(),
		UpdatedAt:  world.GetUpdatedAt(),
		CreatedAt:  world.GetCreatedAt(),
	}
}

func ParseWorldModel(worldModel WorldModel) (world worldmodel.World, err error) {
	bound, err := worldcommonmodel.NewBound(
		worldcommonmodel.NewPosition(worldModel.BoundFromX, worldModel.BoundFromZ),
		worldcommonmodel.NewPosition(worldModel.BoundToX, worldModel.BoundToZ),
	)
	if err != nil {
		return world, err
	}
	return worldmodel.LoadWorld(
		globalcommonmodel.NewWorldId(worldModel.Id),
		globalcommonmodel.NewUserId(worldModel.UserId),
		worldModel.Name,
		bound,
		worldModel.CreatedAt,
		worldModel.UpdatedAt,
	), nil
}
