package pgmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
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

func NewOccupiedPositionModels(unit unitmodel.UnitEntity) []OccupiedPositionModel {
	return lo.Map(unit.GetOccupiedPositions(), func(position worldcommonmodel.Position, _ int) OccupiedPositionModel {
		return OccupiedPositionModel{
			WorldId: unit.GetWorldId().Uuid(),
			PosX:    position.GetX(),
			PosZ:    position.GetZ(),
			UnitId:  unit.GetId().Uuid(),
		}
	})
}
