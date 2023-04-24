package gameappsrv

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/jsondto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/service/gamedomainsrv"
	"github.com/samber/lo"
)

type Service interface {
	GetNearbyPlayers(GetNearbyPlayersQuery) (myPlayerDto jsondto.PlayerAggDto, ohterPlayerDtos []jsondto.PlayerAggDto, err error)
	GetNearbyUnits(GetNearbyUnitsQuery) (unitDtos []jsondto.UnitAggDto, err error)
	GetPlayer(GetPlayerQuery) (jsondto.PlayerAggDto, error)
	EnterWorld(EnterWorldCommand) error
	Move(MoveCommand) error
	LeaveWorld(LeaveWorldCommand) error
	ChangeHeldItem(ChangeHeldItemCommand) error
	PlaceItem(PlaceItemCommand) error
	RemoveItem(RemoveItemCommand) error
}

type serve struct {
	worldRepo         worldmodel.Repo
	playerRepo        playermodel.Repo
	unitRepo          unitmodel.Repo
	itemRepo          itemmodel.Repo
	gameDomainService gamedomainsrv.Service
}

func NewService(worldRepo worldmodel.Repo, playerRepo playermodel.Repo, unitRepo unitmodel.Repo, itemRepo itemmodel.Repo, gameDomainService gamedomainsrv.Service) Service {
	return &serve{
		worldRepo:         worldRepo,
		playerRepo:        playerRepo,
		unitRepo:          unitRepo,
		itemRepo:          itemRepo,
		gameDomainService: gameDomainService,
	}
}

func (serve *serve) GetNearbyPlayers(query GetNearbyPlayersQuery) (
	myPlayerDto jsondto.PlayerAggDto, otherPlayerDtos []jsondto.PlayerAggDto, err error,
) {
	player, err := serve.playerRepo.Get(commonmodel.NewPlayerIdVo(query.PlayerId))
	if err != nil {
		return myPlayerDto, otherPlayerDtos, err
	}

	players, err := serve.playerRepo.GetPlayersAround(commonmodel.NewWorldIdVo(query.WorldId), player.GetPosition())
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
	unitDtos []jsondto.UnitAggDto, err error,
) {
	player, err := serve.playerRepo.Get(commonmodel.NewPlayerIdVo(query.PlayerId))
	if err != nil {
		return unitDtos, err
	}

	visionBound := player.GetVisionBound()
	units, err := serve.unitRepo.GetUnitsInBound(commonmodel.NewWorldIdVo(query.WorldId), visionBound)
	if err != nil {
		return unitDtos, err
	}
	unitDtos = lo.Map(units, func(unit unitmodel.UnitAgg, _ int) jsondto.UnitAggDto {
		return jsondto.NewUnitAggDto(unit)
	})

	return unitDtos, err
}

func (serve *serve) GetPlayer(query GetPlayerQuery) (playerDto jsondto.PlayerAggDto, err error) {
	player, err := serve.playerRepo.Get(commonmodel.NewPlayerIdVo(query.PlayerId))
	if err != nil {
		return playerDto, err
	}
	return jsondto.NewPlayerAggDto(player), nil
}

func (serve *serve) EnterWorld(command EnterWorldCommand) error {
	return serve.gameDomainService.EnterWorld(commonmodel.NewWorldIdVo(command.WorldId), commonmodel.NewPlayerIdVo(command.PlayerId))
}

func (serve *serve) Move(command MoveCommand) error {
	return serve.gameDomainService.Move(commonmodel.NewWorldIdVo(command.WorldId), commonmodel.NewPlayerIdVo(command.PlayerId), commonmodel.NewDirectionVo(command.Direction))
}

func (serve *serve) LeaveWorld(command LeaveWorldCommand) error {
	return serve.gameDomainService.LeaveWorld(commonmodel.NewWorldIdVo(command.WorldId), commonmodel.NewPlayerIdVo(command.PlayerId))
}

func (serve *serve) PlaceItem(command PlaceItemCommand) error {
	return serve.gameDomainService.PlaceItem(commonmodel.NewWorldIdVo(command.WorldId), commonmodel.NewPlayerIdVo(command.PlayerId))
}

func (serve *serve) ChangeHeldItem(command ChangeHeldItemCommand) error {
	return serve.gameDomainService.ChangeHeldItem(commonmodel.NewWorldIdVo(command.WorldId), commonmodel.NewPlayerIdVo(command.PlayerId), commonmodel.NewItemIdVo(command.ItemId))
}

func (serve *serve) RemoveItem(command RemoveItemCommand) error {
	return serve.gameDomainService.RemoveItem(commonmodel.NewWorldIdVo(command.WorldId), commonmodel.NewPlayerIdVo(command.PlayerId))
}
