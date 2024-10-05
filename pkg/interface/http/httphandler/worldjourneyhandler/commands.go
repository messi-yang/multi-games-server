package worldjourneyhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
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
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Payload   struct {
		PlayerId     uuid.UUID           `json:"playerId"`
		PlayerAction dto.PlayerActionDto `json:"action"`
	} `json:"payload"`
}

type sendPlayerIntoPortalCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Payload   struct {
		PlayerId uuid.UUID `json:"playerId"`
		UnitId   uuid.UUID `json:"unitId"`
	} `json:"payload"`
}

type changePlayerHeldItemCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Payload   struct {
		PlayerId uuid.UUID `json:"playerId"`
		ItemId   uuid.UUID `json:"itemId"`
	} `json:"payload"`
}

type createStaticUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Payload   struct {
		UnitId        uuid.UUID       `json:"unitId"`
		ItemId        uuid.UUID       `json:"itemId"`
		UnitPosition  dto.PositionDto `json:"unitPosition"`
		UnitDirection int8            `json:"unitDirection"`
	} `json:"payload"`
}

type removeStaticUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Paylaod   struct {
		UnitId uuid.UUID `json:"unitId"`
	} `json:"payload"`
}

type createFenceUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Paylaod   struct {
		UnitId        uuid.UUID       `json:"unitId"`
		ItemId        uuid.UUID       `json:"itemId"`
		UnitPosition  dto.PositionDto `json:"unitPosition"`
		UnitDirection int8            `json:"unitDirection"`
	} `json:"payload"`
}

type removeFenceUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Paylaod   struct {
		UniId uuid.UUID `json:"unitId"`
	} `json:"payload"`
}

type createPortalUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Paylaod   struct {
		UnitId        uuid.UUID       `json:"unitId"`
		ItemId        uuid.UUID       `json:"itemId"`
		UnitPosition  dto.PositionDto `json:"unitPosition"`
		UnitDirection int8            `json:"unitDirection"`
	} `json:"payload"`
}

type removePortalUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Paylaod   struct {
		UnitId uuid.UUID `json:"unitId"`
	} `json:"payload"`
}

type createLinkUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Paylaod   struct {
		UnitId        uuid.UUID       `json:"unitId"`
		ItemId        uuid.UUID       `json:"itemId"`
		UnitPosition  dto.PositionDto `json:"unitPosition"`
		UnitDirection int8            `json:"unitDirection"`
		UnitLabel     *string         `json:"unitLabel"`
		UnitUrl       string          `json:"unitUrl"`
	} `json:"payload"`
}

type removeLinkUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Paylaod   struct {
		UnitId uuid.UUID `json:"unitId"`
	} `json:"payload"`
}

type createEmbedUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Paylaod   struct {
		UnitId        uuid.UUID       `json:"unitId"`
		ItemId        uuid.UUID       `json:"itemId"`
		UnitPosition  dto.PositionDto `json:"unitPosition"`
		UnitDirection int8            `json:"unitDirection"`
		UnitLabel     *string         `json:"unitLabel"`
		UnitEmbedCode string          `json:"unitEmbedCode"`
	} `json:"payload"`
}

type removeEmbedUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Paylaod   struct {
		UnitId uuid.UUID `json:"unitId"`
	} `json:"payload"`
}

type rotateUnitCommand struct {
	Id        uuid.UUID   `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Name      commandName `json:"name"`
	Paylaod   struct {
		UnitId uuid.UUID `json:"unitId"`
	} `json:"payload"`
}
