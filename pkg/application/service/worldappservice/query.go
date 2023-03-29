package worldappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type GetWorldQuery struct {
	WorldId worldmodel.WorldIdVo
}

type GetWorldsQuery struct{}
