package gameroomsockethandler

import (
	"encoding/json"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator/presenter/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator/presenter/dto/coordinatedto"
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

func extractInformationFromZoomAreaAction(msg []byte) (valueobject.Area, error) {
	var action zoomAreaAction
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return valueobject.Area{}, err
	}

	area, err := areadto.FromDto(action.Payload.Area)
	if err != nil {
		return valueobject.Area{}, err
	}

	return area, nil
}

type reviveUnitsActionPayload struct {
	Coordinates []coordinatedto.Dto `json:"coordinates"`
	ActionedAt  time.Time           `json:"actionedAt"`
}
type reviveUnitsAction struct {
	Type    actionType               `json:"type"`
	Payload reviveUnitsActionPayload `json:"payload"`
}

func extractInformationFromReviveUnitsAction(msg []byte) ([]valueobject.Coordinate, error) {
	var action reviveUnitsAction
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}

	coordinates, err := coordinatedto.FromDtoList(action.Payload.Coordinates)
	if err != nil {
		return nil, err
	}

	return coordinates, nil
}
