package worldcommonmodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"

// A block has a fixed dimension, and its x and z indicate where
// it is on the map.
// If a block is 10x10 and its at (3, 7), the bound of the block will be
// from (3 * 10, 7 * 10) to (3 + 1) * 10 - 1, (7 + 1) * 10 - 1), which is from (30, 70) to (39, 79)
type Block struct {
	x int
	z int
}

// Interface Implementation Check
var _ domain.ValueObject[Block] = (*Block)(nil)

func NewBlock(x int, z int) Block {
	return Block{
		x: x,
		z: z,
	}
}

func (block Block) IsEqual(otherBlock Block) bool {
	return block.x == otherBlock.x && block.z == otherBlock.z
}

func (block Block) GetX() int {
	return block.x
}

func (block Block) GetZ() int {
	return block.z
}

func (block Block) GetDimension() Dimension {
	dimension, _ := NewDimension(50, 50)
	return dimension
}

func (block Block) GetBound() Bound {
	dimension := block.GetDimension()
	dimensionWidth := dimension.GetWidth()
	dimensionDepth := dimension.GetDepth()

	bound, _ := NewBound(NewPosition(block.x*dimensionWidth, block.z*dimensionDepth), NewPosition((block.x+1)*dimensionWidth-1, (block.z+1)*dimensionDepth-1))
	return bound
}
