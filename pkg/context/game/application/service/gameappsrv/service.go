package gameappsrv

import (
	"errors"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/worldmodel/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	GetNearbyPlayers(GetNearbyPlayersQuery) (myPlayerDto dto.PlayerDto, ohterPlayerDtos []dto.PlayerDto, err error)
	GetNearbyUnits(GetNearbyUnitsQuery) (unitDtos []dto.UnitDto, err error)
	GetUnit(GetUnitQuery) (dto.UnitDto, error)
	GetPlayer(GetPlayerQuery) (dto.PlayerDto, error)
	EnterWorld(EnterWorldCommand) (playerId uuid.UUID, err error)
	Move(MoveCommand) error
	LeaveWorld(LeaveWorldCommand) error
	ChangeHeldItem(ChangeHeldItemCommand) error
	PlaceItem(PlaceItemCommand) error
	RemoveItem(RemoveItemCommand) error
}

type serve struct {
	worldRepo         worldmodel.WorldRepo
	playerRepo        playermodel.PlayerRepo
	unitRepo          unitmodel.UnitRepo
	itemRepo          itemmodel.ItemRepo
	gameDomainService service.GameService
}

func NewService(
	worldRepo worldmodel.WorldRepo,
	playerRepo playermodel.PlayerRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
	gameDomainService service.GameService,
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
	player, err := serve.playerRepo.Get(playermodel.NewPlayerId(query.PlayerId))
	if err != nil {
		return myPlayerDto, otherPlayerDtos, err
	}

	players, err := serve.playerRepo.GetPlayersAround(sharedkernelmodel.NewWorldId(query.WorldId), player.GetPosition())
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
	player, err := serve.playerRepo.Get(playermodel.NewPlayerId(query.PlayerId))
	if err != nil {
		return unitDtos, err
	}

	visionBound := player.GetVisionBound()
	units, err := serve.unitRepo.QueryUnitsInBound(sharedkernelmodel.NewWorldId(query.WorldId), visionBound)
	if err != nil {
		return unitDtos, err
	}
	unitDtos = lo.Map(units, func(unit unitmodel.Unit, _ int) dto.UnitDto {
		return dto.NewUnitDto(unit)
	})

	return unitDtos, err
}

func (serve *serve) GetUnit(query GetUnitQuery) (unitDto dto.UnitDto, err error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	position := commonmodel.NewPosition(query.Position.X, query.Position.Z)
	unit, err := serve.unitRepo.Get(unitmodel.NewUnitId(worldId, position))
	if err != nil {
		return unitDto, err
	}
	return dto.NewUnitDto(unit), nil
}

func (serve *serve) GetPlayer(query GetPlayerQuery) (playerDto dto.PlayerDto, err error) {
	player, err := serve.playerRepo.Get(playermodel.NewPlayerId(query.PlayerId))
	if err != nil {
		return playerDto, err
	}
	return dto.NewPlayerDto(player), nil
}

func (serve *serve) EnterWorld(command EnterWorldCommand) (plyaerIdDto uuid.UUID, err error) {
	playerId, err := serve.gameDomainService.EnterWorld(sharedkernelmodel.NewWorldId(command.WorldId))
	if err != nil {
		return plyaerIdDto, err
	}
	return playerId.Uuid(), nil
}

func (serve *serve) Move(command MoveCommand) error {
	return serve.gameDomainService.Move(sharedkernelmodel.NewWorldId(command.WorldId), playermodel.NewPlayerId(command.PlayerId), commonmodel.NewDirection(command.Direction))
}

func (serve *serve) LeaveWorld(command LeaveWorldCommand) error {
	return serve.gameDomainService.LeaveWorld(sharedkernelmodel.NewWorldId(command.WorldId), playermodel.NewPlayerId(command.PlayerId))
}

func (serve *serve) PlaceItem(command PlaceItemCommand) error {
	return serve.gameDomainService.PlaceItem(sharedkernelmodel.NewWorldId(command.WorldId), playermodel.NewPlayerId(command.PlayerId))
}

func (serve *serve) ChangeHeldItem(command ChangeHeldItemCommand) error {
	return serve.gameDomainService.ChangeHeldItem(sharedkernelmodel.NewWorldId(command.WorldId), playermodel.NewPlayerId(command.PlayerId), commonmodel.NewItemId(command.ItemId))
}

func (serve *serve) RemoveItem(command RemoveItemCommand) error {
	return serve.gameDomainService.RemoveItem(sharedkernelmodel.NewWorldId(command.WorldId), playermodel.NewPlayerId(command.PlayerId))
}
