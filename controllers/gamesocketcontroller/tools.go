package gamesocketcontroller

import (
	"encoding/json"
	"math/rand"
)

func getEventTypeFromMessage(msg []byte) (*eventType, error) {
	var newEvent event
	err := json.Unmarshal(msg, &newEvent)
	if err != nil {
		return nil, err
	}

	return &newEvent.Type, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomHash(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
