package itemmodel

type ItemAgg struct {
	id       ItemIdVo
	name     string
	assetSrc string
}

func NewItemAgg(id ItemIdVo, name string, assetSrc string) ItemAgg {
	return ItemAgg{
		id:       id,
		name:     name,
		assetSrc: assetSrc,
	}
}

func (item *ItemAgg) GetId() ItemIdVo {
	return item.id
}

func (item *ItemAgg) GetName() string {
	return item.name
}

func (item *ItemAgg) GetAssetSrc() string {
	return item.assetSrc
}
