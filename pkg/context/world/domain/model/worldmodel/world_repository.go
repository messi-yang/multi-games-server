package worldmodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"

type WorldRepo interface {
	Add(World) error
	Update(World) error
	Delete(World) error
	Get(sharedkernelmodel.WorldId) (World, error)
	GetWorldsOfUser(sharedkernelmodel.UserId) ([]World, error)
	Query(limit int, offset int) ([]World, error)
}
