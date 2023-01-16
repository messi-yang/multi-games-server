package commonmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type UnitVo struct {
	itemId itemmodel.ItemIdVo
}

func NewUnitVo(itemId itemmodel.ItemIdVo) UnitVo {
	return UnitVo{
		itemId: itemId,
	}
}

func (gu UnitVo) GetItemId() itemmodel.ItemIdVo {
	return gu.itemId
}

func (gu UnitVo) SetItemId(itemId itemmodel.ItemIdVo) UnitVo {
	return UnitVo{
		itemId: itemId,
	}
}
