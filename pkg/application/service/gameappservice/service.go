package gameappservice

import (
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
	GetPlayer(GetPlayerQuery) (playermodel.PlayerAgg, error)
	EnterWorld(EnterWorldCommand) error
	Move(MoveCommand) error
	LeaveWorld(LeaveWorldCommand) error
	ChangeHeldItem(ChangeHeldItemCommand) error
	PlaceItem(PlaceItemCommand) error
	RemoveItem(RemoveItemCommand) error
}

type serve struct {
	worldRepository  worldmodel.Repository
	playerRepository playermodel.Repository
	unitRepository   unitmodel.Repository
	itemRepository   itemmodel.Repository
	gameService      service.GameService
}

func NewService(worldRepository worldmodel.Repository, playerRepository playermodel.Repository, unitRepository unitmodel.Repository, itemRepository itemmodel.Repository, gameService service.GameService) Service {
	return &serve{
		worldRepository:  worldRepository,
		playerRepository: playerRepository,
		unitRepository:   unitRepository,
		itemRepository:   itemRepository,
		gameService:      gameService,
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

func (serve *serve) GetPlayer(query GetPlayerQuery) (playermodel.PlayerAgg, error) {
	return serve.playerRepository.Get(query.PlayerId)
}

func (serve *serve) EnterWorld(command EnterWorldCommand) error {
	return serve.gameService.EnterWorld(command.WorldId, command.PlayerId)
}

func (serve *serve) Move(command MoveCommand) error {
	return serve.gameService.Move(command.WorldId, command.PlayerId, command.Direction)
}

func (serve *serve) LeaveWorld(command LeaveWorldCommand) error {
	return serve.gameService.LeaveWorld(command.WorldId, command.PlayerId)
}

func (serve *serve) PlaceItem(command PlaceItemCommand) error {
	return serve.gameService.PlaceItem(command.WorldId, command.PlayerId)
}

func (serve *serve) ChangeHeldItem(command ChangeHeldItemCommand) error {
	return serve.gameService.ChangeHeldItem(command.WorldId, command.PlayerId, command.ItemId)
}

func (serve *serve) RemoveItem(command RemoveItemCommand) error {
	return serve.gameService.RemoveItem(command.WorldId, command.PlayerId)
}
