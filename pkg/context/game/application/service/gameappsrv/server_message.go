package gameappsrv

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
	Name     ServerMessageName
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func NewPlayerJoinedServerMessage(worldId uuid.UUID, playerId uuid.UUID) PlayerJoinedServerMessage {
	return PlayerJoinedServerMessage{
		Name:     PlayerJoined,
		WorldId:  worldId,
		PlayerId: playerId,
	}
}

type PlayerLeftServerMessage struct {
	Name     ServerMessageName
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func NewPlayerLeftServerMessage(worldId uuid.UUID, playerId uuid.UUID) PlayerLeftServerMessage {
	return PlayerLeftServerMessage{
		Name:     PlayerLeft,
		WorldId:  worldId,
		PlayerId: playerId,
	}
}

type PlayerMovedServerMessage struct {
	Name     ServerMessageName
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func NewPlayerMovedServerMessage(worldId uuid.UUID, playerId uuid.UUID) PlayerMovedServerMessage {
	return PlayerMovedServerMessage{
		Name:     PlayerMoved,
		WorldId:  worldId,
		PlayerId: playerId,
	}
}

func NewWorldServerMessageChannel(worldIdDto uuid.UUID) string {
	return fmt.Sprintf("WORLD_%s_CHANNEL", worldIdDto)
}
