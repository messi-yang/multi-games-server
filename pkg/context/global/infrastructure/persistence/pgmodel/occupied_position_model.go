package pgmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/google/uuid"
)

type OccupiedPositionModel struct {
	WorldId uuid.UUID `gorm:"not null"`
	PosX    int       `gorm:"not null"`
	PosZ    int       `gorm:"not null"`
	UnitId  uuid.UUID `gorm:"not null"`
}

func (OccupiedPositionModel) TableName() string {
	return "occupied_positions"
}

func NewOccupiedPositionsFromUnit(unit unitmodel.UnitEntity) []OccupiedPositionModel {
	return []OccupiedPositionModel{
		{
			WorldId: unit.GetWorldId().Uuid(),
			PosX:    unit.GetPosition().GetX(),
			PosZ:    unit.GetPosition().GetZ(),
			UnitId:  unit.GetId().Uuid(),
		},
	}
}
