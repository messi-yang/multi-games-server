package domainservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/repository"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

type GameDomainService interface {
	CreateGame(mapSize valueobject.MapSize) (entity.Game, error)
	GetGame(gameId uuid.UUID) (entity.Game, error)
	GetFirstGameId() (uuid.UUID, error)
}

type gameDomainServiceImplement struct {
	gameRepository repository.GameRepository
}

type GameDomainServiceConfiguration struct {
	GameRepository repository.GameRepository
}

func NewGameDomainService(configuration GameDomainServiceConfiguration) GameDomainService {
	return &gameDomainServiceImplement{
		gameRepository: configuration.GameRepository,
	}
}

func (service *gameDomainServiceImplement) CreateGame(mapSize valueobject.MapSize) (entity.Game, error) {
	newGame := entity.NewGame(uuid.New(), entity.NewUnitMap(mapSize))
	err := service.gameRepository.Add(newGame)
	if err != nil {
		return entity.Game{}, err
	}

	return newGame, nil
}

func (service *gameDomainServiceImplement) GetGame(gameId uuid.UUID) (entity.Game, error) {
	game, err := service.gameRepository.Get(gameId)
	if err != nil {
		return entity.Game{}, err
	}

	return game, nil
}

func (service *gameDomainServiceImplement) GetFirstGameId() (uuid.UUID, error) {
	gameId, err := service.gameRepository.GetFirstGameId()
	return gameId, err
}
