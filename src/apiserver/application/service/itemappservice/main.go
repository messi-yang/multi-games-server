package itemappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type Service interface {
	GetAllItems() []itemmodel.Item
}

type serve struct {
	itemRepo itemmodel.Repo
}

func New(itemRepo itemmodel.Repo) Service {
	return &serve{itemRepo: itemRepo}
}

func (serve *serve) GetAllItems() []itemmodel.Item {
	return serve.itemRepo.GetAllItems()
}
