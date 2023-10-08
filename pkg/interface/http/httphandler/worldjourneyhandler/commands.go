package worldjourneyhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type commandName string

const (
	pingCommandName                 commandName = "PING"
	displayerErrorCommandName       commandName = "DISPLAY_ERROR"
	enterWorldCommandName           commandName = "ENTER_WORLD"
	addPlayerCommandName            commandName = "ADD_PLAYER"
	movePlayerCommandName           commandName = "MOVE_PLAYER"
	changePlayerHeldItemCommandName commandName = "CHANGE_PLAYER_HELD_ITEM"
	removePlayerCommandName         commandName = "REMOVE_PLAYER"
	createStaticUnitCommandName     commandName = "CREATE_STATIC_UNIT"
	createPortalUnitCommandName     commandName = "CREATE_PORTAL_UNIT"
	rotateUnitCommandName           commandName = "ROTATE_UNIT"
	removeUnitCommandName           commandName = "REMOVE_UNIT"
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

type movePlayerCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	PlayerId  uuid.UUID       `json:"playerId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
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

type createPortalUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type rotateUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	Position  dto.PositionDto `json:"position"`
}

type removeUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	Position  dto.PositionDto `json:"position"`
}
