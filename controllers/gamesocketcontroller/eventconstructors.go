package gamesocketcontroller

func constructErrorHappenedEvent(clientMessage string) errorHappenedEvent {
	return errorHappenedEvent{
		Type: errorHappenedEventType,
		Payload: struct {
			ClientMessage string `json:"clientMessage"`
		}{
			ClientMessage: clientMessage,
		},
	}
}
