package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"

type PlayerVm struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewPlayerVm(player livegamemodel.PlayerEntity) PlayerVm {
	return PlayerVm{
		Id:   player.GetId().ToString(),
		Name: player.GetName(),
	}
}
