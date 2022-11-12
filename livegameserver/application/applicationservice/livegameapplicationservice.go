package applicationservice

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/rediseventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/adapter/applicationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/dto"
	"github.com/google/uuid"
)

var (
	ErrEventNotFound = errors.New("event was not found")
)

type LiveGameApplicationService struct {
	liveGameService *service.LiveGameService
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
	liveGameService, _ := service.NewLiveGameService(
		service.WithGameMemory(),
	)
	return func(service *LiveGameApplicationService) error {
		service.liveGameService = liveGameService
		return nil
	}
}

func (grs *LiveGameApplicationService) CreateLiveGame(dimensionDto dto.DimensionDto) (valueobject.GameId, error) {
	dimension, err := dimensionDto.ToValueObject()
	if err != nil {
		return valueobject.GameId{}, err
	}
	gameId, err := grs.liveGameService.CreateLiveGame(dimension)
	if err != nil {
		return valueobject.GameId{}, err
	}

	return gameId, nil
}

func (grs *LiveGameApplicationService) ReviveUnitsInLiveGame(gameId valueobject.GameId, coordinateDtos []dto.CoordinateDto) error {
	coordinates, err := dto.ParseCoordinateDtos(coordinateDtos)
	if err != nil {
		return err
	}

	updatedGame, err := grs.liveGameService.ReviveUnitsInLiveGame(gameId, coordinates)
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
			applicationevent.NewZoomedAreaUpdatedApplicationEvent(dto.NewAreaDto(area), dto.NewUnitBlockDto(unitBlock)),
		)
	}
	return nil
}

func (grs *LiveGameApplicationService) AddPlayerToLiveGame(gameId valueobject.GameId, playerId uuid.UUID) error {
	updatedGame, err := grs.liveGameService.AddPlayerToLiveGame(gameId, valueobject.NewPlayerId(playerId))
	if err != nil {
		return err
	}
	dimensionDto := dto.NewDimensionDto(updatedGame.GetDimension())

	rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.GameInfoUpdatedApplicationEvent](),
	).Publish(
		applicationevent.NewGameInfoUpdatedApplicationEventTopic(gameId.GetId(), playerId),
		applicationevent.NewGameInfoUpdatedApplicationEvent(dimensionDto),
	)

	return nil
}

func (grs *LiveGameApplicationService) RemovePlayerFromLiveGame(gameId valueobject.GameId, playerId uuid.UUID) error {
	_, err := grs.liveGameService.RemovePlayerFromLiveGame(gameId, valueobject.NewPlayerId(playerId))
	if err != nil {
		return err
	}

	return nil
}

func (grs *LiveGameApplicationService) AddZoomedAreaToLiveGame(gameId valueobject.GameId, playerId uuid.UUID, areaDto dto.AreaDto) error {
	area, err := areaDto.ToValueObject()
	if err != nil {
		return err
	}

	updatedGame, err := grs.liveGameService.AddZoomedAreaToLiveGame(gameId, valueobject.NewPlayerId(playerId), area)
	if err != nil {
		return err
	}

	unitBlock, err := updatedGame.GetUnitBlockByArea(area)
	if err != nil {
		return err
	}

	unitBlockDto := dto.NewUnitBlockDto(unitBlock)

	rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.AreaZoomedApplicationEvent](),
	).Publish(
		applicationevent.NewAreaZoomedApplicationEventTopic(gameId.GetId(), playerId),
		applicationevent.NewAreaZoomedApplicationEvent(areaDto, unitBlockDto),
	)

	return nil
}

func (grs *LiveGameApplicationService) RemoveZoomedAreaFromLiveGame(gameId valueobject.GameId, playerId uuid.UUID) error {
	_, err := grs.liveGameService.RemoveZoomedAreaFromLiveGame(gameId, valueobject.NewPlayerId(playerId))
	if err != nil {
		return err
	}

	return nil
}
