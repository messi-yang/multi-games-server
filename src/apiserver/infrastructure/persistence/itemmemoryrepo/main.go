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
		stoneItemDefaultId, _ := itemmodel.NewItemId("2a0c8f7f-48dc-4553-86b4-2bbc3786bb66")
		torchItemDefaultId, _ := itemmodel.NewItemId("31c8cc9e-42a3-4a42-86d2-905ca37305ba")

		fmt.Println(uuid.New())
		fmt.Println(uuid.New())
		fmt.Println(uuid.New())

		serverUrl := os.Getenv("SERVER_URL")

		singleton = &memoryRepo{
			items: []itemmodel.Item{
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
