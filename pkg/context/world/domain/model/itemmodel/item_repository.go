package itemmodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"

type ItemRepo interface {
	Add(item Item) error
	Update(item Item) error
	Get(itemId worldcommonmodel.ItemId) (Item, error)
	GetAll() ([]Item, error)
	GetItemsWithIds(itemIds []worldcommonmodel.ItemId) ([]Item, error)
	GetItemsOfCompatibleUnitType(compatibleUnitType worldcommonmodel.UnitType) ([]Item, error)
	GetFirstItem() (Item, error)
}
