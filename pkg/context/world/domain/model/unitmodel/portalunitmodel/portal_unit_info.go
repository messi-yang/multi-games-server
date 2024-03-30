package portalunitmodel

import "github.com/google/uuid"

type PortalUnitInfo struct {
	TargetPos *struct {
		X int `json:"x"`
		Z int `json:"z"`
	} `json:"targetPosition"`
	TargetUnitId *uuid.UUID `json:"target_unit_id"`
}
