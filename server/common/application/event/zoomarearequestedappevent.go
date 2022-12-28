package event

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/adapter/common/dto/jsondto"
)

type ZoomAreaRequestedAppEvent struct {
	LiveGameId string                    `json:"liveGameId"`
	PlayerId   string                    `json:"playerId"`
	Area       commonjsondto.AreaJsonDto `json:"area"`
}

func NewZoomAreaRequestedAppEvent(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area) *ZoomAreaRequestedAppEvent {
	return &ZoomAreaRequestedAppEvent{
		LiveGameId: liveGameId.ToString(),
		PlayerId:   playerId.ToString(),
		Area:       commonjsondto.NewAreaJsonDto(area),
	}
}

func DeserializeZoomAreaRequestedAppEvent(message []byte) ZoomAreaRequestedAppEvent {
	var event ZoomAreaRequestedAppEvent
	json.Unmarshal(message, &event)
	return event
}

func NewZoomAreaRequestedAppEventChannel() string {
	return "zoom-area-requested"
}

func (event *ZoomAreaRequestedAppEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *ZoomAreaRequestedAppEvent) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := livegamemodel.NewLiveGameId(event.LiveGameId)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *ZoomAreaRequestedAppEvent) GetPlayerId() (commonmodel.PlayerId, error) {
	playerId, err := commonmodel.NewPlayerId(event.PlayerId)
	if err != nil {
		return commonmodel.PlayerId{}, err
	}
	return playerId, nil
}

func (event *ZoomAreaRequestedAppEvent) GetArea() (commonmodel.Area, error) {
	area, err := event.Area.ToValueObject()
	if err != nil {
		return commonmodel.Area{}, err
	}
	return area, nil
}
