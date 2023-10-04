package worldjourneyhandler

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type serverMessageName string

const (
	unitCreatedServerMessageName           serverMessageName = "UNIT_CREATED"
	unitRotatedServerMessageName           serverMessageName = "UNIT_ROTATED"
	unitRemovedServerMessageName           serverMessageName = "UNIT_REMOVED"
	playerJoinedServerMessageName          serverMessageName = "PLAYER_JOINED"
	playerLeftServerMessageName            serverMessageName = "PLAYER_LEFT"
	playerMovedServerMessageName           serverMessageName = "PLAYER_MOVED"
	playerHeldItemChangedServerMessageName serverMessageName = "PLAYER_HELD_ITEM_CHANGED"
)

type ServerMessage struct {
	Name serverMessageName
}

type unitCreatedServerMessage struct {
	Name serverMessageName
	Unit dto.UnitDto
}

func newUnitCreatedServerMessage(unit dto.UnitDto) unitCreatedServerMessage {
	return unitCreatedServerMessage{
		Name: unitCreatedServerMessageName,
		Unit: unit,
	}
}

type unitRotateServerMessage struct {
	Name     serverMessageName
	Position dto.PositionDto
}

func newUnitRotatedServerMessage(position dto.PositionDto) unitRotateServerMessage {
	return unitRotateServerMessage{
		Name:     unitRotatedServerMessageName,
		Position: position,
	}
}

type unitRemovedServerMessage struct {
	Name     serverMessageName
	WorldId  uuid.UUID
	Position dto.PositionDto
}

func newUnitRemovedServerMessage(worldId uuid.UUID, position dto.PositionDto) unitRemovedServerMessage {
	return unitRemovedServerMessage{
		Name:     unitRemovedServerMessageName,
		WorldId:  worldId,
		Position: position,
	}
}

type playerJoinedServerMessage struct {
	Name   serverMessageName
	Player dto.PlayerDto
}

func newPlayerJoinedServerMessage(player dto.PlayerDto) playerJoinedServerMessage {
	return playerJoinedServerMessage{
		Name:   playerJoinedServerMessageName,
		Player: player,
	}
}

type playerLeftServerMessage struct {
	Name     serverMessageName
	PlayerId uuid.UUID
}

func newPlayerLeftServerMessage(playerId uuid.UUID) playerLeftServerMessage {
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

type playerHeldItemChangedServerMessage struct {
	Name     serverMessageName
	PlayerId uuid.UUID
	ItemId   uuid.UUID
}

func newPlayerHeldItemChangedServerMessage(playerId uuid.UUID, itemId uuid.UUID) playerHeldItemChangedServerMessage {
	return playerHeldItemChangedServerMessage{
		Name:     playerHeldItemChangedServerMessageName,
		PlayerId: playerId,
		ItemId:   itemId,
	}
}

func newWorldServerMessageChannel(worldIdDto uuid.UUID) string {
	return fmt.Sprintf("WORLD_%s_CHANNEL", worldIdDto)
}
