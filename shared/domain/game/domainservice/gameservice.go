package domainservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

type GameDomainService interface {
	CreateGame(mapSize valueobject.MapSize) entity.Game
	GetGame(gameId uuid.UUID) (entity.Game, error)
	GetAllGames() []entity.Game
}

type gameDomainServiceImplement struct {
}

type GameDomainServiceConfiguration struct {
}

func NewGameDomainService(configuration GameDomainServiceConfiguration) GameDomainService {
	return &gameDomainServiceImplement{}
}

func (service *gameDomainServiceImplement) CreateGame(mapSize valueobject.MapSize) entity.Game {
	return entity.NewGame(mapSize, time.Second.Microseconds())
}

func (service *gameDomainServiceImplement) GetGame(gameId uuid.UUID) (entity.Game, error) {
	return entity.Game{}, nil
}

func (service *gameDomainServiceImplement) GetAllGames() []entity.Game {
	return nil
}
