package worldmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"

type WorldRepo interface {
	Add(World) error
	Update(World) error
	Get(sharedkernelmodel.WorldId) (World, error)
	Query(limit int, offset int) ([]World, error)
}
