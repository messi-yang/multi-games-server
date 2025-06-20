package commandmodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"

type CommandRepo interface {
	Add(Command) error
	GetCommandsOfGame(gameId gamemodel.GameId) ([]Command, error)
}
