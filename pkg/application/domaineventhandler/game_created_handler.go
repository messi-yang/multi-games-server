package domaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
)

type GameCreatedHandler struct{}

func NewGameCreatedHandler() memdomaineventhandler.Handler {
	return &GameCreatedHandler{}
}

func ProvideGameCreatedHandler() memdomaineventhandler.Handler {
	return NewGameCreatedHandler()
}

// Set all other games' selected flags to false when a new game is created
func (handler GameCreatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	gameCreated := domainEvent.(gamemodel.GameCreated)
	gameId := gameCreated.GetGameId()
	roomId := gameCreated.GetRoomId()

	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	gameRepo := pgrepo.NewGameRepo(uow, domainEventDispatcher)

	selectedGamesOfRoom, err := gameRepo.GetSelectedGamesByRoomId(roomId)
	if err != nil {
		return err
	}

	for _, game := range selectedGamesOfRoom {
		if game.GetId() == gameId {
			continue
		}

		game.SetSelected(false)
		err = gameRepo.Update(game)
		if err != nil {
			return err
		}
	}

	return nil
}
