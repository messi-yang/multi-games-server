package gamesocketcontroller

import (
	"encoding/json"

	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/areadto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/unitspatterndto"
)

type actionType string

const (
	watchAreaActionType   actionType = "WATCH_AREA"
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

type watchAreaActionPayload struct {
	Area areadto.AreaDTO `json:"area"`
}
type watchAreaAction struct {
	Type    actionType             `json:"type"`
	Payload watchAreaActionPayload `json:"payload"`
}

func extractWatchAreaActionFromMessage(msg []byte) (*watchAreaAction, error) {
	var action watchAreaAction
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}

	return &action, nil
}

type reviveUnitsActionPayload struct {
	Coordinate coordinatedto.CoordinateDTO     `json:"coordinate"`
	Pattern    unitspatterndto.UnitsPatternDTO `json:"pattern"`
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
