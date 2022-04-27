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

func extractWatchUnitsActionFromMessage(msg []byte) (*watchUnitsAction, error) {
	var action watchUnitsAction
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}

	return &action, nil
}
