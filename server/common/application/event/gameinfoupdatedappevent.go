package event

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/adapter/common/dto/jsondto"
)

type GameInfoUpdatedAppEvent struct {
	LiveGameId string                         `json:"liveGameId"`
	PlayerId   string                         `json:"playerId"`
	Dimension  commonjsondto.DimensionJsonDto `json:"dimension"`
}

func NewGameInfoUpdatedAppEvent(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, dimension commonmodel.Dimension) *GameInfoUpdatedAppEvent {
	return &GameInfoUpdatedAppEvent{
		LiveGameId: liveGameId.ToString(),
		PlayerId:   playerId.ToString(),
		Dimension:  commonjsondto.NewDimensionJsonDto(dimension),
	}
}

func DeserializeGameInfoUpdatedAppEvent(message []byte) GameInfoUpdatedAppEvent {
	var event GameInfoUpdatedAppEvent
	json.Unmarshal(message, &event)
	return event
}

func NewGameInfoUpdatedAppEventChannel(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) string {
	return fmt.Sprintf("game-info-updated-live-game-id-%s-player-id-%s", liveGameId.ToString(), playerId.ToString())
}

func (event *GameInfoUpdatedAppEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *GameInfoUpdatedAppEvent) GetDimension() (commonmodel.Dimension, error) {
	dimension, err := event.Dimension.ToValueObject()
	if err != nil {
		return commonmodel.Dimension{}, nil
	}
	return dimension, nil
}
