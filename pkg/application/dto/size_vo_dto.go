package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"

type SizeVoDto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewSizeVoDto(size commonmodel.SizeVo) SizeVoDto {
	return SizeVoDto{
		Width:  size.GetWidth(),
		Height: size.GetHeight(),
	}
}
