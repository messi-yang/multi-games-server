package itemmodel

type Item struct {
	id       ItemId
	name     string
	assetSrc string
}

func New(id ItemId, name string, assetSrc string) Item {
	return Item{
		id:       id,
		name:     name,
		assetSrc: assetSrc,
	}
}

func (item *Item) GetId() ItemId {
	return item.id
}

func (item *Item) GetName() string {
	return item.name
}

func (item *Item) GetAssetSrc() string {
	return item.assetSrc
}
