package itemmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"

type ItemRepo interface {
	Add(item Item) error
	Update(item Item) error
	Get(itemId commonmodel.ItemId) (Item, error)
	GetAll() ([]Item, error)
	GetFirstItem() (Item, error)
}
