package itemappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/itemviewmodel"
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
	itemViewModels := lo.Map(items, func(item itemmodel.Item, _ int) itemviewmodel.ViewModel {
		return itemviewmodel.New(item)
	})
	presenter.OnSuccess(itemViewModels)
}
