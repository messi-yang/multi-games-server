package worldjourneyhandler

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type serverMessageName string

const (
	unitCreatedServerMessageName  serverMessageName = "UNIT_CREATED"
	unitDeletedServerMessageName  serverMessageName = "UNIT_DELETED"
	playerJoinedServerMessageName serverMessageName = "PLAYER_JOINED"
	playerLeftServerMessageName   serverMessageName = "PLAYER_LEFT"
	playerMovedServerMessageName  serverMessageName = "PLAYER_MOVED"
)

type ServerMessage struct {
	Name serverMessageName
}

type unitCreatedServerMessage struct {
	Name serverMessageName
	Unit dto.UnitDto
}

func newunitCreatedServerMessage(unit dto.UnitDto) unitCreatedServerMessage {
	return unitCreatedServerMessage{
		Name: unitCreatedServerMessageName,
		Unit: unit,
	}
}

type unitDeletedServerMessage struct {
	Name     serverMessageName
	WorldId  uuid.UUID
	Position dto.PositionDto
}

func newunitDeletedServerMessage(worldId uuid.UUID, position dto.PositionDto) unitDeletedServerMessage {
	return unitDeletedServerMessage{
		Name:     unitDeletedServerMessageName,
		WorldId:  worldId,
		Position: position,
	}
}

type playerJoinedServerMessage struct {
	Name   serverMessageName
	Player dto.PlayerDto
}

func newplayerJoinedServerMessage(player dto.PlayerDto) playerJoinedServerMessage {
	return playerJoinedServerMessage{
		Name:   playerJoinedServerMessageName,
		Player: player,
	}
}

type playerLeftServerMessage struct {
	Name     serverMessageName
	PlayerId uuid.UUID
}

func newplayerLeftServerMessage(playerId uuid.UUID) playerLeftServerMessage {
	return playerLeftServerMessage{
		Name:     playerLeftServerMessageName,
		PlayerId: playerId,
	}
}

type playerMovedServerMessage struct {
	Name   serverMessageName
	Player dto.PlayerDto
}

func newPlayerMovedServerMessage(player dto.PlayerDto) playerMovedServerMessage {
	return playerMovedServerMessage{
		Name:   playerMovedServerMessageName,
		Player: player,
	}
}

func newWorldServerMessageChannel(worldIdDto uuid.UUID) string {
	return fmt.Sprintf("WORLD_%s_CHANNEL", worldIdDto)
}
