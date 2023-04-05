package gameappservice

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/jsondto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/service"
	"github.com/samber/lo"
)

type Service interface {
	GetNearbyPlayers(GetNearbyPlayersQuery) (myPlayerDto jsondto.PlayerAggDto, ohterPlayerDtos []jsondto.PlayerAggDto, err error)
	GetNearbyUnits(GetNearbyUnitsQuery) (visionBoundDto jsondto.BoundVoDto, unitDtos []jsondto.UnitVoDto, err error)
	GetPlayer(GetPlayerQuery) (jsondto.PlayerAggDto, error)
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
	myPlayerDto jsondto.PlayerAggDto, otherPlayerDtos []jsondto.PlayerAggDto, err error,
) {
	player, err := serve.playerRepository.Get(playermodel.NewPlayerIdVo(query.PlayerId))
	if err != nil {
		return myPlayerDto, otherPlayerDtos, err
	}

	players, err := serve.playerRepository.GetPlayersAround(worldmodel.NewWorldIdVo(query.WorldId), player.GetPosition())
	if err != nil {
		return myPlayerDto, otherPlayerDtos, err
	}

	myPlayer, myPlayrFound := lo.Find(players, func(_player playermodel.PlayerAgg) bool {
		return _player.GetId().IsEqual(player.GetId())
	})
	if !myPlayrFound {
		return myPlayerDto, otherPlayerDtos, errors.New("my player not found in players")
	}
	myPlayerDto = jsondto.NewPlayerAggDto(myPlayer)
	otherPlayers := lo.Filter(players, func(_player playermodel.PlayerAgg, _ int) bool {
		return !_player.GetId().IsEqual(player.GetId())
	})
	otherPlayerDtos = lo.Map(otherPlayers, func(_player playermodel.PlayerAgg, _ int) jsondto.PlayerAggDto {
		return jsondto.NewPlayerAggDto(_player)
	})

	return myPlayerDto, otherPlayerDtos, err
}

func (serve *serve) GetNearbyUnits(query GetNearbyUnitsQuery) (
	visionBoundDto jsondto.BoundVoDto, unitDtos []jsondto.UnitVoDto, err error,
) {
	player, err := serve.playerRepository.Get(playermodel.NewPlayerIdVo(query.PlayerId))
	if err != nil {
		return visionBoundDto, unitDtos, err
	}

	visionBound := player.GetVisionBound()
	visionBoundDto = jsondto.NewBoundVoDto(visionBound)
	units, err := serve.unitRepository.GetUnitsInBound(worldmodel.NewWorldIdVo(query.WorldId), visionBound)
	if err != nil {
		return visionBoundDto, unitDtos, err
	}
	unitDtos = lo.Map(units, func(unit unitmodel.UnitAgg, _ int) jsondto.UnitVoDto {
		return jsondto.NewUnitVoDto(unit)
	})

	return visionBoundDto, unitDtos, err
}

func (serve *serve) GetPlayer(query GetPlayerQuery) (playerDto jsondto.PlayerAggDto, err error) {
	player, err := serve.playerRepository.Get(playermodel.NewPlayerIdVo(query.PlayerId))
	if err != nil {
		return playerDto, err
	}
	return jsondto.NewPlayerAggDto(player), nil
}

func (serve *serve) EnterWorld(command EnterWorldCommand) error {
	return serve.gameService.EnterWorld(worldmodel.NewWorldIdVo(command.WorldId), playermodel.NewPlayerIdVo(command.PlayerId))
}

func (serve *serve) Move(command MoveCommand) error {
	return serve.gameService.Move(worldmodel.NewWorldIdVo(command.WorldId), playermodel.NewPlayerIdVo(command.PlayerId), commonmodel.NewDirectionVo(command.Direction))
}

func (serve *serve) LeaveWorld(command LeaveWorldCommand) error {
	return serve.gameService.LeaveWorld(worldmodel.NewWorldIdVo(command.WorldId), playermodel.NewPlayerIdVo(command.PlayerId))
}

func (serve *serve) PlaceItem(command PlaceItemCommand) error {
	return serve.gameService.PlaceItem(worldmodel.NewWorldIdVo(command.WorldId), playermodel.NewPlayerIdVo(command.PlayerId))
}

func (serve *serve) ChangeHeldItem(command ChangeHeldItemCommand) error {
	return serve.gameService.ChangeHeldItem(worldmodel.NewWorldIdVo(command.WorldId), playermodel.NewPlayerIdVo(command.PlayerId), itemmodel.NewItemIdVo(command.ItemId))
}

func (serve *serve) RemoveItem(command RemoveItemCommand) error {
	return serve.gameService.RemoveItem(worldmodel.NewWorldIdVo(command.WorldId), playermodel.NewPlayerIdVo(command.PlayerId))
}
