package pgmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type PortalUnitInfoModel struct {
	Id           uuid.UUID  `gorm:"not null"`
	WorldId      uuid.UUID  `gorm:"not null"`
	TargetPosX   *int       `gorm:""`
	TargetPosZ   *int       `gorm:""`
	TargetUnitId *uuid.UUID `gorm:"not null"`
}

func (PortalUnitInfoModel) TableName() string {
	return "portal_unit_infos"
}

func NewPortalUnitInfoModel(portalUnit portalunitmodel.PortalUnit) PortalUnitInfoModel {
	targetPosition := portalUnit.GetTargetPosition()
	targetUnitId := portalUnit.GetTargetUnitId()

	return PortalUnitInfoModel{
		Id:      portalUnit.GetId().Uuid(),
		WorldId: portalUnit.GetWorldId().Uuid(),
		TargetPosX: lo.TernaryF(
			targetPosition == nil,
			func() *int { return nil },
			func() *int { return commonutil.ToPointer(targetPosition.GetX()) },
		),
		TargetPosZ: lo.TernaryF(
			targetPosition == nil,
			func() *int { return nil },
			func() *int { return commonutil.ToPointer(targetPosition.GetZ()) },
		),
		TargetUnitId: lo.TernaryF(
			targetUnitId == nil,
			func() *uuid.UUID { return nil },
			func() *uuid.UUID { return commonutil.ToPointer(targetUnitId.Uuid()) },
		),
	}
}
