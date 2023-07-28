package viewmodel

import (
	iam_dto "github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/dto"
	world_dto "github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
)

type WorldViewModel struct {
	world_dto.WorldDto
	UserDto iam_dto.UserDto `json:"user"`
}
