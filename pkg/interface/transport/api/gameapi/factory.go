package gameapi

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gameapiservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/messaging/redisinteventpublisher"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres"
)

func newGameAppService(presenter gameapiservice.Presenter) (gameAppService gameapiservice.Service, err error) {
	intEventPublisher := redisinteventpublisher.New()
	itemRepository, err := postgres.NewItemRepository()
	if err != nil {
		return gameAppService, err
	}
	playerRepository := memrepo.NewPlayerMemRepository()
	worldRepository, err := postgres.NewWorldRepository()
	if err != nil {
		return gameAppService, err
	}
	unitRepository, err := postgres.NewUnitRepository()
	if err != nil {
		return gameAppService, err
	}
	return gameapiservice.NewService(
		presenter, intEventPublisher, worldRepository, playerRepository, unitRepository, itemRepository,
	), nil
}
