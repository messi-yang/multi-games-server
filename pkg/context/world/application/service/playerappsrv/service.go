package playerappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/playermodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	GetPlayers(GetPlayersQuery) (playerDtos []dto.PlayerDto, err error)
	GetPlayer(GetPlayerQuery) (dto.PlayerDto, error)
	EnterWorld(EnterWorldCommand) (playerId uuid.UUID, err error)
	Move(MoveCommand) error
	LeaveWorld(LeaveWorldCommand) error
	ChangeHeldItem(ChangeHeldItemCommand) error
}

type serve struct {
	playerRepo playermodel.PlayerRepo
}

func NewService(
	playerRepo playermodel.PlayerRepo,
) Service {
	return &serve{
		playerRepo: playerRepo,
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

func (serve *serve) Move(command MoveCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	playerId := playermodel.NewPlayerId(command.PlayerId)
	direction := worldcommonmodel.NewDirection(command.Direction)

	player, err := serve.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}

	if !direction.IsEqual(player.GetDirection()) {
		player.Move(player.GetPosition(), direction)
		return serve.playerRepo.Update(player)
	}

	newItemPos := player.GetPositionOneStepFoward()

	player.Move(newItemPos, direction)

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

func (serve *serve) ChangeHeldItem(command ChangeHeldItemCommand) error {
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
