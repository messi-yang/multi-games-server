package pgmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/embedunitmodel"
	"github.com/google/uuid"
)

type EmbedUnitInfoModel struct {
	Id        uuid.UUID `gorm:"not null"`
	WorldId   uuid.UUID `gorm:"not null"`
	EmbedCode string    `gorm:"not null"`
}

func (EmbedUnitInfoModel) TableName() string {
	return "embed_unit_infos"
}

func NewEmbedUnitInfoModel(embedUnit embedunitmodel.EmbedUnit) EmbedUnitInfoModel {
	return EmbedUnitInfoModel{
		Id:        embedUnit.GetId().Uuid(),
		WorldId:   embedUnit.GetWorldId().Uuid(),
		EmbedCode: embedUnit.GetEmbedCode().String(),
	}
}
