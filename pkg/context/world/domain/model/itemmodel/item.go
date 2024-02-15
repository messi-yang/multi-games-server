package itemmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type Item struct {
	id                 worldcommonmodel.ItemId
	compatibleUnitType worldcommonmodel.UnitType
	name               string
	traversable        bool
	thumbnailSrc       string
	modelSources       []string
}

// Interface Implementation Check
var _ domain.Aggregate[worldcommonmodel.ItemId] = (*Item)(nil)

func LoadItem(
	id worldcommonmodel.ItemId,
	compatibleUnitType worldcommonmodel.UnitType,
	name string,
	traversable bool,
	thumbnailSrc string,
	modelSources []string,
) Item {
	return Item{
		id:                 id,
		compatibleUnitType: compatibleUnitType,
		name:               name,
		traversable:        traversable,
		thumbnailSrc:       thumbnailSrc,
		modelSources:       modelSources,
	}
}

func (item *Item) GetId() worldcommonmodel.ItemId {
	return item.id
}

func (item *Item) GetCompatibleUnitType() worldcommonmodel.UnitType {
	return item.compatibleUnitType
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

func (item *Item) GetModelSources() []string {
	return item.modelSources
}
