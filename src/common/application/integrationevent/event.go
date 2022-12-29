package integrationevent

import "encoding/json"

type Event struct {
	Name string `json:"name"`
}

func New(bytes []byte) Event {
	var event *Event
	json.Unmarshal(bytes, &event)

	return *event
}
