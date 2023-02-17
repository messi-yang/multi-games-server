package itemmodel

type ItemAgg struct {
	id          ItemIdVo
	name        string
	traversable bool
	assetSrc    string
	modelSrc    string
}

func NewItemAgg(id ItemIdVo, name string, traversable bool, assetSrc string, modelSrc string) ItemAgg {
	return ItemAgg{
		id:          id,
		name:        name,
		traversable: traversable,
		assetSrc:    assetSrc,
		modelSrc:    modelSrc,
	}
}

func (item *ItemAgg) GetId() ItemIdVo {
	return item.id
}

func (item *ItemAgg) GetName() string {
	return item.name
}

func (item *ItemAgg) IsTraversable() bool {
	return item.traversable
}

func (item *ItemAgg) GetAssetSrc() string {
	return item.assetSrc
}

func (item *ItemAgg) GetModelSrc() string {
	return item.modelSrc
}
