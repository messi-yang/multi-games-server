package gamesocketcontroller

import (
	"encoding/json"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/dto/coordinatedto"
)

type actionType string

const (
	zoomAreaActionType    actionType = "ZOOM_AREA"
	reviveUnitsActionType actionType = "REVIVE_UNITS"
)

type action struct {
	Type actionType `json:"type"`
}

func getActionTypeFromMessage(msg []byte) (*actionType, error) {
	var action action
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}

	return &action.Type, nil
}

type zoomAreaActionPayload struct {
	Area       areadto.Dto `json:"area"`
	ActionedAt time.Time   `json:"actionedAt"`
}
type zoomAreaAction struct {
	Type    actionType            `json:"type"`
	Payload zoomAreaActionPayload `json:"payload"`
}

func extractZoomAreaActionFromMessage(msg []byte) (*zoomAreaAction, error) {
	var action zoomAreaAction
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}

	return &action, nil
}

type reviveUnitsActionPayload struct {
	Coordinates []coordinatedto.Dto `json:"coordinates"`
	ActionedAt  time.Time           `json:"actionedAt"`
}
type reviveUnitsAction struct {
	Type    actionType               `json:"type"`
	Payload reviveUnitsActionPayload `json:"payload"`
}

func extractReviveUnitsActionFromMessage(msg []byte) (*reviveUnitsAction, error) {
	var action reviveUnitsAction
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}

	return &action, nil
}
