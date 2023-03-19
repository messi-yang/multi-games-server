package seedcmd

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/uuidutil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres"
	"github.com/samber/lo"
)

func Exec() {
	fmt.Println("Start seeding Postgres database")

	itemRepository, err := postgres.NewItemRepository()
	if err != nil {
		return
	}

	items := []itemmodel.ItemAgg{
		itemmodel.NewItemAgg(
			itemmodel.NewItemIdVo(uuidutil.UnsafelyNewUuid("3c28537a-80c2-4ac1-917b-b1cd517c6b5e")),
			"stone", false, "/asset/item/stone/thumbnail.png", "/asset/item/stone/model.gltf",
		),
		itemmodel.NewItemAgg(
			itemmodel.NewItemIdVo(uuidutil.UnsafelyNewUuid("34af14ab-42c5-4c55-a787-44f32012354e")),
			"torch", true, "/asset/item/torch/thumbnail.png", "/asset/item/torch/model.gltf",
		),
		itemmodel.NewItemAgg(
			itemmodel.NewItemIdVo(uuidutil.UnsafelyNewUuid("414b5703-91d1-42fc-a007-36dd8f25e329")),
			"tree", false, "/asset/item/tree/thumbnail.png", "/asset/item/tree/model.gltf",
		),
		itemmodel.NewItemAgg(
			itemmodel.NewItemIdVo(uuidutil.UnsafelyNewUuid("41de86e6-07a1-4a5d-ba6f-152d07f3ba1e")),
			"fan", false, "/asset/item/fan/thumbnail.png", "/asset/item/fan/model.gltf",
		),
	}

	lo.ForEach(items, func(item itemmodel.ItemAgg, _ int) {
		if _, err = itemRepository.Get(item.GetId()); err != nil {
			fmt.Printf("Add new item \"%s\"\n", item.GetName())
			itemRepository.Add(item)
		} else {
			fmt.Printf("Update existing item \"%s\"\n", item.GetName())
			itemRepository.Update(item)
		}
	})

	fmt.Println("Finished seeding Postgres database")
}
