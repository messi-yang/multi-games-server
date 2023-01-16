package itemmodel

type ItemAgr struct {
	id       ItemIdVo
	name     string
	assetSrc string
}

func NewItemAgr(id ItemIdVo, name string, assetSrc string) ItemAgr {
	return ItemAgr{
		id:       id,
		name:     name,
		assetSrc: assetSrc,
	}
}

func (item *ItemAgr) GetId() ItemIdVo {
	return item.id
}

func (item *ItemAgr) GetName() string {
	return item.name
}

func (item *ItemAgr) GetAssetSrc() string {
	return item.assetSrc
}
