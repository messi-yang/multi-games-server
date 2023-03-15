package playermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type PlayerAgg struct {
	id        PlayerIdVo
	worldId   worldmodel.WorldIdVo    // The id of the world the player belongs to
	name      string                  // The name of the player
	position  commonmodel.PositionVo  // The current position of the player
	direction commonmodel.DirectionVo // The direction where the player is facing
}

func NewPlayerAgg(id PlayerIdVo, worldId worldmodel.WorldIdVo, name string, position commonmodel.PositionVo, direction commonmodel.DirectionVo) PlayerAgg {
	return PlayerAgg{
		id:        id,
		worldId:   worldId,
		name:      name,
		position:  position,
		direction: direction,
	}
}

func (agg *PlayerAgg) GetId() PlayerIdVo {
	return agg.id
}

func (agg *PlayerAgg) GetWorldId() worldmodel.WorldIdVo {
	return agg.worldId
}

func (agg *PlayerAgg) GetName() string {
	return agg.name
}

func (agg *PlayerAgg) GetPosition() commonmodel.PositionVo {
	return agg.position
}

func (agg *PlayerAgg) ChangePosition(position commonmodel.PositionVo) {
	agg.position = position
}

// 	xDistance := int(math.Abs(float64(agg.position.GetX() - agg.lastGotUnits.GetX())))
// 	zDistance := int(math.Abs(float64(agg.position.GetZ() - agg.lastGotUnits.GetZ())))
// 	if xDistance >= MARK_POSITION_DISTANCE || zDistance >= MARK_POSITION_DISTANCE {
// 		agg.lastGotUnits = agg.position
// 	}

func (agg *PlayerAgg) GetDirection() commonmodel.DirectionVo {
	return agg.direction
}

func (agg *PlayerAgg) ChangeDirection(direction commonmodel.DirectionVo) {
	agg.direction = direction
}

func (agg *PlayerAgg) GetVisionBound() commonmodel.BoundVo {
	playerPosition := agg.GetPosition()

	fromX := playerPosition.GetX() - 25
	toX := playerPosition.GetX() + 25

	fromY := playerPosition.GetZ() - 25
	toY := playerPosition.GetZ() + 25

	from := commonmodel.NewPositionVo(fromX, fromY)
	to := commonmodel.NewPositionVo(toX, toY)
	bound, _ := commonmodel.NewBoundVo(from, to)

	return bound
}

func (agg *PlayerAgg) CanSeeAnyPositions(positions []commonmodel.PositionVo) bool {
	bound := agg.GetVisionBound()
	return bound.CoverAnyPositions(positions)
}
