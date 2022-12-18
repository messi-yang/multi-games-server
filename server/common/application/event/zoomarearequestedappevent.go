package event

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type ZoomAreaRequestedAppEvent struct {
	LiveGameId string                    `json:"liveGameId"`
	PlayerId   string                    `json:"playerId"`
	Area       commonjsondto.AreaJsonDto `json:"area"`
}

func NewZoomAreaRequestedAppEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) *ZoomAreaRequestedAppEvent {
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

func (event *ZoomAreaRequestedAppEvent) GetPlayerId() (gamecommonmodel.PlayerId, error) {
	playerId, err := gamecommonmodel.NewPlayerId(event.PlayerId)
	if err != nil {
		return gamecommonmodel.PlayerId{}, err
	}
	return playerId, nil
}

func (event *ZoomAreaRequestedAppEvent) GetArea() (gamecommonmodel.Area, error) {
	area, err := event.Area.ToValueObject()
	if err != nil {
		return gamecommonmodel.Area{}, err
	}
	return area, nil
}
