package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/blockmodel"
	"github.com/google/uuid"
)

type BlockIdDto struct {
	WorldId uuid.UUID `json:"worldId"`
	X       int       `json:"x"`
	Z       int       `json:"z"`
}

func NewBlockIdDto(blockId blockmodel.BlockId) BlockIdDto {
	return BlockIdDto{
		WorldId: blockId.GetWorldId().Uuid(),
		X:       blockId.GetX(),
		Z:       blockId.GetZ(),
	}
}

func (dto BlockIdDto) ToValueObject() blockmodel.BlockId {
	return blockmodel.NewBlockId(
		globalcommonmodel.NewWorldId(dto.WorldId), dto.X, dto.Z,
	)
}
