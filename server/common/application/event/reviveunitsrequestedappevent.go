package event

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type ReviveUnitsRequestedAppEvent struct {
	LiveGameId  string                      `json:"liveGameId"`
	Coordinates []jsondto.CoordinateJsonDto `json:"coordinates"`
}

func NewReviveUnitsRequestedAppEvent(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) *ReviveUnitsRequestedAppEvent {
	return &ReviveUnitsRequestedAppEvent{
		LiveGameId:  liveGameId.ToString(),
		Coordinates: commonjsondto.NewCoordinateJsonDtos(coordinates),
	}
}

func DeserializeReviveUnitsRequestedAppEvent(message []byte) ReviveUnitsRequestedAppEvent {
	var event ReviveUnitsRequestedAppEvent
	json.Unmarshal(message, &event)
	return event
}

func NewReviveUnitsRequestedAppEventChannel() string {
	return "revive-units-requested"
}

func (event *ReviveUnitsRequestedAppEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *ReviveUnitsRequestedAppEvent) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := livegamemodel.NewLiveGameId(event.LiveGameId)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *ReviveUnitsRequestedAppEvent) GetCoordinates() ([]gamecommonmodel.Coordinate, error) {
	coordinates, err := jsondto.ParseCoordinateJsonDtos(event.Coordinates)
	if err != nil {
		return nil, err
	}
	return coordinates, nil
}
