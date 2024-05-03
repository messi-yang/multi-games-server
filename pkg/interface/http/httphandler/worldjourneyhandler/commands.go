package worldjourneyhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type commandName string

const (
	changePlayerActionCommandName   commandName = "CHANGE_PLAYER_ACTION"
	sendPlayerIntoPortalCommandName commandName = "SEND_PLAYER_INTO_PORTAL"
	changePlayerHeldItemCommandName commandName = "CHANGE_PLAYER_HELD_ITEM"
	createStaticUnitCommandName     commandName = "CREATE_STATIC_UNIT"
	removeStaticUnitCommandName     commandName = "REMOVE_STATIC_UNIT"
	createFenceUnitCommandName      commandName = "CREATE_FENCE_UNIT"
	removeFenceUnitCommandName      commandName = "REMOVE_FENCE_UNIT"
	createPortalUnitCommandName     commandName = "CREATE_PORTAL_UNIT"
	removePortalUnitCommandName     commandName = "REMOVE_PORTAL_UNIT"
	createLinkUnitCommandName       commandName = "CREATE_LINK_UNIT"
	removeLinkUnitCommandName       commandName = "REMOVE_LINK_UNIT"
	createEmbedUnitCommandName      commandName = "CREATE_EMBED_UNIT"
	removeEmbedUnitCommandName      commandName = "REMOVE_EMBED_UNIT"
	rotateUnitCommandName           commandName = "ROTATE_UNIT"
)

type command struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
}

type changePlayerActionCommand struct {
	Id        uuid.UUID           `json:"id"`
	Timestamp int64               `json:"timestamp"`
	Name      commandName         `json:"name"`
	PlayerId  uuid.UUID           `json:"playerId"`
	Action    dto.PlayerActionDto `json:"action"`
}

type sendPlayerIntoPortalCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	PlayerId  uuid.UUID   `json:"playerId"`
	UnitId    uuid.UUID   `json:"unitId"`
}

type changePlayerHeldItemCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	PlayerId  uuid.UUID   `json:"playerId"`
	ItemId    uuid.UUID   `json:"itemId"`
}

type createStaticUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	UnitId    uuid.UUID       `json:"unitId"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type removeStaticUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	UnitId    uuid.UUID   `json:"unitId"`
}

type createFenceUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	UnitId    uuid.UUID       `json:"unitId"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type removeFenceUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	UniId     uuid.UUID   `json:"unitId"`
}

type createPortalUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	UnitId    uuid.UUID       `json:"unitId"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
}

type removePortalUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	UnitId    uuid.UUID   `json:"unitId"`
}

type createLinkUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	UnitId    uuid.UUID       `json:"unitId"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
	Label     *string         `json:"label"`
	Url       string          `json:"url"`
}

type removeLinkUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	UnitId    uuid.UUID   `json:"unitId"`
}

type createEmbedUnitCommand struct {
	Id        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
	Name      commandName     `json:"name"`
	UnitId    uuid.UUID       `json:"unitId"`
	ItemId    uuid.UUID       `json:"itemId"`
	Position  dto.PositionDto `json:"position"`
	Direction int8            `json:"direction"`
	Label     *string         `json:"label"`
	EmbedCode string          `json:"embedCode"`
}

type removeEmbedUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	UnitId    uuid.UUID   `json:"unitId"`
}

type rotateUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	UnitId    uuid.UUID   `json:"unitId"`
}
