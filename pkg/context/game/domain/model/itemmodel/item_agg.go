package itemmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"

type ItemAgg struct {
	id           commonmodel.ItemIdVo
	name         string
	traversable  bool
	thumbnailSrc string
	modelSrc     string
}

func NewItemAgg(id commonmodel.ItemIdVo, name string, traversable bool, thumbnailSrc string, modelSrc string) ItemAgg {
	return ItemAgg{
		id:           id,
		name:         name,
		traversable:  traversable,
		thumbnailSrc: thumbnailSrc,
		modelSrc:     modelSrc,
	}
}

func (item *ItemAgg) GetId() commonmodel.ItemIdVo {
	return item.id
}

func (item *ItemAgg) GetName() string {
	return item.name
}

func (item *ItemAgg) GetTraversable() bool {
	return item.traversable
}

func (item *ItemAgg) GetThumbnailSrc() string {
	return item.thumbnailSrc
}

func (item *ItemAgg) GetModelSrc() string {
	return item.modelSrc
}
