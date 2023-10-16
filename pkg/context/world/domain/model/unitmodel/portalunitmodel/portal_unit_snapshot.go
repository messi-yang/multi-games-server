package portalunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type PortalUnitSnapshot struct {
	TargetPos *struct {
		X int `json:"x"`
		Z int `json:"z"`
	} `json:"targetPosition"`
}

// Interface Implementation Check
var _ domain.ValueObject[PortalUnitSnapshot] = (*PortalUnitSnapshot)(nil)

func NewPortalUnitSnapshot(targetPosition *worldcommonmodel.Position) PortalUnitSnapshot {
	if targetPosition == nil {
		return PortalUnitSnapshot{
			TargetPos: nil,
		}
	} else {
		return PortalUnitSnapshot{
			TargetPos: &struct {
				X int `json:"x"`
				Z int `json:"z"`
			}{
				X: targetPosition.GetX(),
				Z: targetPosition.GetZ(),
			},
		}
	}
}

func (snapshot PortalUnitSnapshot) IsEqual(otherSnapshot PortalUnitSnapshot) bool {
	return snapshot == otherSnapshot
}
