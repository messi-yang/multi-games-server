package gameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/service"
)

type Service interface {
	GetNearbyPlayers(GetNearbyPlayersQuery) ([]playermodel.PlayerAgg, error)
	GetNearbyUnits(GetNearbyUnitsQuery) (commonmodel.BoundVo, []unitmodel.UnitAgg, error)
	AddPlayer(AddPlayerCommand) error
	MovePlayer(MovePlayerCommand) (isVisionBoundUpdated bool, err error)
	RemovePlayer(RemovePlayerCommand) error
	PlaceItem(PlaceItemCommand) error
	DestroyItem(DestroyItemCommand) error
}

type serve struct {
	intEventPublisher intevent.Publisher
	worldRepository   worldmodel.Repository
	playerRepository  playermodel.Repository
	unitRepository    unitmodel.Repository
	itemRepository    itemmodel.Repository
	gameService       service.GameService
}

func NewService(intEventPublisher intevent.Publisher, worldRepository worldmodel.Repository, playerRepository playermodel.Repository, unitRepository unitmodel.Repository, itemRepository itemmodel.Repository) Service {
	return &serve{
		intEventPublisher: intEventPublisher,
		worldRepository:   worldRepository,
		playerRepository:  playerRepository,
		unitRepository:    unitRepository,
		itemRepository:    itemRepository,
		gameService:       service.NewGameService(worldRepository, playerRepository, unitRepository, itemRepository),
	}
}

func (serve *serve) GetNearbyPlayers(query GetNearbyPlayersQuery) (
	players []playermodel.PlayerAgg, err error,
) {
	player, err := serve.playerRepository.Get(query.PlayerId)
	if err != nil {
		return players, err
	}

	return serve.playerRepository.GetPlayersAround(query.WorldId, player.GetPosition())
}

func (serve *serve) GetNearbyUnits(query GetNearbyUnitsQuery) (
	visionBound commonmodel.BoundVo, units []unitmodel.UnitAgg, err error,
) {
	player, err := serve.playerRepository.Get(query.PlayerId)
	if err != nil {
		return visionBound, units, err
	}

	visionBound = player.GetVisionBound()
	units, err = serve.unitRepository.GetUnitsInBound(query.WorldId, visionBound)
	if err != nil {
		return visionBound, units, err
	}

	return visionBound, units, nil
}

func (serve *serve) AddPlayer(command AddPlayerCommand) error {
	return serve.gameService.AddPlayer(command.WorldId, command.PlayerId)
}

func (serve *serve) MovePlayer(command MovePlayerCommand) (isVisionBoundUpdated bool, err error) {
	return serve.gameService.MovePlayer(command.WorldId, command.PlayerId, command.Direction)
}

func (serve *serve) RemovePlayer(command RemovePlayerCommand) error {
	return serve.gameService.RemovePlayer(command.WorldId, command.PlayerId)
}

func (serve *serve) PlaceItem(command PlaceItemCommand) error {
	return serve.gameService.PlaceItem(command.WorldId, command.PlayerId, command.ItemId)
}

func (serve *serve) DestroyItem(command DestroyItemCommand) error {
	return serve.gameService.DestroyItem(command.WorldId, command.PlayerId)
}
