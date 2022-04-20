package gamesocketcontroller

import "encoding/json"

func getEventTypeFromMessage(msg []byte) (*eventType, error) {
	var newEvent event
	err := json.Unmarshal(msg, &newEvent)
	if err != nil {
		return nil, err
	}

	return &newEvent.Type, nil
}
