package embedunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type EmbedUnit struct {
	unitmodel.UnitEntity
	embedCode worldcommonmodel.EmbedCode
}

// Interface Implementation Check
var _ domain.Aggregate[EmbedUnitId] = (*EmbedUnit)(nil)

func NewEmbedUnit(
	id EmbedUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	label *string,
	embedCode worldcommonmodel.EmbedCode,
) EmbedUnit {
	return EmbedUnit{
		UnitEntity: unitmodel.NewUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			label,
			worldcommonmodel.NewEmbedUnitType(),
			nil,
		),
		embedCode: embedCode,
	}
}

func LoadEmbedUnit(
	id EmbedUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	label *string,
	embedCode worldcommonmodel.EmbedCode,
) EmbedUnit {
	return EmbedUnit{
		UnitEntity: unitmodel.LoadUnitEntity(
			unitmodel.NewUnitId(id.Uuid()),
			worldId,
			position,
			itemId,
			direction,
			label,
			worldcommonmodel.NewEmbedUnitType(),
			nil,
		),
		embedCode: embedCode,
	}
}

func (unit *EmbedUnit) GetId() EmbedUnitId {
	return NewEmbedUnitId(unit.UnitEntity.GetId().Uuid())
}

func (unit *EmbedUnit) GetEmbedCode() worldcommonmodel.EmbedCode {
	return unit.embedCode
}

func (unit *EmbedUnit) Delete() {
}
