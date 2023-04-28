package itemmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
)

type Item struct {
	id           commonmodel.ItemId
	name         string
	traversable  bool
	thumbnailSrc string
	modelSrc     string
	domainEvents []domainmodel.DomainEvent
}

// Interface Implementation Check
var _ domainmodel.Aggregate = (*Item)(nil)

func NewItem(id commonmodel.ItemId, name string, traversable bool, thumbnailSrc string, modelSrc string) Item {
	return Item{
		id:           id,
		name:         name,
		traversable:  traversable,
		thumbnailSrc: thumbnailSrc,
		modelSrc:     modelSrc,
		domainEvents: []domainmodel.DomainEvent{},
	}
}

func (item *Item) AddDomainEvent(domainEvent domainmodel.DomainEvent) {
	item.domainEvents = append(item.domainEvents, domainEvent)
}

func (item *Item) GetDomainEvents() []domainmodel.DomainEvent {
	return item.domainEvents
}

func (item *Item) GetId() commonmodel.ItemId {
	return item.id
}

func (item *Item) GetName() string {
	return item.name
}

func (item *Item) GetTraversable() bool {
	return item.traversable
}

func (item *Item) GetThumbnailSrc() string {
	return item.thumbnailSrc
}

func (item *Item) GetModelSrc() string {
	return item.modelSrc
}
