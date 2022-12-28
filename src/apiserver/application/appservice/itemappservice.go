package appservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainservice"
)

type ItemAppService interface {
	GetAllItems() []itemmodel.Item
}

type itemAppServe struct {
	itemDomainService domainservice.ItemDomainService
}

func NewItemAppService(itemDomainService domainservice.ItemDomainService) ItemAppService {
	return &itemAppServe{itemDomainService: itemDomainService}
}

func (serve *itemAppServe) GetAllItems() []itemmodel.Item {
	return serve.itemDomainService.GetAllItems()
}
