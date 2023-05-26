package worldmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"

type Repo interface {
	Add(World) error
	Update(World) error
	Get(commonmodel.WorldId) (World, error)
	Query(limit int, offset int) ([]World, error)
}
