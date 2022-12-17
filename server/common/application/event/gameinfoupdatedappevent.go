package event

import (
	"encoding/json"
	"fmt"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type GameInfoUpdatedAppEvent struct {
	LiveGameId commonjsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   commonjsondto.PlayerIdJsonDto   `json:"playerId"`
	Dimension  commonjsondto.DimensionJsonDto  `json:"dimension"`
}

func NewGameInfoUpdatedAppEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, dimension gamecommonmodel.Dimension) *GameInfoUpdatedAppEvent {
	return &GameInfoUpdatedAppEvent{
		LiveGameId: commonjsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   commonjsondto.NewPlayerIdJsonDto(playerId),
		Dimension:  commonjsondto.NewDimensionJsonDto(dimension),
	}
}

func DeserializeGameInfoUpdatedAppEvent(message []byte) GameInfoUpdatedAppEvent {
	var event GameInfoUpdatedAppEvent
	json.Unmarshal(message, &event)
	return event
}

func NewGameInfoUpdatedAppEventChannel(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("game-info-updated-live-game-id-%s-player-id-%s", liveGameId.GetId().String(), playerId.GetId().String())
}

func (event *GameInfoUpdatedAppEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *GameInfoUpdatedAppEvent) GetDimension() (gamecommonmodel.Dimension, error) {
	dimension, err := event.Dimension.ToValueObject()
	if err != nil {
		return gamecommonmodel.Dimension{}, nil
	}
	return dimension, nil
}
