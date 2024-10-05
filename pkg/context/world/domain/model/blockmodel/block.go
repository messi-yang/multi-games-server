package blockmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

// A block has a fixed dimension, and its x and z indicate where
// it is on the map.
// If a block is 10x10 and its at (3, 7), the bound of the block will be
// from (3 * 10, 7 * 10) to (3 + 1) * 10 - 1, (7 + 1) * 10 - 1), which is from (30, 70) to (39, 79)
type Block struct {
	blockId BlockId
}

// Interface Implementation Check
var _ domain.Aggregate = (*Block)(nil)

func LoadBlock(blockId BlockId) Block {
	return Block{
		blockId: blockId,
	}
}

func (block Block) GetId() BlockId {
	return block.blockId
}

func (block Block) GetX() int {
	return block.blockId.GetX()
}

func (block Block) GetZ() int {
	return block.blockId.GetZ()
}

func (block Block) GetDimension() worldcommonmodel.Dimension {
	dimension, _ := worldcommonmodel.NewDimension(50, 50)
	return dimension
}

func (block Block) GetBound() worldcommonmodel.Bound {
	dimension := block.GetDimension()
	dimensionWidth := dimension.GetWidth()
	dimensionDepth := dimension.GetDepth()

	x := block.GetX()
	z := block.GetZ()

	bound, _ := worldcommonmodel.NewBound(worldcommonmodel.NewPosition(x*dimensionWidth, z*dimensionDepth), worldcommonmodel.NewPosition((x+1)*dimensionWidth-1, (z+1)*dimensionDepth-1))
	return bound
}
