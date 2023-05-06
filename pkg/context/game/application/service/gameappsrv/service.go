package gameappsrv

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/service/gamedomainsrv"
	"github.com/samber/lo"
)

type Service interface {
	GetNearbyPlayers(GetNearbyPlayersQuery) (myPlayerDto dto.PlayerDto, ohterPlayerDtos []dto.PlayerDto, err error)
	GetNearbyUnits(GetNearbyUnitsQuery) (unitDtos []dto.UnitDto, err error)
	GetPlayer(GetPlayerQuery) (dto.PlayerDto, error)
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

func NewService(
	worldRepo worldmodel.Repo,
	playerRepo playermodel.Repo,
	unitRepo unitmodel.Repo,
	itemRepo itemmodel.Repo,
	gameDomainService gamedomainsrv.Service,
) Service {
	return &serve{
		worldRepo:         worldRepo,
		playerRepo:        playerRepo,
		unitRepo:          unitRepo,
		itemRepo:          itemRepo,
		gameDomainService: gameDomainService,
	}
}

func (serve *serve) GetNearbyPlayers(query GetNearbyPlayersQuery) (
	myPlayerDto dto.PlayerDto, otherPlayerDtos []dto.PlayerDto, err error,
) {
	player, err := serve.playerRepo.Get(commonmodel.NewPlayerId(query.PlayerId))
	if err != nil {
		return myPlayerDto, otherPlayerDtos, err
	}

	players, err := serve.playerRepo.GetPlayersAround(commonmodel.NewWorldId(query.WorldId), player.GetPosition())
	if err != nil {
		return myPlayerDto, otherPlayerDtos, err
	}

	myPlayer, myPlayrFound := lo.Find(players, func(_player playermodel.Player) bool {
		return _player.GetId().IsEqual(player.GetId())
	})
	if !myPlayrFound {
		return myPlayerDto, otherPlayerDtos, errors.New("my player not found in players")
	}
	myPlayerDto = dto.NewPlayerDto(myPlayer)
	otherPlayers := lo.Filter(players, func(_player playermodel.Player, _ int) bool {
		return !_player.GetId().IsEqual(player.GetId())
	})
	otherPlayerDtos = lo.Map(otherPlayers, func(_player playermodel.Player, _ int) dto.PlayerDto {
		return dto.NewPlayerDto(_player)
	})

	return myPlayerDto, otherPlayerDtos, err
}

func (serve *serve) GetNearbyUnits(query GetNearbyUnitsQuery) (
	unitDtos []dto.UnitDto, err error,
) {
	player, err := serve.playerRepo.Get(commonmodel.NewPlayerId(query.PlayerId))
	if err != nil {
		return unitDtos, err
	}

	visionBound := player.GetVisionBound()
	units, err := serve.unitRepo.QueryUnitsInBound(commonmodel.NewWorldId(query.WorldId), visionBound)
	if err != nil {
		return unitDtos, err
	}
	unitDtos = lo.Map(units, func(unit unitmodel.Unit, _ int) dto.UnitDto {
		return dto.NewUnitDto(unit)
	})

	return unitDtos, err
}

func (serve *serve) GetPlayer(query GetPlayerQuery) (playerDto dto.PlayerDto, err error) {
	player, err := serve.playerRepo.Get(commonmodel.NewPlayerId(query.PlayerId))
	if err != nil {
		return playerDto, err
	}
	return dto.NewPlayerDto(player), nil
}

func (serve *serve) EnterWorld(command EnterWorldCommand) error {
	return serve.gameDomainService.EnterWorld(commonmodel.NewWorldId(command.WorldId), commonmodel.NewPlayerId(command.PlayerId))
}

func (serve *serve) Move(command MoveCommand) error {
	return serve.gameDomainService.Move(commonmodel.NewWorldId(command.WorldId), commonmodel.NewPlayerId(command.PlayerId), commonmodel.NewDirection(command.Direction))
}

func (serve *serve) LeaveWorld(command LeaveWorldCommand) error {
	return serve.gameDomainService.LeaveWorld(commonmodel.NewWorldId(command.WorldId), commonmodel.NewPlayerId(command.PlayerId))
}

func (serve *serve) PlaceItem(command PlaceItemCommand) error {
	return serve.gameDomainService.PlaceItem(commonmodel.NewWorldId(command.WorldId), commonmodel.NewPlayerId(command.PlayerId))
}

func (serve *serve) ChangeHeldItem(command ChangeHeldItemCommand) error {
	return serve.gameDomainService.ChangeHeldItem(commonmodel.NewWorldId(command.WorldId), commonmodel.NewPlayerId(command.PlayerId), commonmodel.NewItemId(command.ItemId))
}

func (serve *serve) RemoveItem(command RemoveItemCommand) error {
	return serve.gameDomainService.RemoveItem(commonmodel.NewWorldId(command.WorldId), commonmodel.NewPlayerId(command.PlayerId))
}
