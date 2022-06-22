package areadto

import "github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"

type AreaDTO struct {
	From coordinatedto.CoordinateDTO `json:"from"`
	To   coordinatedto.CoordinateDTO `json:"to"`
}
