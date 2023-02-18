package playermodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
)

type PlayerAgg struct {
	id       PlayerIdVo
	gameId   gamemodel.GameIdVo
	name     string
	location commonmodel.LocationVo
}

func NewPlayerAgg(id PlayerIdVo, gameId gamemodel.GameIdVo, name string, location commonmodel.LocationVo) PlayerAgg {
	return PlayerAgg{
		id:       id,
		gameId:   gameId,
		name:     name,
		location: location,
	}
}

func (p *PlayerAgg) GetId() PlayerIdVo {
	return p.id
}

func (p *PlayerAgg) GetGameId() gamemodel.GameIdVo {
	return p.gameId
}

func (p *PlayerAgg) GetName() string {
	return p.name
}

func (p *PlayerAgg) GetLocation() commonmodel.LocationVo {
	return p.location
}

func (p *PlayerAgg) SetLocation(location commonmodel.LocationVo) {
	p.location = location
}

func (p *PlayerAgg) GetVisionBound() commonmodel.BoundVo {
	playerLocation := p.GetLocation()

	fromX := playerLocation.GetX() - 35
	toX := playerLocation.GetX() + 35

	fromY := playerLocation.GetZ() - 35
	toY := playerLocation.GetZ() + 35

	from := commonmodel.NewLocationVo(fromX, fromY)
	to := commonmodel.NewLocationVo(toX, toY)
	bound, _ := commonmodel.NewBoundVo(from, to)

	return bound
}

func (p *PlayerAgg) CanSeeAnyLocations(locations []commonmodel.LocationVo) bool {
	bound := p.GetVisionBound()
	return bound.CoverAnyLocations(locations)
}
