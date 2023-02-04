package appservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/samber/lo"
)

type ItemAppService interface {
	GetAllItems(presenter Presenter)
}

type itemAppServe struct {
	itemRepo itemmodel.Repo
}

func NewItemAppService(itemRepo itemmodel.Repo) ItemAppService {
	return &itemAppServe{itemRepo: itemRepo}
}

func (itemAppServe *itemAppServe) GetAllItems(presenter Presenter) {
	items := itemAppServe.itemRepo.GetAll()
	itemModels := lo.Map(items, func(item itemmodel.ItemAgg, _ int) viewmodel.ItemVm {
		return viewmodel.NewItemVm(item)
	})
	presenter.OnSuccess(itemModels)
}
