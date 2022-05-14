package gamesocketcontroller

type actionType string

const (
	watchAreaActionType   actionType = "WATCH_AREA"
	reviveUnitsActionType actionType = "REVIVE_UNITS"
)

type action struct {
	Type actionType `json:"type"`
}

type watchAreaActionPayload struct {
	Area AreaDTO `json:"area"`
}
type watchAreaAction struct {
	Type    actionType             `json:"type"`
	Payload watchAreaActionPayload `json:"payload"`
}

type reviveUnitsActionPayload struct {
	Coordinates []CoordinateDTO `json:"coordinates"`
}
type reviveUnitsAction struct {
	Type    actionType               `json:"type"`
	Payload reviveUnitsActionPayload `json:"payload"`
}
