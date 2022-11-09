package applicationservice

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/rediseventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/port/adapter/applicationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/port/dto"
	"github.com/google/uuid"
)

var (
	ErrEventNotFound = errors.New("event was not found")
)

type GameApplicationService struct {
	gameService *service.GameService
}

type gameApplicationServiceConfiguration func(service *GameApplicationService) error

func NewGameApplicationService(cfgs ...gameApplicationServiceConfiguration) (*GameApplicationService, error) {
	service := &GameApplicationService{}
	for _, cfg := range cfgs {
		err := cfg(service)
		if err != nil {
			return nil, err
		}
	}
	return service, nil
}

func WithGameService() gameApplicationServiceConfiguration {
	gameService, _ := service.NewGameService(
		service.WithGameMemory(),
	)
	return func(service *GameApplicationService) error {
		service.gameService = gameService
		return nil
	}
}

func (grs *GameApplicationService) CreateGame(dimensionDto dto.DimensionDto) (uuid.UUID, error) {
	dimension, err := dimensionDto.ToValueObject()
	if err != nil {
		return uuid.Nil, err
	}
	gameId, err := grs.gameService.CreateGame(dimension)
	if err != nil {
		return uuid.Nil, err
	}

	return gameId, nil
}

func (grs *GameApplicationService) ReviveUnitsInGame(gameId uuid.UUID, coordinateDtos []dto.CoordinateDto) error {
	coordinates, err := dto.ParseCoordinateDtos(coordinateDtos)
	if err != nil {
		return err
	}

	updatedGame, err := grs.gameService.ReviveUnitsInGame(gameId, coordinates)
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
			applicationevent.NewZoomedAreaUpdatedApplicationEventTopic(updatedGame.GetId(), dto.NewPlayerIdDto(playerId.GetId())),
			applicationevent.NewZoomedAreaUpdatedApplicationEvent(dto.NewAreaDto(area), dto.NewUnitBlockDto(unitBlock)),
		)
	}
	return nil
}

func (grs *GameApplicationService) AddPlayerToGame(gameId uuid.UUID, playerIdDto dto.PlayerIdDto) error {
	playerId := playerIdDto.ToValueObject()
	updatedGame, err := grs.gameService.AddPlayerToGame(gameId, playerId)
	if err != nil {
		return err
	}
	dimensionDto := dto.NewDimensionDto(updatedGame.GetDimension())

	rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.GameInfoUpdatedApplicationEvent](),
	).Publish(
		applicationevent.NewGameInfoUpdatedApplicationEventTopic(gameId, playerIdDto),
		applicationevent.NewGameInfoUpdatedApplicationEvent(dimensionDto),
	)

	return nil
}

func (grs *GameApplicationService) RemovePlayerFromGame(gameId uuid.UUID, playerIdDto dto.PlayerIdDto) error {
	playerId := playerIdDto.ToValueObject()
	_, err := grs.gameService.RemovePlayerFromGame(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}

func (grs *GameApplicationService) AddZoomedAreaToGame(gameId uuid.UUID, playerIdDto dto.PlayerIdDto, areaDto dto.AreaDto) error {
	area, err := areaDto.ToValueObject()
	if err != nil {
		return err
	}
	playerId := playerIdDto.ToValueObject()

	updatedGame, err := grs.gameService.AddZoomedAreaToGame(gameId, playerId, area)
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
		applicationevent.NewAreaZoomedApplicationEventTopic(gameId, playerIdDto),
		applicationevent.NewAreaZoomedApplicationEvent(areaDto, unitBlockDto),
	)

	return nil
}

func (grs *GameApplicationService) RemoveZoomedAreaFromGame(gameId uuid.UUID, playerIdDto dto.PlayerIdDto) error {
	playerId := playerIdDto.ToValueObject()
	_, err := grs.gameService.RemoveZoomedAreaFromGame(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}
