package itemservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/itemmodel"
	"github.com/google/uuid"
)

type ItemService interface {
	GetAllItems() []itemmodel.Item
}

type itemServe struct{}

func NewItemServe() ItemService {
	return &itemServe{}
}

func (serve *itemServe) GetAllItems() []itemmodel.Item {
	stoneItemDefaultId, _ := uuid.Parse("4632b3c0-f748-4c46-954a-93a5cb4bc767")
	// fmt.Println(uuid.New())

	return []itemmodel.Item{
		itemmodel.NewItem(itemmodel.NewItemId(stoneItemDefaultId), "stone"),
	}
}
