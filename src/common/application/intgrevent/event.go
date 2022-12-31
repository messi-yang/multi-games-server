package intgrevent

import "encoding/json"

type Event struct {
	Name string `json:"name"`
}

func Parse(bytes []byte) (Event, error) {
	var event Event
	err := json.Unmarshal(bytes, &event)
	if err != nil {
		return Event{}, err
	}

	return event, nil
}
