package service

import "fmt"

var (
	errBoundAlreadyHasUnit          = fmt.Errorf("the bound already has unit")
	errUnitCannotBeAtOriginPosition = fmt.Errorf("cannot put unit at origin position")
)
