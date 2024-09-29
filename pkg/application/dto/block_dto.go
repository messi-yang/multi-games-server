package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type BlockDto struct {
	X int `json:"x"`
	Z int `json:"z"`
}

func NewBlockDto(block worldcommonmodel.Block) BlockDto {
	return BlockDto{
		X: block.GetX(),
		Z: block.GetZ(),
	}
}

func (dto BlockDto) ToValueObject() (block worldcommonmodel.Block) {
	return worldcommonmodel.NewBlock(
		dto.X,
		dto.Z,
	)
}
