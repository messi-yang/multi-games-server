package portalunithttphandler

import "github.com/dum-dum-genius/zossi-server/pkg/application/dto"

type getTargetPositionResponse struct {
	TargetPosition *dto.PositionDto `json:"targetPosition"`
}
