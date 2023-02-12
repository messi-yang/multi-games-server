package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"

type SizeDto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewSizeDto(size commonmodel.SizeVo) SizeDto {
	return SizeDto{
		Width:  size.GetWidth(),
		Height: size.GetHeight(),
	}
}
