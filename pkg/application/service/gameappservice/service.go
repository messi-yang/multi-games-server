package gameappservice

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
	GetPlayersAroundPlayer(query GetPlayersAroundPlayerQuery) ([]dto.PlayerAggDto, error)
	GetUnitsVisibleByPlayer(query GetUnitsVisibleByPlayerQuery) (dto.BoundVoDto, []dto.UnitVoDto, error)
	AddPlayer(command AddPlayerCommand) ([]dto.ItemAggDto, []dto.PlayerAggDto, dto.BoundVoDto, []dto.UnitVoDto, error)
	MovePlayer(command MovePlayerCommand) error
	RemovePlayer(command RemovePlayerCommand) error
	PlaceItem(command PlaceItemCommand) error
	DestroyItem(command DestroyItemCommand) error
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

func (serve *serve) GetPlayersAroundPlayer(query GetPlayersAroundPlayerQuery) (
	playerDtos []dto.PlayerAggDto, err error,
) {
	worldId, playerId := query.Parse()

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return playerDtos, err
	}

	players, err := serve.playerRepository.GetPlayersAround(worldId, player.GetPosition())
	if err != nil {
		return playerDtos, err
	}

	playerDtos = lo.Map(players, func(player playermodel.PlayerAgg, _ int) dto.PlayerAggDto {
		return dto.NewPlayerAggDto(player)
	})

	return playerDtos, nil
}

func (serve *serve) GetUnitsVisibleByPlayer(query GetUnitsVisibleByPlayerQuery) (
	boundDto dto.BoundVoDto, unitDtos []dto.UnitVoDto, err error,
) {
	worldId, playerId := query.Parse()

	player, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return boundDto, unitDtos, err
	}

	visionBound := player.GetVisionBound()
	units, err := serve.unitRepository.GetUnitsInBound(worldId, visionBound)
	if err != nil {
		return boundDto, unitDtos, err
	}

	boundDto = dto.NewBoundVoDto(visionBound)
	unitDtos = lo.Map(units, func(unit unitmodel.UnitAgg, _ int) dto.UnitVoDto {
		return dto.NewUnitVoDto(unit)
	})

	return boundDto, unitDtos, nil
}

func (serve *serve) AddPlayer(command AddPlayerCommand) (
	itemDtos []dto.ItemAggDto, playerDtos []dto.PlayerAggDto, visionBoundDto dto.BoundVoDto, unitDtos []dto.UnitVoDto, err error,
) {
	worldId, playerId := command.Parse()

	err = serve.gameService.AddPlayer(worldId, playerId)
	if err != nil {
		return itemDtos, playerDtos, visionBoundDto, unitDtos, err
	}

	newPlayer, err := serve.playerRepository.Get(playerId)
	if err != nil {
		return itemDtos, playerDtos, visionBoundDto, unitDtos, err
	}

	items, err := serve.itemRepository.GetAll()
	if err != nil {
		return itemDtos, playerDtos, visionBoundDto, unitDtos, err
	}

	players, err := serve.playerRepository.GetPlayersAround(worldId, newPlayer.GetPosition())
	if err != nil {
		return itemDtos, playerDtos, visionBoundDto, unitDtos, err
	}

	visionBound := newPlayer.GetVisionBound()

	units, err := serve.unitRepository.GetUnitsInBound(worldId, visionBound)
	if err != nil {
		return itemDtos, playerDtos, visionBoundDto, unitDtos, err
	}

	itemDtos = lo.Map(items, func(item itemmodel.ItemAgg, _ int) dto.ItemAggDto {
		return dto.NewItemAggDto(item)
	})

	playerDtos = lo.Map(players, func(p playermodel.PlayerAgg, _ int) dto.PlayerAggDto {
		return dto.NewPlayerAggDto(p)
	})

	unitDtos = lo.Map(units, func(unit unitmodel.UnitAgg, _ int) dto.UnitVoDto {
		return dto.NewUnitVoDto(unit)
	})

	visionBoundDto = dto.NewBoundVoDto(visionBound)

	err = serve.publishPlayersUpdatedEventToNearPlayers(worldId, playerId)
	if err != nil {
		return itemDtos, playerDtos, visionBoundDto, unitDtos, err
	}
	return itemDtos, playerDtos, visionBoundDto, unitDtos, nil
}

func (serve *serve) MovePlayer(command MovePlayerCommand) error {
	worldId, playerId, direction := command.Parse()

	isVisionBoundUpdated, err := serve.gameService.MovePlayer(worldId, playerId, direction)
	if err != nil {
		return err
	}

	err = serve.publishPlayersUpdatedEventToNearPlayers(worldId, playerId)
	if err != nil {
		return err
	}

	if isVisionBoundUpdated {
		err = serve.intEventPublisher.Publish(
			NewVisionBoundUpdatedIntEventChannel(worldId.Uuid(), playerId.Uuid()),
			VisionBoundUpdatedIntEvent{},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (serve *serve) RemovePlayer(command RemovePlayerCommand) error {
	worldId, playerId := command.Parse()

	err := serve.gameService.RemovePlayer(worldId, playerId)
	if err != nil {
		return err
	}

	err = serve.publishPlayersUpdatedEventToNearPlayers(worldId, playerId)
	if err != nil {
		return err
	}
	return nil
}

func (serve *serve) PlaceItem(command PlaceItemCommand) error {
	worldId, playerId, itemId := command.Parse()

	err := serve.gameService.PlaceItem(worldId, playerId, itemId)
	if err != nil {
		return err
	}

	err = serve.publishUnitsUpdatedEventToNearPlayers(playerId)
	if err != nil {
		return err
	}
	return nil
}

func (serve *serve) DestroyItem(command DestroyItemCommand) error {
	worldId, playerId := command.Parse()

	err := serve.gameService.DestroyItem(worldId, playerId)
	if err != nil {
		return err
	}

	err = serve.publishUnitsUpdatedEventToNearPlayers(playerId)
	if err != nil {
		return err
	}
	return nil
}
