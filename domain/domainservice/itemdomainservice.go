package domainservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/itemmodel"
)

type ItemDomainService interface {
	GetAllItems() []itemmodel.Item
}

type itemDomainServe struct{}

func NewItemDomainServe() ItemDomainService {
	return &itemDomainServe{}
}

func (serve *itemDomainServe) GetAllItems() []itemmodel.Item {
	stoneItemDefaultId, _ := itemmodel.NewItemId("4632b3c0-f748-4c46-954a-93a5cb4bc767")
	woodItemDefaultId, _ := itemmodel.NewItemId("7ab5bc4e-9596-4a54-adee-8807512cbbb4")
	sandItemDefaultId, _ := itemmodel.NewItemId("3e8a5704-6de6-4156-96ea-25076aa82b35")
	mudItemDefaultId, _ := itemmodel.NewItemId("6c2da2d6-d47b-4bc3-b884-889f9b7ba882")
	glassItemDefaultId, _ := itemmodel.NewItemId("ec112b60-826c-438b-85eb-d5e3b2a428b5")
	steelItemDefaultId, _ := itemmodel.NewItemId("f462b570-b02d-437a-8b10-563fec84ee96")

	return []itemmodel.Item{
		itemmodel.NewItem(stoneItemDefaultId, "stone"),
		itemmodel.NewItem(woodItemDefaultId, "wood"),
		itemmodel.NewItem(sandItemDefaultId, "sand"),
		itemmodel.NewItem(mudItemDefaultId, "mud"),
		itemmodel.NewItem(glassItemDefaultId, "glass"),
		itemmodel.NewItem(steelItemDefaultId, "steel"),
	}
}
