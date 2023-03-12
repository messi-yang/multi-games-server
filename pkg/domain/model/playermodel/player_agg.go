package playermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type PlayerAgg struct {
	id        PlayerIdVo
	worldId   worldmodel.WorldIdVo
	name      string
	position  commonmodel.PositionVo
	direction commonmodel.DirectionVo
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

func (p *PlayerAgg) GetId() PlayerIdVo {
	return p.id
}

func (p *PlayerAgg) GetWorldId() worldmodel.WorldIdVo {
	return p.worldId
}

func (p *PlayerAgg) GetName() string {
	return p.name
}

func (p *PlayerAgg) GetPosition() commonmodel.PositionVo {
	return p.position
}

func (p *PlayerAgg) SetPosition(position commonmodel.PositionVo) {
	p.position = position
}

func (p *PlayerAgg) GetDirection() commonmodel.DirectionVo {
	return p.direction
}

func (p *PlayerAgg) SetDirection(direction commonmodel.DirectionVo) {
	p.direction = direction
}

func (p *PlayerAgg) GetVisionBound() commonmodel.BoundVo {
	playerPosition := p.GetPosition()

	fromX := playerPosition.GetX() - 25
	toX := playerPosition.GetX() + 25

	fromY := playerPosition.GetZ() - 25
	toY := playerPosition.GetZ() + 25

	from := commonmodel.NewPositionVo(fromX, fromY)
	to := commonmodel.NewPositionVo(toX, toY)
	bound, _ := commonmodel.NewBoundVo(from, to)

	return bound
}

func (p *PlayerAgg) CanSeeAnyPositions(positions []commonmodel.PositionVo) bool {
	bound := p.GetVisionBound()
	return bound.CoverAnyPositions(positions)
}
