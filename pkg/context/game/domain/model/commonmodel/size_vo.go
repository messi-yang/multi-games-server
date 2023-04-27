package commonmodel

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/valueobject"
)

type ErrInvalidSizeVo struct {
	width  int
	height int
}

func (e *ErrInvalidSizeVo) Error() string {
	return fmt.Sprintf("width (%d) and height(%d) of size must be greater than 0", e.width, e.height)
}

type SizeVo struct {
	width  int
	height int
}

// Interface Implementation Check
var _ valueobject.ValueObject[SizeVo] = (*SizeVo)(nil)

func NewSizeVo(width int, height int) (SizeVo, error) {
	if width < 1 || height < 1 {
		return SizeVo{}, &ErrInvalidSizeVo{width: width, height: height}
	}

	return SizeVo{
		width:  width,
		height: height,
	}, nil
}

func (size SizeVo) IsEqual(anotherVo SizeVo) bool {
	return size.width == anotherVo.width && size.height == anotherVo.height
}

func (size SizeVo) GetWidth() int {
	return size.width
}

func (size SizeVo) GetHeight() int {
	return size.height
}
