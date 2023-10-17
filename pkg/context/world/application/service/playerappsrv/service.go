package playerappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

var (
	ErrPosHasNoUnits       = fmt.Errorf("the position has no units")
	ErrUnitIsNotPortalType = fmt.Errorf("the unit is not in type of portal")
)

type Service interface {
	GetPlayers(GetPlayersQuery) (playerDtos []dto.PlayerDto, err error)
	GetPlayer(GetPlayerQuery) (dto.PlayerDto, error)
	EnterWorld(EnterWorldCommand) (playerId uuid.UUID, err error)
	MovePlayer(MovePlayerCommand) error
	SendPlayerIntoPortal(SendPlayerIntoPortalCommand) error
	LeaveWorld(LeaveWorldCommand) error
	ChangePlayerHeldItem(ChangePlayerHeldItemCommand) error
}

type serve struct {
	playerRepo     playermodel.PlayerRepo
	unitRepo       unitmodel.UnitRepo
	portalUnitRepo portalunitmodel.PortalUnitRepo
}

func NewService(
	playerRepo playermodel.PlayerRepo,
	unitRepo unitmodel.UnitRepo,
	portalUnitRepo portalunitmodel.PortalUnitRepo,
) Service {
	return &serve{
		playerRepo:     playerRepo,
		unitRepo:       unitRepo,
		portalUnitRepo: portalUnitRepo,
	}
}

func (serve *serve) GetPlayers(query GetPlayersQuery) (
	playerDtos []dto.PlayerDto, err error,
) {
	players, err := serve.playerRepo.GetPlayersOfWorld(globalcommonmodel.NewWorldId(query.WorldId))
	if err != nil {
		return playerDtos, err
	}
	playerDtos = lo.Map(players, func(_player playermodel.Player, _ int) dto.PlayerDto {
		return dto.NewPlayerDto(_player)
	})

	return playerDtos, nil
}

func (serve *serve) GetPlayer(query GetPlayerQuery) (playerDto dto.PlayerDto, err error) {
	player, err := serve.playerRepo.Get(globalcommonmodel.NewWorldId(query.WorldId), playermodel.NewPlayerId(query.PlayerId))
	if err != nil {
		return playerDto, err
	}
	return dto.NewPlayerDto(player), nil
}

func (serve *serve) EnterWorld(command EnterWorldCommand) (plyaerIdDto uuid.UUID, err error) {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	playerHeldItemId := worldcommonmodel.NewItemId(command.PlayerHeldItemId)

	direction := worldcommonmodel.NewDownDirection()
	newPlayer := playermodel.NewPlayer(
		worldId,
		command.PlayerName,
		worldcommonmodel.NewPosition(0, 0),
		direction,
		&playerHeldItemId,
	)

	if err = serve.playerRepo.Add(newPlayer); err != nil {
		return plyaerIdDto, err
	}
	return newPlayer.GetId().Uuid(), nil
}

func (serve *serve) MovePlayer(command MovePlayerCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	playerId := playermodel.NewPlayerId(command.PlayerId)
	position := worldcommonmodel.NewPosition(command.Position.X, command.Position.Z)
	direction := worldcommonmodel.NewDirection(command.Direction)

	player, err := serve.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}

	player.Move(position, direction)

	return serve.playerRepo.Update(player)
}

func (serve *serve) SendPlayerIntoPortal(command SendPlayerIntoPortalCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	playerId := playermodel.NewPlayerId(command.PlayerId)
	position := worldcommonmodel.NewPosition(command.Position.X, command.Position.Z)

	player, err := serve.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}

	unitId := unitmodel.NewUnitId(worldId, position)

	unitAtPosition, err := serve.unitRepo.Find(unitId)
	if err != nil {
		return err
	}

	if unitAtPosition == nil {
		return ErrPosHasNoUnits
	}

	if !unitAtPosition.GetType().IsPortal() {
		return ErrUnitIsNotPortalType
	}

	portalUnit, err := serve.portalUnitRepo.Get(unitId)
	if err != nil {
		return err
	}
	targetPosition := portalUnit.GetTargetPosition()
	if targetPosition != nil {
		player.Teleport(*targetPosition)
	} else {
		return nil
	}

	return serve.playerRepo.Update(player)
}

func (serve *serve) LeaveWorld(command LeaveWorldCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	playerId := playermodel.NewPlayerId(command.PlayerId)

	player, err := serve.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}
	return serve.playerRepo.Delete(player)
}

func (serve *serve) ChangePlayerHeldItem(command ChangePlayerHeldItemCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	playerId := playermodel.NewPlayerId(command.PlayerId)
	itemId := worldcommonmodel.NewItemId(command.ItemId)

	player, err := serve.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}

	player.ChangeHeldItem(itemId)
	return serve.playerRepo.Update(player)
}
