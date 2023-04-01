package playermodel

import (
	"math"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

func calculatePlayerVisionBound(pos commonmodel.PositionVo) commonmodel.BoundVo {
	fromX := pos.GetX() - 30
	toX := pos.GetX() + 30

	fromY := pos.GetZ() - 30
	toY := pos.GetZ() + 30

	from := commonmodel.NewPositionVo(fromX, fromY)
	to := commonmodel.NewPositionVo(toX, toY)
	bound, _ := commonmodel.NewBoundVo(from, to)

	return bound
}

type PlayerAgg struct {
	id          PlayerIdVo              // Id of the player
	worldId     worldmodel.WorldIdVo    // The id of the world the player belongs to
	name        string                  // The name of the player
	position    commonmodel.PositionVo  // The current position of the player
	direction   commonmodel.DirectionVo // The direction where the player is facing
	visionBound commonmodel.BoundVo     // The vision bound of the player
	heldItemId  *itemmodel.ItemIdVo     // Optional, The item held by the player
}

func NewPlayerAgg(id PlayerIdVo, worldId worldmodel.WorldIdVo, name string, position commonmodel.PositionVo, direction commonmodel.DirectionVo, heldItemId *itemmodel.ItemIdVo) PlayerAgg {
	player := PlayerAgg{
		id:          id,
		worldId:     worldId,
		name:        name,
		position:    position,
		direction:   direction,
		visionBound: calculatePlayerVisionBound(position),
		heldItemId:  heldItemId,
	}
	return player
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

func (agg *PlayerAgg) GetDirection() commonmodel.DirectionVo {
	return agg.direction
}

func (agg *PlayerAgg) ChangeDirection(direction commonmodel.DirectionVo) {
	agg.direction = direction
}

func (agg *PlayerAgg) ShallUpdateVisionBound() bool {
	visionBoundCenterPos := agg.visionBound.GetCenterPos()
	xDistance := int(math.Abs(float64(agg.position.GetX() - visionBoundCenterPos.GetX())))
	zDistance := int(math.Abs(float64(agg.position.GetZ() - visionBoundCenterPos.GetZ())))
	return xDistance >= 10 || zDistance >= 10
}

func (agg *PlayerAgg) UpdateVisionBound() {
	agg.visionBound = calculatePlayerVisionBound(agg.GetPosition())
}

func (agg *PlayerAgg) GetVisionBound() commonmodel.BoundVo {
	return agg.visionBound
}

func (agg *PlayerAgg) CanSeeAnyPositions(positions []commonmodel.PositionVo) bool {
	bound := agg.GetVisionBound()
	return bound.CoverAnyPositions(positions)
}

func (agg *PlayerAgg) GetPositionOneStepFoward() commonmodel.PositionVo {
	direction := agg.direction
	position := agg.position

	if direction.IsUp() {
		return position.Shift(0, -1)
	} else if direction.IsRight() {
		return position.Shift(1, 0)
	} else if direction.IsDown() {
		return position.Shift(0, 1)
	} else if direction.IsLeft() {
		return position.Shift(-1, 0)
	} else {
		return position.Shift(0, 1)
	}
}

func (agg *PlayerAgg) HasHeldItem() bool {
	return agg.heldItemId != nil
}

func (agg *PlayerAgg) GetHeldItemId() *itemmodel.ItemIdVo {
	return agg.heldItemId
}
