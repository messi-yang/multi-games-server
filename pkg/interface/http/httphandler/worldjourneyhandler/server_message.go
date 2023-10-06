package worldjourneyhandler

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type serverMessageName string

const (
	staticUnitCreatedServerMessageName     serverMessageName = "STATIC_UNIT_CREATED"
	portalUnitCreatedServerMessageName     serverMessageName = "PORTAL_UNIT_CREATED"
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

type staticUnitCreatedServerMessage struct {
	Name      serverMessageName
	ItemId    uuid.UUID
	Position  dto.PositionDto
	Direction int8
}

func newStaticUnitCreatedServerMessage(itemId uuid.UUID, position dto.PositionDto, direction int8) staticUnitCreatedServerMessage {
	return staticUnitCreatedServerMessage{
		Name:      staticUnitCreatedServerMessageName,
		ItemId:    itemId,
		Position:  position,
		Direction: direction,
	}
}

type portalUnitCreatedServerMessage struct {
	Name      serverMessageName
	ItemId    uuid.UUID
	Position  dto.PositionDto
	Direction int8
}

func newPortalUnitCreatedServerMessage(itemId uuid.UUID, position dto.PositionDto, direction int8) portalUnitCreatedServerMessage {
	return portalUnitCreatedServerMessage{
		Name:      portalUnitCreatedServerMessageName,
		ItemId:    itemId,
		Position:  position,
		Direction: direction,
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
	Name      serverMessageName
	PlayerId  uuid.UUID
	Position  dto.PositionDto
	Direction int8
}

func newPlayerMovedServerMessage(playerId uuid.UUID, position dto.PositionDto, direction int8) playerMovedServerMessage {
	return playerMovedServerMessage{
		Name:      playerMovedServerMessageName,
		PlayerId:  playerId,
		Position:  position,
		Direction: direction,
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
