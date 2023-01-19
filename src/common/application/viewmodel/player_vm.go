package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"

type PlayerVm struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewPlayerVm(playerAgg playermodel.PlayerAgg) PlayerVm {
	return PlayerVm{
		Id:   playerAgg.GetId().ToString(),
		Name: playerAgg.GetName(),
	}
}
