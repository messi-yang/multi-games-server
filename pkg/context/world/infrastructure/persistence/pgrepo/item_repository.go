package pgrepo

import (
	"fmt"
	"os"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

func newModelFromItem(item itemmodel.Item) pgmodel.ItemModel {
	return pgmodel.ItemModel{
		Id:                 item.GetId().Uuid(),
		CompatibleUnitType: pgmodel.UnitTypeEnum(item.GetCompatibleUnitType().String()),
		Name:               item.GetName(),
		Traversable:        item.GetTraversable(),
		ThumbnailSrc:       item.GetThumbnailSrc(),
		ModelSrc:           item.GetModelSrc(),
	}
}

func parseModelToItem(itemModel pgmodel.ItemModel) (item itemmodel.Item, err error) {
	serverUrl := os.Getenv("SERVER_URL")
	compatibleUnitType, err := worldcommonmodel.NewUnitType(string(itemModel.CompatibleUnitType))
	if err != nil {
		return item, err
	}
	return itemmodel.LoadItem(
		worldcommonmodel.NewItemId(itemModel.Id),
		compatibleUnitType,
		itemModel.Name,
		itemModel.Traversable,
		fmt.Sprintf("%s%s", serverUrl, itemModel.ThumbnailSrc),
		fmt.Sprintf("%s%s", serverUrl, itemModel.ModelSrc),
	), nil
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
	itemModel := newModelFromItem(item)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&itemModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&item)
}

func (repo *itemRepo) Update(item itemmodel.Item) error {
	itemModel := newModelFromItem(item)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Model(&pgmodel.ItemModel{}).Where(
			"id = ?",
			item.GetId().Uuid(),
		).Select("*").Updates(&itemModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&item)
}

func (repo *itemRepo) Get(itemId worldcommonmodel.ItemId) (item itemmodel.Item, err error) {
	itemModel := pgmodel.ItemModel{Id: itemId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&itemModel).Error
	}); err != nil {
		return item, err
	}

	return parseModelToItem(itemModel)
}

func (repo *itemRepo) GetAll() (items []itemmodel.Item, err error) {
	var itemModels []pgmodel.ItemModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Find(&itemModels).Error
	}); err != nil {
		return items, err
	}

	return commonutil.MapWithError(itemModels, func(_ int, itemModel pgmodel.ItemModel) (itemmodel.Item, error) {
		return parseModelToItem(itemModel)
	})
}

func (repo *itemRepo) GetItemsOfIds(itemIds []worldcommonmodel.ItemId) (items []itemmodel.Item, err error) {
	itemIdDtos := lo.Map(itemIds, func(itemId worldcommonmodel.ItemId, _ int) uuid.UUID {
		return itemId.Uuid()
	})
	var itemModels []pgmodel.ItemModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(itemIdDtos).Find(&itemModels).Error
	}); err != nil {
		return items, err
	}

	return commonutil.MapWithError(itemModels, func(_ int, itemModel pgmodel.ItemModel) (itemmodel.Item, error) {
		return parseModelToItem(itemModel)
	})
}

func (repo *itemRepo) GetItemsOfCompatibleUnitType(compatibleUnitType worldcommonmodel.UnitType) (items []itemmodel.Item, err error) {
	var itemModels []pgmodel.ItemModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"compatible_unit_type = ?",
			compatibleUnitType.String(),
		).Find(&itemModels, pgmodel.UnitModel{}).Error
	}); err != nil {
		return items, err
	}

	return commonutil.MapWithError(itemModels, func(_ int, itemModel pgmodel.ItemModel) (itemmodel.Item, error) {
		return parseModelToItem(itemModel)
	})
}

func (repo *itemRepo) GetFirstItem() (item itemmodel.Item, err error) {
	itemModel := pgmodel.ItemModel{}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&itemModel).Error
	}); err != nil {
		return item, err
	}

	return parseModelToItem(itemModel)
}
