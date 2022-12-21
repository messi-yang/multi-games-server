package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/service/itemservice"
)

type ItemAppService interface {
	GetAllItems() []itemmodel.Item
}

type itemAppServe struct {
	itemService itemservice.ItemService
}

func NewItemAppService(itemService itemservice.ItemService) ItemAppService {
	return &itemAppServe{itemService: itemService}
}

func (serve *itemAppServe) GetAllItems() []itemmodel.Item {
	return serve.itemService.GetAllItems()
}
