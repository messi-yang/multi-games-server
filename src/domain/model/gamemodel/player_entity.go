package gamemodel

import (
	"math"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type PlayerEntity struct {
	id            PlayerIdVo
	name          string
	lastGotViewAt commonmodel.LocationVo
	location      commonmodel.LocationVo
}

func NewPlayerEntity(id PlayerIdVo, name string, location commonmodel.LocationVo) PlayerEntity {
	return PlayerEntity{
		id:            id,
		name:          name,
		lastGotViewAt: location,
		location:      location,
	}
}

func (p *PlayerEntity) GetId() PlayerIdVo {
	return p.id
}

func (p *PlayerEntity) GetName() string {
	return p.name
}

func (p *PlayerEntity) GetLastGotViewAt() commonmodel.LocationVo {
	return p.lastGotViewAt
}

func (p *PlayerEntity) SetLastGotViewAt(location commonmodel.LocationVo) {
	p.lastGotViewAt = location
}

func (p *PlayerEntity) GetLocation() commonmodel.LocationVo {
	return p.location
}

func (p *PlayerEntity) SetLocation(location commonmodel.LocationVo) {
	p.location = location
}

func (p *PlayerEntity) IsNewViewNeeded() bool {
	xOffset := int(math.Abs(float64(p.location.GetX() - p.lastGotViewAt.GetX())))
	yOffset := int(math.Abs(float64(p.location.GetY() - p.lastGotViewAt.GetY())))

	return xOffset > 10 || yOffset > 10
}
