package worldmodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"

type WorldRepo interface {
	Add(World) error
	Update(World) error
	Delete(World) error
	Get(globalcommonmodel.WorldId) (World, error)
	GetWorldsOfUser(globalcommonmodel.UserId) ([]World, error)
	Query(limit int, offset int) ([]World, error)
}
