package itemappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/samber/lo"
)

type Service interface {
	GetAllItems(presenter Presenter)
}

type serve struct {
	itemRepo itemmodel.Repo
}

func New(itemRepo itemmodel.Repo) Service {
	return &serve{itemRepo: itemRepo}
}

func (serve *serve) GetAllItems(presenter Presenter) {
	items := serve.itemRepo.GetAllItems()
	itemCameraModels := lo.Map(items, func(item itemmodel.ItemAgr, _ int) viewmodel.ItemVm {
		return viewmodel.NewItemVm(item)
	})
	presenter.OnSuccess(itemCameraModels)
}
