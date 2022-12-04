package event

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type ReviveUnitsRequestedApplicationEvent struct {
	LiveGameId  commonjsondto.LiveGameIdJsonDto `json:"liveGameId"`
	Coordinates []jsondto.CoordinateJsonDto     `json:"coordinates"`
}

func NewReviveUnitsRequestedApplicationEvent(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) *ReviveUnitsRequestedApplicationEvent {
	return &ReviveUnitsRequestedApplicationEvent{
		LiveGameId:  commonjsondto.NewLiveGameIdJsonDto(liveGameId),
		Coordinates: commonjsondto.NewCoordinateJsonDtos(coordinates),
	}
}

func DeserializeReviveUnitsRequestedApplicationEvent(message []byte) ReviveUnitsRequestedApplicationEvent {
	var event ReviveUnitsRequestedApplicationEvent
	json.Unmarshal(message, &event)
	return event
}

func NewReviveUnitsRequestedApplicationEventChannel() string {
	return "revive-units-requested"
}

func (event *ReviveUnitsRequestedApplicationEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *ReviveUnitsRequestedApplicationEvent) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := event.LiveGameId.ToValueObject()
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *ReviveUnitsRequestedApplicationEvent) GetCoordinates() ([]gamecommonmodel.Coordinate, error) {
	coordinates, err := jsondto.ParseCoordinateJsonDtos(event.Coordinates)
	if err != nil {
		return nil, err
	}
	return coordinates, nil
}
