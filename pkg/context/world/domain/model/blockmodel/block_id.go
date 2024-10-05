package blockmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type BlockId struct {
	worldId globalcommonmodel.WorldId
	x       int
	z       int
}

// Interface Implementation Check
var _ domain.ValueObject[BlockId] = (*BlockId)(nil)

func NewBlockId(worldId globalcommonmodel.WorldId, x int, z int) BlockId {
	return BlockId{
		worldId: worldId,
		x:       x,
		z:       z,
	}
}

func (blockId BlockId) IsEqual(otherBlockId BlockId) bool {
	return blockId.x == otherBlockId.x && blockId.z == otherBlockId.z
}

func (blockId BlockId) GetWorldId() globalcommonmodel.WorldId {
	return blockId.worldId
}

func (blockId BlockId) GetX() int {
	return blockId.x
}

func (blockId BlockId) GetZ() int {
	return blockId.z
}
