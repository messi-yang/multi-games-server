package pgmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/colorunitmodel"
	"github.com/google/uuid"
)

type ColorUnitInfoModel struct {
	Id      uuid.UUID `gorm:"not null"`
	WorldId uuid.UUID `gorm:"not null"`
	Color   string    `gorm:"not null"`
}

func (ColorUnitInfoModel) TableName() string {
	return "color_unit_infos"
}

func NewColorUnitInfoModel(colorUnit colorunitmodel.ColorUnit) ColorUnitInfoModel {
	return ColorUnitInfoModel{
		Id:      colorUnit.GetId().Uuid(),
		WorldId: colorUnit.GetWorldId().Uuid(),
		Color:   colorUnit.GetColor().HexString(),
	}
}
