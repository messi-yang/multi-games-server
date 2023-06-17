package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/itemmodel"
	"github.com/google/uuid"
)

type ItemDto struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Traversable bool      `json:"traversable"`
	AssetSrc    string    `json:"thumbnailSrc"`
	ModelSrc    string    `json:"modelSrc"`
}

func NewItemDto(item itemmodel.Item) ItemDto {
	return ItemDto{
		Id:          item.GetId().Uuid(),
		Name:        item.GetName(),
		Traversable: item.GetTraversable(),
		AssetSrc:    item.GetThumbnailSrc(),
		ModelSrc:    item.GetModelSrc(),
	}
}
