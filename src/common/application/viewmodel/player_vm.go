package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"

type PlayerVm struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewPlayerVm(playerAgr playermodel.Player) PlayerVm {
	return PlayerVm{
		Id:   playerAgr.GetId().ToString(),
		Name: playerAgr.GetName(),
	}
}
