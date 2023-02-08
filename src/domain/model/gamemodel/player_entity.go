package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type PlayerEntity struct {
	id       PlayerIdVo
	name     string
	location commonmodel.LocationVo
}

func NewPlayerEntity(id PlayerIdVo, name string, location commonmodel.LocationVo) PlayerEntity {
	return PlayerEntity{
		id:       id,
		name:     name,
		location: location,
	}
}

func (p *PlayerEntity) GetId() PlayerIdVo {
	return p.id
}

func (p *PlayerEntity) GetName() string {
	return p.name
}

func (p *PlayerEntity) GetLocation() commonmodel.LocationVo {
	return p.location
}

func (p *PlayerEntity) SetLocation(location commonmodel.LocationVo) {
	p.location = location
}
