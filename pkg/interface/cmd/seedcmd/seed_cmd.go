package seedcmd

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres"
)

func Exec() {
	fmt.Println("Start seeding Postgres database")

	itemRepository, err := postgres.NewItemRepository()
	if err != nil {
		return
	}

	itemId, _ := itemmodel.ParseItemIdVo("3c28537a-80c2-4ac1-917b-b1cd517c6b5e")
	if _, err = itemRepository.Get(itemId); err != nil {
		itemRepository.Add(itemmodel.NewItemAgg(itemId, "stone", false, "/asset/item/stone/thumbnail.png", "/asset/item/stone/model.gltf"))
	} else {
		fmt.Println("Item with id of 3c28537a-80c2-4ac1-917b-b1cd517c6b5e found")
	}

	itemId, _ = itemmodel.ParseItemIdVo("34af14ab-42c5-4c55-a787-44f32012354e")
	if _, err = itemRepository.Get(itemId); err != nil {
		itemRepository.Add(itemmodel.NewItemAgg(itemId, "torch", true, "/asset/item/torch/thumbnail.png", "/asset/item/torch/model.gltf"))
	} else {
		fmt.Println("Item with id of 34af14ab-42c5-4c55-a787-44f32012354e found")
	}

	itemId, _ = itemmodel.ParseItemIdVo("414b5703-91d1-42fc-a007-36dd8f25e329")
	if _, err = itemRepository.Get(itemId); err != nil {
		itemRepository.Add(itemmodel.NewItemAgg(itemId, "tree", false, "/asset/item/tree/thumbnail.png", "/asset/item/tree/model.gltf"))
	} else {
		fmt.Println("Item with id of 414b5703-91d1-42fc-a007-36dd8f25e329 found")
	}

	fmt.Println("Finished seeding Postgres database")
}
