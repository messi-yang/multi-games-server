package gamesocketcontroller

import (
	"encoding/json"
)

func getActionTypeFromMessage(msg []byte) (*actionType, error) {
	var action action
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}

	return &action.Type, nil
}

func extractWatchAreaActionFromMessage(msg []byte) (*watchAreaAction, error) {
	var action watchAreaAction
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}

	return &action, nil
}

func extractReviveUnitsActionFromMessage(msg []byte) (*reviveUnitsAction, error) {
	var action reviveUnitsAction
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}

	return &action, nil
}
