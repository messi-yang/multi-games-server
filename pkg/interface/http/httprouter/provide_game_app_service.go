package httprouter

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gameappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gamerappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/itemappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/worldappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/service/gamedomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
)

func provideGamerAppService() (gamerappsrv.Service, error) {
	gamerRepo, err := pgrepo.NewGamerRepo()
	if err != nil {
		return nil, err
	}
	return gamerappsrv.NewService(gamerRepo), nil
}

func provideWorldAppService() (worldappsrv.Service, error) {
	worldRepo, err := pgrepo.NewWorldRepo()
	if err != nil {
		return nil, err
	}
	itemRepo, err := pgrepo.NewItemRepo()
	if err != nil {
		return nil, err
	}
	unitRepo, err := pgrepo.NewUnitRepo()
	if err != nil {
		return nil, err
	}
	return worldappsrv.NewService(worldRepo, unitRepo, itemRepo), nil
}

func provideItemAppService() (itemappsrv.Service, error) {
	itemRepo, err := pgrepo.NewItemRepo()
	if err != nil {
		return nil, err
	}
	return itemappsrv.NewService(itemRepo), nil
}

func provideGameAppService() (gameappsrv.Service, error) {
	itemRepo, err := pgrepo.NewItemRepo()
	if err != nil {
		return nil, err
	}
	playerRepo := memrepo.NewPlayerMemRepo()
	worldRepo, err := pgrepo.NewWorldRepo()
	if err != nil {
		return nil, err
	}
	unitRepo, err := pgrepo.NewUnitRepo()
	if err != nil {
		return nil, err
	}
	gameDomainService := gamedomainsrv.NewService(worldRepo, playerRepo, unitRepo, itemRepo)
	return gameappsrv.NewService(
		worldRepo, playerRepo, unitRepo, itemRepo, gameDomainService,
	), nil
}
