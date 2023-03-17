package psqlmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/google/uuid"
)

type UnitModel struct {
	WorldId   uuid.UUID `gorm:"not null"`
	PosX      int       `gorm:"not null"`
	PosZ      int       `gorm:"not null"`
	ItemId    uuid.UUID `gorm:"not null"`
	Direction int8      `gorm:"not null"`
}

func (UnitModel) TableName() string {
	return "units"
}

func NewUnitModel(unit unitmodel.UnitAgg) UnitModel {
	return UnitModel{
		WorldId:   unit.GetWorldId().Uuid(),
		PosX:      unit.GetPosition().GetX(),
		PosZ:      unit.GetPosition().GetZ(),
		ItemId:    unit.GetItemId().Uuid(),
		Direction: unit.GetDirection().Int8(),
	}
}

func (model UnitModel) ToAggregate() unitmodel.UnitAgg {
	direction, _ := commonmodel.NewDirectionVo(model.Direction)
	return unitmodel.NewUnitAgg(
		worldmodel.NewWorldIdVo(model.WorldId),
		commonmodel.NewPositionVo(model.PosX, model.PosZ),
		itemmodel.NewItemIdVo(model.ItemId),
		direction,
	)
}
