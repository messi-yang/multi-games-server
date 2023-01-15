package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type SizeVm struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewSizeVm(size commonmodel.Size) SizeVm {
	return SizeVm{
		Width:  size.GetWidth(),
		Height: size.GetHeight(),
	}
}

func (dto SizeVm) ToValueObject() (commonmodel.Size, error) {
	return commonmodel.NewSize(dto.Width, dto.Height)
}
