package worldjourneyappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/dto"
	"github.com/google/uuid"
)

type ServerMessageName string

const (
	UnitCreated  ServerMessageName = "UNIT_CREATED"
	UnitDeleted  ServerMessageName = "UNIT_DELETED"
	PlayerJoined ServerMessageName = "PLAYER_JOINED"
	PlayerLeft   ServerMessageName = "PLAYER_LEFT"
	PlayerMoved  ServerMessageName = "PLAYER_MOVED"
)

type ServerMessage struct {
	Name ServerMessageName
}

type UnitCreatedServerMessage struct {
	Name ServerMessageName
	Unit dto.UnitDto
}

func NewUnitCreatedServerMessage(unit dto.UnitDto) UnitCreatedServerMessage {
	return UnitCreatedServerMessage{
		Name: UnitCreated,
		Unit: unit,
	}
}

type UnitDeletedServerMessage struct {
	Name     ServerMessageName
	WorldId  uuid.UUID
	Position dto.PositionDto
}

func NewUnitDeletedServerMessage(worldId uuid.UUID, position dto.PositionDto) UnitDeletedServerMessage {
	return UnitDeletedServerMessage{
		Name:     UnitDeleted,
		WorldId:  worldId,
		Position: position,
	}
}

type PlayerJoinedServerMessage struct {
	Name   ServerMessageName
	Player dto.PlayerDto
}

func NewPlayerJoinedServerMessage(player dto.PlayerDto) PlayerJoinedServerMessage {
	return PlayerJoinedServerMessage{
		Name:   PlayerJoined,
		Player: player,
	}
}

type PlayerLeftServerMessage struct {
	Name     ServerMessageName
	PlayerId uuid.UUID
}

func NewPlayerLeftServerMessage(playerId uuid.UUID) PlayerLeftServerMessage {
	return PlayerLeftServerMessage{
		Name:     PlayerLeft,
		PlayerId: playerId,
	}
}

type PlayerMovedServerMessage struct {
	Name   ServerMessageName
	Player dto.PlayerDto
}

func NewPlayerMovedServerMessage(player dto.PlayerDto) PlayerMovedServerMessage {
	return PlayerMovedServerMessage{
		Name:   PlayerMoved,
		Player: player,
	}
}

func NewWorldServerMessageChannel(worldIdDto uuid.UUID) string {
	return fmt.Sprintf("WORLD_%s_CHANNEL", worldIdDto)
}
