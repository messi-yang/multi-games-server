package pgrepo

import (
	"fmt"
	"os"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/samber/lo"
)

func newItemModel(item itemmodel.Item) pgmodel.ItemModel {
	return pgmodel.ItemModel{
		Id:           item.GetId().Uuid(),
		Name:         item.GetName(),
		Traversable:  item.GetTraversable(),
		ThumbnailSrc: item.GetThumbnailSrc(),
		ModelSrc:     item.GetModelSrc(),
	}
}

func parseItemModel(itemModel pgmodel.ItemModel) itemmodel.Item {
	serverUrl := os.Getenv("SERVER_URL")
	return itemmodel.NewItem(
		commonmodel.NewItemId(itemModel.Id),
		itemModel.Name,
		itemModel.Traversable,
		fmt.Sprintf("%s%s", serverUrl, itemModel.ThumbnailSrc),
		fmt.Sprintf("%s%s", serverUrl, itemModel.ModelSrc),
	)
}

type itemRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewItemRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository itemmodel.ItemRepo) {
	return &itemRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *itemRepo) Add(item itemmodel.Item) error {
	itemModel := newItemModel(item)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&itemModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&item)
}

func (repo *itemRepo) Update(item itemmodel.Item) error {
	itemModel := newItemModel(item)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Save(&itemModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&item)
}

func (repo *itemRepo) Get(itemId commonmodel.ItemId) (item itemmodel.Item, err error) {
	itemModel := pgmodel.ItemModel{Id: itemId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&itemModel).Error
	}); err != nil {
		return item, err
	}

	return parseItemModel(itemModel), nil
}

func (repo *itemRepo) GetAll() (items []itemmodel.Item, err error) {
	var itemModels []pgmodel.ItemModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Find(&itemModels).Error
	}); err != nil {
		return items, err
	}

	items = lo.Map(itemModels, func(itemModel pgmodel.ItemModel, _ int) itemmodel.Item {
		return parseItemModel(itemModel)
	})
	return items, nil
}

func (repo *itemRepo) GetFirstItem() (item itemmodel.Item, err error) {
	itemModel := pgmodel.ItemModel{}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&itemModel).Error
	}); err != nil {
		return item, err
	}

	return parseItemModel(itemModel), nil
}
