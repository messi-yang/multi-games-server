package service

import "fmt"

var (
	errUnitExceededBoundary         = fmt.Errorf("unit exceeds the boundary")
	errPositionAlreadyHasUnit       = fmt.Errorf("the position already has unit")
	errUnitCannotBeAtOriginPosition = fmt.Errorf("cannot put unit at origin position")
)
