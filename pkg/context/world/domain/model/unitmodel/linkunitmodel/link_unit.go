package linkunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type LinkUnit struct {
	unitmodel.UnitEntity
	url globalcommonmodel.Url
}

// Interface Implementation Check
var _ domain.Aggregate = (*LinkUnit)(nil)

func NewLinkUnit(
	id LinkUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
	label *string,
	url globalcommonmodel.Url,
) LinkUnit {
	return LinkUnit{
		UnitEntity: unitmodel.NewUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			dimension,
			label,
			nil,
			worldcommonmodel.NewLinkUnitType(),
		),
		url: url,
	}
}

func LoadLinkUnit(
	id LinkUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
	label *string,
	url globalcommonmodel.Url,
) LinkUnit {
	return LinkUnit{
		UnitEntity: unitmodel.LoadUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			dimension,
			label,
			nil,
			worldcommonmodel.NewLinkUnitType(),
		),
		url: url,
	}
}

func (unit *LinkUnit) GetId() LinkUnitId {
	return NewLinkUnitId(unit.UnitEntity.GetId().Uuid())
}

func (unit *LinkUnit) GetUrl() globalcommonmodel.Url {
	return unit.url
}

func (unit *LinkUnit) Delete() {
}
