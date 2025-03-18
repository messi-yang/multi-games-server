package commandmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/google/uuid"
)

type ChangePlayerActionCommandPayloadJson struct {
	PlayerId     uuid.UUID           `json:"playerId"`
	PlayerAction dto.PlayerActionDto `json:"action"`
}

type SendPlayerIntoPortalCommandPayloadJson struct {
	PlayerId uuid.UUID `json:"playerId"`
	UnitId   uuid.UUID `json:"unitId"`
}

type ChangePlayerHeldItemCommandPayloadJson struct {
	PlayerId uuid.UUID `json:"playerId"`
	ItemId   uuid.UUID `json:"itemId"`
}

type CreateStaticUnitCommandPayloadJson struct {
	UnitId        uuid.UUID       `json:"unitId"`
	ItemId        uuid.UUID       `json:"itemId"`
	UnitPosition  dto.PositionDto `json:"unitPosition"`
	UnitDirection int8            `json:"unitDirection"`
}

type RemoveStaticUnitCommandPayloadJson struct {
	UnitId uuid.UUID `json:"unitId"`
}

type CreateFenceUnitCommandPayloadJson struct {
	UnitId        uuid.UUID       `json:"unitId"`
	ItemId        uuid.UUID       `json:"itemId"`
	UnitPosition  dto.PositionDto `json:"unitPosition"`
	UnitDirection int8            `json:"unitDirection"`
}

type RemoveFenceUnitCommandPayloadJson struct {
	UnitId uuid.UUID `json:"unitId"`
}

type CreatePortalUnitCommandPayloadJson struct {
	UnitId        uuid.UUID       `json:"unitId"`
	ItemId        uuid.UUID       `json:"itemId"`
	UnitPosition  dto.PositionDto `json:"unitPosition"`
	UnitDirection int8            `json:"unitDirection"`
}

type RemovePortalUnitCommandPayloadJson struct {
	UnitId uuid.UUID `json:"unitId"`
}

type CreateLinkUnitCommandPayloadJson struct {
	UnitId        uuid.UUID       `json:"unitId"`
	ItemId        uuid.UUID       `json:"itemId"`
	UnitPosition  dto.PositionDto `json:"unitPosition"`
	UnitDirection int8            `json:"unitDirection"`
	UnitLabel     *string         `json:"unitLabel"`
	UnitUrl       string          `json:"unitUrl"`
}

type RemoveLinkUnitCommandPayloadJson struct {
	UnitId uuid.UUID `json:"unitId"`
}

type CreateEmbedUnitCommandPayloadJson struct {
	UnitId        uuid.UUID       `json:"unitId"`
	ItemId        uuid.UUID       `json:"itemId"`
	UnitPosition  dto.PositionDto `json:"unitPosition"`
	UnitDirection int8            `json:"unitDirection"`
	UnitLabel     *string         `json:"unitLabel"`
	UnitEmbedCode string          `json:"unitEmbedCode"`
}

type RemoveEmbedUnitCommandPayloadJson struct {
	UnitId uuid.UUID `json:"unitId"`
}

type CreateColorUnitCommandPayloadJson struct {
	UnitId        uuid.UUID       `json:"unitId"`
	ItemId        uuid.UUID       `json:"itemId"`
	UnitPosition  dto.PositionDto `json:"unitPosition"`
	UnitDirection int8            `json:"unitDirection"`
	UnitLabel     *string         `json:"unitLabel"`
	UnitColor     string          `json:"unitColor"`
}

type RemoveColorUnitCommandPayloadJson struct {
	UnitId uuid.UUID `json:"unitId"`
}

type CreateSignUnitCommandPayloadJson struct {
	UnitId        uuid.UUID       `json:"unitId"`
	ItemId        uuid.UUID       `json:"itemId"`
	UnitPosition  dto.PositionDto `json:"unitPosition"`
	UnitDirection int8            `json:"unitDirection"`
	UnitLabel     string          `json:"unitLabel"`
}

type RemoveSignUnitCommandPayloadJson struct {
	UnitId uuid.UUID `json:"unitId"`
}

type RotateUnitCommandPayloadJson struct {
	UnitId uuid.UUID `json:"unitId"`
}
