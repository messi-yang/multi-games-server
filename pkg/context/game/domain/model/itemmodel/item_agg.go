package itemmodel

type ItemAgg struct {
	id           ItemIdVo
	name         string
	traversable  bool
	thumbnailSrc string
	modelSrc     string
}

func NewItemAgg(id ItemIdVo, name string, traversable bool, thumbnailSrc string, modelSrc string) ItemAgg {
	return ItemAgg{
		id:           id,
		name:         name,
		traversable:  traversable,
		thumbnailSrc: thumbnailSrc,
		modelSrc:     modelSrc,
	}
}

func (item *ItemAgg) GetId() ItemIdVo {
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
