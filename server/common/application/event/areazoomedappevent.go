package event

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/adapter/common/dto/jsondto"
)

type AreaZoomedAppEvent struct {
	LiveGameId string                         `json:"liveGameId"`
	PlayerId   string                         `json:"playerId"`
	Area       commonjsondto.AreaJsonDto      `json:"area"`
	UnitBlock  commonjsondto.UnitBlockJsonDto `json:"unitBlock"`
}

func NewAreaZoomedAppEvent(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area, unitBlock commonmodel.UnitBlock) *AreaZoomedAppEvent {
	return &AreaZoomedAppEvent{
		LiveGameId: liveGameId.ToString(),
		PlayerId:   playerId.ToString(),
		Area:       commonjsondto.NewAreaJsonDto(area),
		UnitBlock:  commonjsondto.NewUnitBlockJsonDto(unitBlock),
	}
}

func DeserializeAreaZoomedAppEvent(message []byte) AreaZoomedAppEvent {
	var event AreaZoomedAppEvent
	json.Unmarshal(message, &event)
	return event
}

func NewAreaZoomedAppEventChannel(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", liveGameId.ToString(), playerId.ToString())
}

func (event *AreaZoomedAppEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *AreaZoomedAppEvent) GetArea() (commonmodel.Area, error) {
	area, err := event.Area.ToValueObject()
	if err != nil {
		return commonmodel.Area{}, err
	}
	return area, nil
}

func (event *AreaZoomedAppEvent) GetUnitBlock() (commonmodel.UnitBlock, error) {
	unitBlock, err := event.UnitBlock.ToValueObject()
	if err != nil {
		return commonmodel.UnitBlock{}, err
	}
	return unitBlock, nil
}
