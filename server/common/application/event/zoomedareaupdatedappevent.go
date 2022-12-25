package event

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type ZoomedAreaUpdatedAppEvent struct {
	LiveGameId string                         `json:"liveGameId"`
	PlayerId   string                         `json:"playerId"`
	Area       commonjsondto.AreaJsonDto      `json:"area"`
	UnitBlock  commonjsondto.UnitBlockJsonDto `json:"unitBlock"`
}

func NewZoomedAreaUpdatedAppEvent(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area, unitBlock commonmodel.UnitBlock) *ZoomedAreaUpdatedAppEvent {
	return &ZoomedAreaUpdatedAppEvent{
		LiveGameId: liveGameId.ToString(),
		PlayerId:   playerId.ToString(),
		Area:       commonjsondto.NewAreaJsonDto(area),
		UnitBlock:  commonjsondto.NewUnitBlockJsonDto(unitBlock),
	}
}

func DeserializeZoomedAreaUpdatedAppEvent(message []byte) ZoomedAreaUpdatedAppEvent {
	var event ZoomedAreaUpdatedAppEvent
	json.Unmarshal(message, &event)
	return event
}

func NewZoomedAreaUpdatedAppEventChannel(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) string {
	return fmt.Sprintf("zoomed-area-updated-live-game-id-%s-player-id-%s", liveGameId.ToString(), playerId.ToString())
}

func (event *ZoomedAreaUpdatedAppEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *ZoomedAreaUpdatedAppEvent) GetArea() (commonmodel.Area, error) {
	area, err := event.Area.ToValueObject()
	if err != nil {
		return commonmodel.Area{}, err
	}
	return area, nil
}

func (event *ZoomedAreaUpdatedAppEvent) GetUnitBlock() (commonmodel.UnitBlock, error) {
	unitBlock, err := event.UnitBlock.ToValueObject()
	if err != nil {
		return commonmodel.UnitBlock{}, err
	}
	return unitBlock, nil
}
