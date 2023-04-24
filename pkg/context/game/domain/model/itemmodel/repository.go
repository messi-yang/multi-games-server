package itemmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"

type Repo interface {
	GetAll() ([]ItemAgg, error)
	Get(itemId commonmodel.ItemIdVo) (ItemAgg, error)
	GetFirstItem() (ItemAgg, error)
	Add(item ItemAgg) error
	Update(item ItemAgg) error
}
