package event

import (
	"encoding/json"
	"fmt"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type GameInfoUpdatedApplicationEvent struct {
	LiveGameId commonjsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   commonjsondto.PlayerIdJsonDto   `json:"playerId"`
	Dimension  commonjsondto.DimensionJsonDto  `json:"dimension"`
}

func NewGameInfoUpdatedApplicationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, dimension gamecommonmodel.Dimension) *GameInfoUpdatedApplicationEvent {
	return &GameInfoUpdatedApplicationEvent{
		LiveGameId: commonjsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   commonjsondto.NewPlayerIdJsonDto(playerId),
		Dimension:  commonjsondto.NewDimensionJsonDto(dimension),
	}
}

func DeserializeGameInfoUpdatedApplicationEvent(message []byte) GameInfoUpdatedApplicationEvent {
	var event GameInfoUpdatedApplicationEvent
	json.Unmarshal(message, &event)
	return event
}

func NewGameInfoUpdatedApplicationEventChannel(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("game-info-updated-live-game-id-%s-player-id-%s", liveGameId.GetId().String(), playerId.GetId().String())
}

func (event *GameInfoUpdatedApplicationEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *GameInfoUpdatedApplicationEvent) GetDimension() (gamecommonmodel.Dimension, error) {
	dimension, err := event.Dimension.ToValueObject()
	if err != nil {
		return gamecommonmodel.Dimension{}, nil
	}
	return dimension, nil
}
