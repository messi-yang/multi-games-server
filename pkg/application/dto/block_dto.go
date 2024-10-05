package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/blockmodel"
)

type BlockDto struct {
	Id BlockIdDto `json:"id"`
}

func NewBlockDto(block blockmodel.Block) BlockDto {
	return BlockDto{
		Id: BlockIdDto{
			WorldId: block.GetId().GetWorldId().Uuid(),
			X:       block.GetX(),
			Z:       block.GetZ(),
		},
	}
}
