package itemappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/itemviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
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
	presenter.OnSuccess(itemviewmodel.BatchNew(items))
}
