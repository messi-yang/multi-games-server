package gameroomsockethandler

import (
	"encoding/json"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
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
	Area       dto.AreaDto `json:"area"`
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

	area, err := action.Payload.Area.ToArea()
	if err != nil {
		return valueobject.Area{}, err
	}

	return area, nil
}

type reviveUnitsActionPayload struct {
	Coordinates []dto.CoordinateDto `json:"coordinates"`
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

	coordinates, err := dto.ParseCoordinateDtos(action.Payload.Coordinates)
	if err != nil {
		return nil, err
	}

	return coordinates, nil
}
