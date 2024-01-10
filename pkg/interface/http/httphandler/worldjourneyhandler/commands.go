package worldjourneyhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type commandName string

const (
	pingCommandName                 commandName = "PING"
	addPlayerCommandName            commandName = "ADD_PLAYER"
	changePlayerActionCommandName   commandName = "CHANGE_PLAYER_ACTION"
	sendPlayerIntoPortalCommandName commandName = "SEND_PLAYER_INTO_PORTAL"
	changePlayerHeldItemCommandName commandName = "CHANGE_PLAYER_HELD_ITEM"
	removePlayerCommandName         commandName = "REMOVE_PLAYER"
	createStaticUnitCommandName     commandName = "CREATE_STATIC_UNIT"
	removeStaticUnitCommandName     commandName = "REMOVE_STATIC_UNIT"
	createFenceUnitCommandName      commandName = "CREATE_FENCE_UNIT"
	removeFenceUnitCommandName      commandName = "REMOVE_FENCE_UNIT"
	createPortalUnitCommandName     commandName = "CREATE_PORTAL_UNIT"
	removePortalUnitCommandName     commandName = "REMOVE_PORTAL_UNIT"
	rotateUnitCommandName           commandName = "ROTATE_UNIT"
)

type command struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
}

type addPlayerCommand struct {
	Id        uuid.UUID     `json:"id"`
	Timestamp int64         `json:"timestamp"`
	Name      commandName   `json:"name"`
	Player    dto.PlayerDto `json:"player"`
}

type changePlayerActionCommand struct {
	Id        uuid.UUID           `json:"id"`
	Timestamp int64               `json:"timestamp"`
	Name      commandName         `json:"name"`
	PlayerId  uuid.UUID           `json:"playerId"`
	Action    dto.PlayerActionDto `json:"action"`
}

type sendPlayerIntoPortalCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	PlayerId  uuid.UUID       `json:"playerId"`
	Position  dto.PositionDto `json:"position"`
}

type changePlayerHeldItemCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	PlayerId  uuid.UUID   `json:"playerId"`
	ItemId    uuid.UUID   `json:"itemId"`
}
type removePlayerCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	PlayerId  uuid.UUID   `json:"playerId"`
}

type createStaticUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type removeStaticUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	Position  dto.PositionDto `json:"position"`
}

type createFenceUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type removeFenceUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	Position  dto.PositionDto `json:"position"`
}

type createPortalUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type removePortalUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	Position  dto.PositionDto `json:"position"`
}

type rotateUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	Position  dto.PositionDto `json:"position"`
}
