package itemmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/common/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
)

type Item struct {
	id                   commonmodel.ItemId
	name                 string
	traversable          bool
	thumbnailSrc         string
	modelSrc             string
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*Item)(nil)

func NewItem(id commonmodel.ItemId, name string, traversable bool, thumbnailSrc string, modelSrc string) Item {
	return Item{
		id:                   id,
		name:                 name,
		traversable:          traversable,
		thumbnailSrc:         thumbnailSrc,
		modelSrc:             modelSrc,
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (item *Item) PopDomainEvents() []domain.DomainEvent {
	return item.domainEventCollector.PopAll()
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
