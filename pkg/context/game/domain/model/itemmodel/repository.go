package itemmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"

type Repo interface {
	Add(item Item) error
	Get(itemId commonmodel.ItemId) (Item, error)
	Update(item Item) error
	GetAll() ([]Item, error)
	GetFirstItem() (Item, error)
}
