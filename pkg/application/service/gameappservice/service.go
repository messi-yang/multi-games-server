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
	FindNearbyPlayers(FindNearbyPlayersQuery) ([]playermodel.PlayerAgg, error)
	QueryUnits(QueryUnitsQuery) (commonmodel.BoundVo, []unitmodel.UnitAgg, error)
	AddPlayer(AddPlayerCommand) ([]itemmodel.ItemAgg, []playermodel.PlayerAgg, commonmodel.BoundVo, []unitmodel.UnitAgg, error)
	MovePlayer(MovePlayerCommand) error
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

func (serve *serve) publishUnitsUpdatedEventTo(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error {
	return serve.intEventPublisher.Publish(
		NewUnitsUpdatedIntEventChannel(worldId.Uuid(), playerId.Uuid()),
		UnitsUpdatedIntEvent{},
	)
}

func (serve *serve) publishUnitsUpdatedEventToNearPlayers(playerId playermodel.PlayerIdVo) error {
	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	players, err := serve.playerRepository.GetPlayersAround(player.GetWorldId(), player.GetPosition())
	if err != nil {
		return err
	}

	for _, player := range players {
		err = serve.publishUnitsUpdatedEventTo(player.GetWorldId(), player.GetId())
		if err != nil {
			return err
		}
	}

	return nil
}

func (serve *serve) publishPlayersUpdatedEventToNearPlayers(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) error {
	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	players, err := serve.playerRepository.GetPlayersAround(worldId, player.GetPosition())
	if err != nil {
		return err
	}

	for _, player := range players {
		err = serve.intEventPublisher.Publish(
			NewPlayersUpdatedIntEventChannel(worldId.Uuid(), player.GetId().Uuid()),
			PlayersUpdatedIntEvent{},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (serve *serve) FindNearbyPlayers(query FindNearbyPlayersQuery) (
	players []playermodel.PlayerAgg, err error,
) {
	player, err := serve.playerRepository.Get(query.PlayerId)
	if err != nil {
		return players, err
	}

	return serve.playerRepository.GetPlayersAround(query.WorldId, player.GetPosition())
}

func (serve *serve) QueryUnits(query QueryUnitsQuery) (
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

func (serve *serve) AddPlayer(command AddPlayerCommand) (
	items []itemmodel.ItemAgg, players []playermodel.PlayerAgg, visionBound commonmodel.BoundVo, units []unitmodel.UnitAgg, err error,
) {
	err = serve.gameService.AddPlayer(command.WorldId, command.PlayerId)
	if err != nil {
		return items, players, visionBound, units, err
	}

	newPlayer, err := serve.playerRepository.Get(command.PlayerId)
	if err != nil {
		return items, players, visionBound, units, err
	}

	items, err = serve.itemRepository.GetAll()
	if err != nil {
		return items, players, visionBound, units, err
	}

	players, err = serve.playerRepository.GetPlayersAround(command.WorldId, newPlayer.GetPosition())
	if err != nil {
		return items, players, visionBound, units, err
	}

	visionBound = newPlayer.GetVisionBound()

	units, err = serve.unitRepository.GetUnitsInBound(command.WorldId, visionBound)
	if err != nil {
		return items, players, visionBound, units, err
	}

	err = serve.publishPlayersUpdatedEventToNearPlayers(command.WorldId, command.PlayerId)
	if err != nil {
		return items, players, visionBound, units, err
	}
	return items, players, visionBound, units, err
}

func (serve *serve) MovePlayer(command MovePlayerCommand) error {
	isVisionBoundUpdated, err := serve.gameService.MovePlayer(command.WorldId, command.PlayerId, command.Direction)
	if err != nil {
		return err
	}

	err = serve.publishPlayersUpdatedEventToNearPlayers(command.WorldId, command.PlayerId)
	if err != nil {
		return err
	}

	if isVisionBoundUpdated {
		err = serve.intEventPublisher.Publish(
			NewVisionBoundUpdatedIntEventChannel(command.WorldId.Uuid(), command.PlayerId.Uuid()),
			VisionBoundUpdatedIntEvent{},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (serve *serve) RemovePlayer(command RemovePlayerCommand) error {
	err := serve.gameService.RemovePlayer(command.WorldId, command.PlayerId)
	if err != nil {
		return err
	}

	err = serve.publishPlayersUpdatedEventToNearPlayers(command.WorldId, command.PlayerId)
	if err != nil {
		return err
	}
	return nil
}

func (serve *serve) PlaceItem(command PlaceItemCommand) error {
	err := serve.gameService.PlaceItem(command.WorldId, command.PlayerId, command.ItemId)
	if err != nil {
		return err
	}

	err = serve.publishUnitsUpdatedEventToNearPlayers(command.PlayerId)
	if err != nil {
		return err
	}
	return nil
}

func (serve *serve) DestroyItem(command DestroyItemCommand) error {
	err := serve.gameService.DestroyItem(command.WorldId, command.PlayerId)
	if err != nil {
		return err
	}

	err = serve.publishUnitsUpdatedEventToNearPlayers(command.PlayerId)
	if err != nil {
		return err
	}
	return nil
}
