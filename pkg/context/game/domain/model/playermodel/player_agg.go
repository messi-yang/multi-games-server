package playermodel

import (
	"math"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
)

func calculatePlayerVisionBound(pos commonmodel.PositionVo) commonmodel.BoundVo {
	fromX := pos.GetX() - 50
	toX := pos.GetX() + 50

	fromY := pos.GetZ() - 50
	toY := pos.GetZ() + 50

	from := commonmodel.NewPositionVo(fromX, fromY)
	to := commonmodel.NewPositionVo(toX, toY)
	bound, _ := commonmodel.NewBoundVo(from, to)

	return bound
}

type PlayerAgg struct {
	id          commonmodel.PlayerIdVo  // Id of the player
	worldId     commonmodel.WorldIdVo   // The id of the world the player belongs to
	name        string                  // The name of the player
	position    commonmodel.PositionVo  // The current position of the player
	direction   commonmodel.DirectionVo // The direction where the player is facing
	visionBound commonmodel.BoundVo     // The vision bound of the player
	heldItemId  *commonmodel.ItemIdVo   // Optional, The item held by the player
}

func NewPlayerAgg(id commonmodel.PlayerIdVo, worldId commonmodel.WorldIdVo, name string, position commonmodel.PositionVo, direction commonmodel.DirectionVo, heldItemId *commonmodel.ItemIdVo) PlayerAgg {
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

func (agg *PlayerAgg) GetId() commonmodel.PlayerIdVo {
	return agg.id
}

func (agg *PlayerAgg) GetWorldId() commonmodel.WorldIdVo {
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

func (agg *PlayerAgg) ChangeHeldItem(itemId commonmodel.ItemIdVo) {
	agg.heldItemId = &itemId
}

func (agg *PlayerAgg) GetHeldItemId() *commonmodel.ItemIdVo {
	return agg.heldItemId
}
