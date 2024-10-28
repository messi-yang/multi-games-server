package globalcommonmodel

import (
	"fmt"
	"strconv"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type Color struct {
	r, g, b int64
}

// Interface Implementation Check
var _ domain.ValueObject[Color] = (*Color)(nil)

func NewColor(r, g, b int64) (color Color, err error) {
	if !isValidColorComponent(r) || !isValidColorComponent(g) || !isValidColorComponent(b) {
		return color, fmt.Errorf("color components must be between 0 and 255")
	}
	return Color{
		r: r,
		g: g,
		b: b,
	}, nil
}

func NewColorFromHexString(hexString string) (color Color, err error) {
	if len(hexString) != 7 || hexString[0] != '#' {
		return color, fmt.Errorf("invalid hex color format")
	}

	r, err := strconv.ParseInt(hexString[1:3], 16, 0)
	if err != nil {
		return color, err
	}
	g, err := strconv.ParseInt(hexString[3:5], 16, 0)
	if err != nil {
		return color, err
	}
	b, err := strconv.ParseInt(hexString[5:7], 16, 0)
	if err != nil {
		return color, err
	}

	return Color{
		r: r,
		g: g,
		b: b,
	}, nil
}

func (color Color) IsEqual(otherColor Color) bool {
	return color == otherColor
}

func (color Color) HexString() string {
	return fmt.Sprintf("#%s%s%s", fmt.Sprintf("%02x", color.r), fmt.Sprintf("%02x", color.g), fmt.Sprintf("%02x", color.b))
}

func isValidColorComponent(value int64) bool {
	return value >= 0 && value <= 255
}
