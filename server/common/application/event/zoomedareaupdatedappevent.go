package event

import (
	"encoding/json"
	"fmt"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type ZoomedAreaUpdatedAppEvent struct {
	LiveGameId string                         `json:"liveGameId"`
	PlayerId   commonjsondto.PlayerIdJsonDto  `json:"playerId"`
	Area       commonjsondto.AreaJsonDto      `json:"area"`
	UnitBlock  commonjsondto.UnitBlockJsonDto `json:"unitBlock"`
}

func NewZoomedAreaUpdatedAppEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area, unitBlock gamecommonmodel.UnitBlock) *ZoomedAreaUpdatedAppEvent {
	return &ZoomedAreaUpdatedAppEvent{
		LiveGameId: liveGameId.ToString(),
		PlayerId:   commonjsondto.NewPlayerIdJsonDto(playerId),
		Area:       commonjsondto.NewAreaJsonDto(area),
		UnitBlock:  commonjsondto.NewUnitBlockJsonDto(unitBlock),
	}
}

func DeserializeZoomedAreaUpdatedAppEvent(message []byte) ZoomedAreaUpdatedAppEvent {
	var event ZoomedAreaUpdatedAppEvent
	json.Unmarshal(message, &event)
	return event
}

func NewZoomedAreaUpdatedAppEventChannel(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", liveGameId.ToString(), playerId.GetId().String())
}

func (event *ZoomedAreaUpdatedAppEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *ZoomedAreaUpdatedAppEvent) GetArea() (gamecommonmodel.Area, error) {
	area, err := event.Area.ToValueObject()
	if err != nil {
		return gamecommonmodel.Area{}, err
	}
	return area, nil
}

func (event *ZoomedAreaUpdatedAppEvent) GetUnitBlock() (gamecommonmodel.UnitBlock, error) {
	unitBlock, err := event.UnitBlock.ToValueObject()
	if err != nil {
		return gamecommonmodel.UnitBlock{}, err
	}
	return unitBlock, nil
}
