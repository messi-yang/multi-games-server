package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type SizeVm struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewSizeVm(size commonmodel.SizeVo) SizeVm {
	return SizeVm{
		Width:  size.GetWidth(),
		Height: size.GetHeight(),
	}
}
