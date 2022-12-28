package itemappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/service/itemdomainservice"
)

type Service interface {
	GetAllItems() []itemmodel.Item
}

type serve struct {
	itemDomainService itemdomainservice.Service
}

func New(itemDomainService itemdomainservice.Service) Service {
	return &serve{itemDomainService: itemDomainService}
}

func (serve *serve) GetAllItems() []itemmodel.Item {
	return serve.itemDomainService.GetAllItems()
}
