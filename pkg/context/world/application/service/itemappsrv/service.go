package itemappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/uuidutil"
	"github.com/samber/lo"
)

type Service interface {
	QueryItems(QueryItemsQuery) ([]dto.ItemDto, error)
	CreateDefaultItems(CreateDefaultItemsCommand) error
}

type serve struct {
	itemRepo itemmodel.ItemRepo
}

func NewService(itemRepo itemmodel.ItemRepo) Service {
	return &serve{
		itemRepo: itemRepo,
	}
}

func (serve *serve) QueryItems(query QueryItemsQuery) (itemDtos []dto.ItemDto, err error) {
	items, err := serve.itemRepo.GetAll()
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(items, func(item itemmodel.Item, _ int) dto.ItemDto {
		return dto.NewItemDto(item)
	}), nil
}

func (serve *serve) CreateDefaultItems(command CreateDefaultItemsCommand) error {
	items := []itemmodel.Item{
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("6f127ae8-f1f8-4ff3-8148-fa8d2fef307a")),
			worldcommonmodel.NewStaticUnitType(),
			"test", true, "/asset/item/test/thumbnail.png", "/asset/item/test/model.gltf",
		),
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("3c28537a-80c2-4ac1-917b-b1cd517c6b5e")),
			worldcommonmodel.NewStaticUnitType(),
			"stone", false, "/asset/item/stone/thumbnail.png", "/asset/item/stone/model.gltf",
		),
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("34af14ab-42c5-4c55-a787-44f32012354e")),
			worldcommonmodel.NewStaticUnitType(),
			"torch", true, "/asset/item/torch/thumbnail.png", "/asset/item/torch/model.gltf",
		),
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("414b5703-91d1-42fc-a007-36dd8f25e329")),
			worldcommonmodel.NewStaticUnitType(),
			"tree", false, "/asset/item/tree/thumbnail.png", "/asset/item/tree/model.gltf",
		),
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("41de86e6-07a1-4a5d-ba6f-152d07f3ba1e")),
			worldcommonmodel.NewStaticUnitType(),
			"fan", false, "/asset/item/fan/thumbnail.png", "/asset/item/fan/model.gltf",
		),
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("c0a15d4a-24b7-4a81-8a39-9bbf4c7d6ccf")),
			worldcommonmodel.NewStaticUnitType(),
			"grass", true, "/asset/item/grass/thumbnail.png", "/asset/item/grass/model.gltf",
		),
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("2b6ab30d-0a2a-4424-b245-99ec2c301844")),
			worldcommonmodel.NewStaticUnitType(),
			"chair 1", false, "/asset/item/chair/thumbnail.png", "/asset/item/chair/model.gltf",
		),
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("fb9d06f8-5d6d-4fa9-bdc5-ab760d55a442")),
			worldcommonmodel.NewStaticUnitType(),
			"potted plant", false, "/asset/item/potted_plant/thumbnail.png", "/asset/item/potted_plant/model.gltf",
		),
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("1b9ba8b1-c13e-4524-bddc-7cc6d981ee2c")),
			worldcommonmodel.NewStaticUnitType(),
			"trash bin", false, "/asset/item/trash_bin/thumbnail.png", "/asset/item/trash_bin/model.gltf",
		),
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("e495468b-e662-49cb-bc5b-96db204ad9d8")),
			worldcommonmodel.NewStaticUnitType(),
			"box", false, "/asset/item/box/thumbnail.png", "/asset/item/box/model.gltf",
		),
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("52bdd7d3-799d-42dd-a2dc-cd438101cfca")),
			worldcommonmodel.NewStaticUnitType(),
			"chair 2", false, "/asset/item/chair_2/thumbnail.png", "/asset/item/chair_2/model.gltf",
		),
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("d4d0850a-dbe0-451c-9e50-6ac280108d1c")),
			worldcommonmodel.NewStaticUnitType(),
			"fence", false, "/asset/item/fence/thumbnail.png", "/asset/item/fence/model.gltf",
		),
		itemmodel.LoadItem(
			worldcommonmodel.NewItemId(uuidutil.UnsafelyNewUuid("c64f732c-9494-4693-b17b-43a2736aa67b")),
			worldcommonmodel.NewPortalUnitType(),
			"portal", true, "/asset/item/portal/thumbnail.png", "/asset/item/portal/model.gltf",
		),
	}

	for _, item := range items {
		if _, err := serve.itemRepo.Get(item.GetId()); err != nil {
			fmt.Printf("Add new item \"%s\"\n", item.GetName())
			if err = serve.itemRepo.Add(item); err != nil {
				return err
			}
		} else {
			fmt.Printf("Update existing item \"%s\"\n", item.GetName())
			if err = serve.itemRepo.Update(item); err != nil {
				return err
			}
		}
	}

	return nil
}
