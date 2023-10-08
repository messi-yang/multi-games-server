package worldjourneyhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/viewmodel"
	"github.com/google/uuid"
)

type eventName string

const (
	worldEnteredEventName     eventName = "WORLD_ENTERED"
	commandSucceededEventName eventName = "COMMAND_SUCCEEDED"
	erroredEventName          eventName = "ERRORED"
)

type event struct {
	Name eventName `json:"name"`
}

type worldEnteredEvent struct {
	Name       eventName                `json:"name"`
	World      viewmodel.WorldViewModel `json:"world"`
	Units      []dto.UnitDto            `json:"units"`
	MyPlayerId uuid.UUID                `json:"myPlayerId"`
	Players    []dto.PlayerDto          `json:"players"`
}

type commandSucceededEvent struct {
	Name    eventName `json:"name"`
	Command any       `json:"command"`
}

type erroredEvent struct {
	Name    eventName `json:"name"`
	Message string    `json:"message"`
}
