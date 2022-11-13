package applicationservice

import (
	commonValueObject "github.com/dum-dum-genius/game-of-liberty-computer/common/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/rediseventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/service"
	liveGameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/adapter/applicationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/dto"
	"github.com/google/uuid"
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
		service.WithGameMemoryRepository(),
	)
	return func(service *LiveGameApplicationService) error {
		service.liveGameService = liveGameService
		return nil
	}
}

func (grs *LiveGameApplicationService) CreateLiveGame(dimensionDto dto.DimensionDto) (liveGameValueObject.LiveGameId, error) {
	dimension, err := dimensionDto.ToValueObject()
	if err != nil {
		return liveGameValueObject.LiveGameId{}, err
	}
	liveGameId, err := grs.liveGameService.CreateLiveGame(dimension)
	if err != nil {
		return liveGameValueObject.LiveGameId{}, err
	}

	return liveGameId, nil
}

func (grs *LiveGameApplicationService) ReviveUnitsInLiveGame(liveGameId liveGameValueObject.LiveGameId, coordinateDtos []dto.CoordinateDto) error {
	coordinates, err := dto.ParseCoordinateDtos(coordinateDtos)
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
			applicationevent.NewZoomedAreaUpdatedApplicationEvent(dto.NewAreaDto(area), dto.NewUnitBlockDto(unitBlock)),
		)
	}
	return nil
}

func (grs *LiveGameApplicationService) AddPlayerToLiveGame(liveGameId liveGameValueObject.LiveGameId, playerId uuid.UUID) error {
	updatedGame, err := grs.liveGameService.AddPlayerToLiveGame(liveGameId, commonValueObject.NewPlayerId(playerId))
	if err != nil {
		return err
	}
	dimensionDto := dto.NewDimensionDto(updatedGame.GetDimension())

	rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.GameInfoUpdatedApplicationEvent](),
	).Publish(
		applicationevent.NewGameInfoUpdatedApplicationEventTopic(liveGameId.GetId(), playerId),
		applicationevent.NewGameInfoUpdatedApplicationEvent(dimensionDto),
	)

	return nil
}

func (grs *LiveGameApplicationService) RemovePlayerFromLiveGame(liveGameId liveGameValueObject.LiveGameId, playerId uuid.UUID) error {
	_, err := grs.liveGameService.RemovePlayerFromLiveGame(liveGameId, commonValueObject.NewPlayerId(playerId))
	if err != nil {
		return err
	}

	return nil
}

func (grs *LiveGameApplicationService) AddZoomedAreaToLiveGame(liveGameId liveGameValueObject.LiveGameId, playerId uuid.UUID, areaDto dto.AreaDto) error {
	area, err := areaDto.ToValueObject()
	if err != nil {
		return err
	}

	updatedGame, err := grs.liveGameService.AddZoomedAreaToLiveGame(liveGameId, commonValueObject.NewPlayerId(playerId), area)
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
		applicationevent.NewAreaZoomedApplicationEventTopic(liveGameId.GetId(), playerId),
		applicationevent.NewAreaZoomedApplicationEvent(areaDto, unitBlockDto),
	)

	return nil
}

func (grs *LiveGameApplicationService) RemoveZoomedAreaFromLiveGame(liveGameId liveGameValueObject.LiveGameId, playerId uuid.UUID) error {
	_, err := grs.liveGameService.RemoveZoomedAreaFromLiveGame(liveGameId, commonValueObject.NewPlayerId(playerId))
	if err != nil {
		return err
	}

	return nil
}
