package itemmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type Item struct {
	id                   worldcommonmodel.ItemId
	name                 string
	traversable          bool
	thumbnailSrc         string
	modelSrc             string
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.Aggregate = (*Item)(nil)

func NewItem(id worldcommonmodel.ItemId, name string, traversable bool, thumbnailSrc string, modelSrc string) Item {
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

func (item *Item) GetId() worldcommonmodel.ItemId {
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
