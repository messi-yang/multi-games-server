package itemmemoryrepo

import (
	"fmt"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/google/uuid"
)

type memoryRepo struct {
	items []itemmodel.Item
}

var singleton *memoryRepo

func New() itemmodel.Repo {
	if singleton == nil {
		appleItemDefaultId, _ := itemmodel.NewItemId("4632b3c0-f748-4c46-954a-93a5cb4bc767")
		bananaItemDefaultId, _ := itemmodel.NewItemId("7ab5bc4e-9596-4a54-adee-8807512cbbb4")
		orangeItemDefaultId, _ := itemmodel.NewItemId("3e8a5704-6de6-4156-96ea-25076aa82b35")
		waterMelonItemDefaultId, _ := itemmodel.NewItemId("6c2da2d6-d47b-4bc3-b884-889f9b7ba882")
		stoneItemDefaultId, _ := itemmodel.NewItemId("2a0c8f7f-48dc-4553-86b4-2bbc3786bb66")
		torchItemDefaultId, _ := itemmodel.NewItemId("31c8cc9e-42a3-4a42-86d2-905ca37305ba")

		fmt.Println(uuid.New())
		fmt.Println(uuid.New())
		fmt.Println(uuid.New())

		serverUrl := os.Getenv("SERVER_URL")

		singleton = &memoryRepo{
			items: []itemmodel.Item{
				itemmodel.New(appleItemDefaultId, "apple", fmt.Sprintf("%s/assets/items/apple.png", serverUrl)),
				itemmodel.New(bananaItemDefaultId, "banana", fmt.Sprintf("%s/assets/items/banana.png", serverUrl)),
				itemmodel.New(orangeItemDefaultId, "orange", fmt.Sprintf("%s/assets/items/orange.png", serverUrl)),
				itemmodel.New(waterMelonItemDefaultId, "water melon", fmt.Sprintf("%s/assets/items/water-melon.png", serverUrl)),
				itemmodel.New(stoneItemDefaultId, "stone", fmt.Sprintf("%s/assets/items/stone.png", serverUrl)),
				itemmodel.New(torchItemDefaultId, "torch", fmt.Sprintf("%s/assets/items/torch.png", serverUrl)),
			},
		}
		return singleton
	}
	return singleton
}

func (repo *memoryRepo) GetAllItems() []itemmodel.Item {
	return repo.items
}
