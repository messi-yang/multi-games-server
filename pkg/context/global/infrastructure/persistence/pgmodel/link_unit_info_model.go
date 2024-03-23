package pgmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/google/uuid"
)

type LinkUnitInfoModel struct {
	Id      uuid.UUID `gorm:"not null"`
	WorldId uuid.UUID `gorm:"not null"`
	Url     string    `gorm:"not null"`
}

func (LinkUnitInfoModel) TableName() string {
	return "link_unit_infos"
}

func NewLinkUnitInfoModel(linkUnit linkunitmodel.LinkUnit) LinkUnitInfoModel {
	return LinkUnitInfoModel{
		Id:      linkUnit.GetId().Uuid(),
		WorldId: linkUnit.GetWorldId().Uuid(),
		Url:     linkUnit.GetUrl().String(),
	}
}
