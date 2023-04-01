package itemappservice

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"

type Service interface {
	GetItems(GetItemsQuery) ([]itemmodel.ItemAgg, error)
}

type serve struct {
	itemRepository itemmodel.Repository
}

func NewService(itemRepository itemmodel.Repository) Service {
	return &serve{
		itemRepository: itemRepository,
	}
}

func (serve *serve) GetItems(query GetItemsQuery) ([]itemmodel.ItemAgg, error) {
	return serve.itemRepository.GetAll()
}
