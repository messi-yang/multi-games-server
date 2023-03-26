package gamesocketappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/service"
	"github.com/samber/lo"
)

type Service interface {
	GetPlayersAroundPlayer(presenter Presenter, query GetPlayersQuery) error
	GetUnitsVisibleByPlayer(presenter Presenter, query GetUnitsVisibleByPlayerQuery) error
	AddPlayer(presenter Presenter, command AddPlayerCommand) error
	MovePlayer(presenter Presenter, command MovePlayerCommand) error
	RemovePlayer(presenter Presenter, command RemovePlayerCommand) error
	PlaceItem(presenter Presenter, command PlaceItemCommand) error
	DestroyItem(presenter Presenter, command DestroyItemCommand) error
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
		NewUnitsUpdatedIntEventChannel(worldId.String(), playerId.String()),
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
			NewPlayersUpdatedIntEventChannel(worldId.String(), player.GetId().String()),
			PlayersUpdatedIntEvent{},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (serve *serve) GetPlayersAroundPlayer(presenter Presenter, query GetPlayersQuery) error {
	worldId, playerId, err := query.Validate()
	if err != nil {
		return err
	}

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	players, err := serve.playerRepository.GetPlayersAround(worldId, player.GetPosition())
	if err != nil {
		return err
	}

	return presenter.OnMessage(PlayersUpdatedResponseDto{
		Type: PlayersUpdatedResponseDtoType,
		Players: lo.Map(players, func(player playermodel.PlayerAgg, _ int) dto.PlayerAggDto {
			return dto.NewPlayerAggDto(player)
		}),
	})
}

func (serve *serve) GetUnitsVisibleByPlayer(presenter Presenter, query GetUnitsVisibleByPlayerQuery) error {
	worldId, playerId, err := query.Validate()
	if err != nil {
		return err
	}

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	playerVisionBound := player.GetVisionBound()
	units, err := serve.unitRepository.GetUnitsInBound(worldId, playerVisionBound)
	if err != nil {
		return err
	}

	return presenter.OnMessage(UnitsUpdatedResponseDto{
		Type:        UnitsUpdatedResponseDtoType,
		VisionBound: dto.NewBoundVoDto(playerVisionBound),
		Units: lo.Map(units, func(unit unitmodel.UnitAgg, _ int) dto.UnitVoDto {
			return dto.NewUnitVoDto(unit)
		}),
	})
}

func (serve *serve) AddPlayer(presenter Presenter, command AddPlayerCommand) error {
	worldId, playerId, err := command.Validate()
	if err != nil {
		return err
	}

	err = serve.gameService.AddPlayer(worldId, playerId)
	if err != nil {
		return err
	}

	newPlayer, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return err
	}

	items, err := serve.itemRepository.GetAll()
	if err != nil {
		return err
	}

	itemDtos := lo.Map(items, func(item itemmodel.ItemAgg, _ int) dto.ItemAggDto {
		return dto.NewItemAggDto(item)
	})

	players, err := serve.playerRepository.GetPlayersAround(worldId, newPlayer.GetPosition())
	if err != nil {
		return err
	}

	playerDtos := lo.Map(players, func(p playermodel.PlayerAgg, _ int) dto.PlayerAggDto {
		return dto.NewPlayerAggDto(p)
	})

	playerVisionBound := newPlayer.GetVisionBound()
	units, err := serve.unitRepository.GetUnitsInBound(worldId, playerVisionBound)
	if err != nil {
		return err
	}

	err = presenter.OnMessage(GameJoinedResponseDto{
		Type:        GameJoinedResponseDtoType,
		Items:       itemDtos,
		PlayerId:    playerId.String(),
		Players:     playerDtos,
		VisionBound: dto.NewBoundVoDto(playerVisionBound),
		Units: lo.Map(units, func(unit unitmodel.UnitAgg, _ int) dto.UnitVoDto {
			return dto.NewUnitVoDto(unit)
		}),
	})
	if err != nil {
		return err
	}

	return serve.publishPlayersUpdatedEventToNearPlayers(worldId, playerId)
}

func (serve *serve) MovePlayer(presenter Presenter, command MovePlayerCommand) error {
	worldId, playerId, direction, err := command.Validate()
	if err != nil {
		return err
	}

	isVisionBoundUpdated, err := serve.gameService.MovePlayer(worldId, playerId, direction)
	if err != nil {
		return err
	}

	err = serve.publishPlayersUpdatedEventToNearPlayers(worldId, playerId)
	if err != nil {
		return err
	}

	if isVisionBoundUpdated {
		return serve.intEventPublisher.Publish(
			NewVisionBoundUpdatedIntEventChannel(worldId.String(), playerId.String()),
			VisionBoundUpdatedIntEvent{},
		)
	}
	return nil
}

func (serve *serve) RemovePlayer(presenter Presenter, command RemovePlayerCommand) error {
	worldId, playerId, err := command.Validate()
	if err != nil {
		return err
	}

	err = serve.gameService.RemovePlayer(worldId, playerId)
	if err != nil {
		return err
	}

	return serve.publishPlayersUpdatedEventToNearPlayers(worldId, playerId)
}

func (serve *serve) PlaceItem(presenter Presenter, command PlaceItemCommand) error {
	worldId, playerId, itemId, err := command.Validate()
	if err != nil {
		return err
	}

	err = serve.gameService.PlaceItem(worldId, playerId, itemId)
	if err != nil {
		return err
	}

	return serve.publishUnitsUpdatedEventToNearPlayers(playerId)
}

func (serve *serve) DestroyItem(presenter Presenter, command DestroyItemCommand) error {
	worldId, playerId, err := command.Validate()
	if err != nil {
		return err
	}

	err = serve.gameService.DestroyItem(worldId, playerId)
	if err != nil {
		return err
	}

	return serve.publishUnitsUpdatedEventToNearPlayers(playerId)
}
