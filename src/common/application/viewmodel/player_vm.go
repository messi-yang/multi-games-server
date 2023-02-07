package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"

type PlayerVm struct {
	Id       string     `json:"id"`
	Name     string     `json:"name"`
	Location LocationVm `json:"location"`
}

func NewPlayerVm(player gamemodel.PlayerEntity) PlayerVm {
	return PlayerVm{
		Id:       player.GetId().ToString(),
		Name:     player.GetName(),
		Location: NewLocationVm(player.GetLocation()),
	}
}
