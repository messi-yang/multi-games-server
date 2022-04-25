package gamesocketcontroller

import (
	"encoding/json"
)

func getActionTypeFromMessage(msg []byte) (*actionType, error) {
	var newAction action
	err := json.Unmarshal(msg, &newAction)
	if err != nil {
		return nil, err
	}

	return &newAction.Type, nil
}
