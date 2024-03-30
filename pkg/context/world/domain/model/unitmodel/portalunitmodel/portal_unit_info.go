package portalunitmodel

import "github.com/google/uuid"

type PortalUnitInfo struct {
	TargetUnitId *uuid.UUID `json:"target_unit_id"`
}
