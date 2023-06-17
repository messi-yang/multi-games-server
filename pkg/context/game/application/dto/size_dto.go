package dto

import "github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commonmodel"

type SizeDto struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func NewSizeDto(size commonmodel.Size) SizeDto {
	return SizeDto{
		Width:  size.GetWidth(),
		Height: size.GetHeight(),
	}
}
