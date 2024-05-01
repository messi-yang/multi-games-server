package worldjourneyhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/viewmodel"
	"github.com/google/uuid"
)

type serverEventName string

const (
	worldEnteredServerEventName     serverEventName = "WORLD_ENTERED"
	playerJoinedServerEventName     serverEventName = "PLAYER_JOINED"
	playerLeftServerEventName       serverEventName = "PLAYER_LEFT"
	commandSucceededServerEventName serverEventName = "COMMAND_SUCCEEDED"
	commandFailedServerEventName    serverEventName = "COMMAND_FAILED"
	erroredServerEventName          serverEventName = "ERRORED"
)

type worldEnteredServerEvent struct {
	Name       serverEventName          `json:"name"`
	World      viewmodel.WorldViewModel `json:"world"`
	Units      []dto.UnitDto            `json:"units"`
	MyPlayerId uuid.UUID                `json:"myPlayerId"`
	Players    []dto.PlayerDto          `json:"players"`
}

type playerJoinedServerEvent struct {
	Name   serverEventName `json:"name"`
	Player dto.PlayerDto
}

type playerLeftServerEvent struct {
	Name     serverEventName `json:"name"`
	PlayerId uuid.UUID
}

type commandSucceededServerEvent struct {
	Name    serverEventName `json:"name"`
	Command any             `json:"command"`
}

type commandFailedServerEvent struct {
	Name         serverEventName `json:"name"`
	CommandId    uuid.UUID       `json:"commandId"`
	ErrorMessage string          `json:"errorMessage"`
}

type erroredServerEvent struct {
	Name    serverEventName `json:"name"`
	Message string          `json:"message"`
}
