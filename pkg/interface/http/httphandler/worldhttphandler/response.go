package worldhttphandler

import (
	iam_dto "github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
)

type worldViewModel struct {
	dto.WorldDto
	User iam_dto.UserDto `json:"user"`
}

type getWorldResponse worldViewModel

type queryWorldsResponse []worldViewModel

type getMyWorldsResponse []worldViewModel

type createWorldResponse worldViewModel

type updateWorldResponse worldViewModel
