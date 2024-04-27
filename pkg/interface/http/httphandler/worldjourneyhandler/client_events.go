package worldjourneyhandler

type clientEventName string

const (
	pingClientEventName             clientEventName = "PING"
	commandRequestedClientEventName clientEventName = "COMMAND_REQUESTED"
)

type clientEvent struct {
	Name clientEventName `json:"name"`
}

type commandRequestedClientEvent[T any] struct {
	Name    clientEventName `json:"name"`
	Command T               `json:"command"`
}
