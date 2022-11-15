package applicationservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/rediseventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/livegame/port/adapter/applicationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/service/gameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/service/livegameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type LiveGameApplicationService struct {
	liveGameService *livegameservice.LiveGameService
	gameService     *gameservice.GameService
}

type liveGameApplicationServiceConfiguration func(service *LiveGameApplicationService) error

func NewLiveGameApplicationService(cfgs ...liveGameApplicationServiceConfiguration) (*LiveGameApplicationService, error) {
	service := &LiveGameApplicationService{}
	for _, cfg := range cfgs {
		err := cfg(service)
		if err != nil {
			return nil, err
		}
	}
	return service, nil
}

func WithLiveGameService() liveGameApplicationServiceConfiguration {
	return func(liveGameApplicationService *LiveGameApplicationService) error {
		liveGameService, _ := livegameservice.NewLiveGameService(
			livegameservice.WithGameMemoryRepository(),
		)
		liveGameApplicationService.liveGameService = liveGameService
		return nil
	}
}

func WithGameService() liveGameApplicationServiceConfiguration {
	return func(liveGameApplicationService *LiveGameApplicationService) error {
		gameService, _ := gameservice.NewGameService(
			gameservice.WithPostgresGameRepository(),
		)
		liveGameApplicationService.gameService = gameService
		return nil
	}
}

func (grs *LiveGameApplicationService) CreateLiveGame(gameId gamemodel.GameId) (livegamemodel.LiveGameId, error) {
	game, err := grs.gameService.GetGame(gameId)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	liveGameId, err := grs.liveGameService.CreateLiveGame(game)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}

	return liveGameId, nil
}

func (grs *LiveGameApplicationService) ReviveUnitsInLiveGame(liveGameId livegamemodel.LiveGameId, coordinatePresenterDtos []presenterdto.CoordinatePresenterDto) error {
	coordinates, err := presenterdto.ParseCoordinatePresenterDtos(coordinatePresenterDtos)
	if err != nil {
		return err
	}

	updatedGame, err := grs.liveGameService.ReviveUnitsInLiveGame(liveGameId, coordinates)
	if err != nil {
		return err
	}

	for playerId, area := range updatedGame.GetZoomedAreas() {
		coordinatesInArea := area.FilterCoordinates(coordinates)
		if len(coordinatesInArea) == 0 {
			continue
		}
		unitBlock, err := updatedGame.GetUnitBlockByArea(area)
		if err != nil {
			continue
		}
		rediseventbus.NewRedisApplicationEventBus(
			rediseventbus.WithRedisInfrastructureService[applicationevent.ZoomedAreaUpdatedApplicationEvent](),
		).Publish(
			applicationevent.NewZoomedAreaUpdatedApplicationEventTopic(updatedGame.GetId().GetId(), playerId.GetId()),
			applicationevent.NewZoomedAreaUpdatedApplicationEvent(presenterdto.NewAreaPresenterDto(area), presenterdto.NewUnitBlockPresenterDto(unitBlock)),
		)
	}
	return nil
}

func (grs *LiveGameApplicationService) AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId uuid.UUID) error {
	updatedGame, err := grs.liveGameService.AddPlayerToLiveGame(liveGameId, gamecommonmodel.NewPlayerId(playerId))
	if err != nil {
		return err
	}
	dimensionPresenterDto := presenterdto.NewDimensionPresenterDto(updatedGame.GetDimension())

	rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.GameInfoUpdatedApplicationEvent](),
	).Publish(
		applicationevent.NewGameInfoUpdatedApplicationEventTopic(liveGameId.GetId(), playerId),
		applicationevent.NewGameInfoUpdatedApplicationEvent(dimensionPresenterDto),
	)

	return nil
}

func (grs *LiveGameApplicationService) RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId uuid.UUID) error {
	_, err := grs.liveGameService.RemovePlayerFromLiveGame(liveGameId, gamecommonmodel.NewPlayerId(playerId))
	if err != nil {
		return err
	}

	return nil
}

func (grs *LiveGameApplicationService) AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId uuid.UUID, areaPresenterDto presenterdto.AreaPresenterDto) error {
	area, err := areaPresenterDto.ToValueObject()
	if err != nil {
		return err
	}

	updatedGame, err := grs.liveGameService.AddZoomedAreaToLiveGame(liveGameId, gamecommonmodel.NewPlayerId(playerId), area)
	if err != nil {
		return err
	}

	unitBlock, err := updatedGame.GetUnitBlockByArea(area)
	if err != nil {
		return err
	}

	unitBlockPresenterDto := presenterdto.NewUnitBlockPresenterDto(unitBlock)

	rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.AreaZoomedApplicationEvent](),
	).Publish(
		applicationevent.NewAreaZoomedApplicationEventTopic(liveGameId.GetId(), playerId),
		applicationevent.NewAreaZoomedApplicationEvent(areaPresenterDto, unitBlockPresenterDto),
	)

	return nil
}

func (grs *LiveGameApplicationService) RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId uuid.UUID) error {
	_, err := grs.liveGameService.RemoveZoomedAreaFromLiveGame(liveGameId, gamecommonmodel.NewPlayerId(playerId))
	if err != nil {
		return err
	}

	return nil
}
