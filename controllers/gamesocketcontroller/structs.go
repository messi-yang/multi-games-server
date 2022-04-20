package gamesocketcontroller

type eventType string

const (
	watchGameBlock eventType = "WATCH_GAME_BLOCK"
)

type event struct {
	Type eventType `json:"action"`
}

type watchGameBlockEvent struct {
	Type    eventType `json:"action"`
	Payload struct {
		FromX int `json:"fromX"`
		FromY int `json:"fromY"`
		ToX   int `json:"toX"`
		ToY   int `json:"toY"`
	} `json:"payload"`
}
