package itemmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
)

type ItemAgg struct {
	id           commonmodel.ItemIdVo
	name         string
	traversable  bool
	thumbnailSrc string
	modelSrc     string
	domainEvents []domainmodel.DomainEvent
}

// Interface Implementation Check
var _ domainmodel.Aggregate = (*ItemAgg)(nil)

func NewItemAgg(id commonmodel.ItemIdVo, name string, traversable bool, thumbnailSrc string, modelSrc string) ItemAgg {
	return ItemAgg{
		id:           id,
		name:         name,
		traversable:  traversable,
		thumbnailSrc: thumbnailSrc,
		modelSrc:     modelSrc,
		domainEvents: []domainmodel.DomainEvent{},
	}
}

func (agg *ItemAgg) AddDomainEvent(domainEvent domainmodel.DomainEvent) {
	agg.domainEvents = append(agg.domainEvents, domainEvent)
}

func (agg *ItemAgg) GetDomainEvents() []domainmodel.DomainEvent {
	return agg.domainEvents
}

func (agg *ItemAgg) GetId() commonmodel.ItemIdVo {
	return agg.id
}

func (agg *ItemAgg) GetName() string {
	return agg.name
}

func (agg *ItemAgg) GetTraversable() bool {
	return agg.traversable
}

func (agg *ItemAgg) GetThumbnailSrc() string {
	return agg.thumbnailSrc
}

func (agg *ItemAgg) GetModelSrc() string {
	return agg.modelSrc
}
