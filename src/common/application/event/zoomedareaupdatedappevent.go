package event

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/viewmodel/areaviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/viewmodel/unitblockviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type ZoomedAreaUpdatedAppEvent struct {
	LiveGameId string                                `json:"liveGameId"`
	PlayerId   string                                `json:"playerId"`
	Area       areaviewmodel.AreaViewModel           `json:"area"`
	UnitBlock  unitblockviewmodel.UnitBlockViewModel `json:"unitBlock"`
}

func NewZoomedAreaUpdatedAppEvent(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area, unitBlock commonmodel.UnitBlock) *ZoomedAreaUpdatedAppEvent {
	return &ZoomedAreaUpdatedAppEvent{
		LiveGameId: liveGameId.ToString(),
		PlayerId:   playerId.ToString(),
		Area:       areaviewmodel.New(area),
		UnitBlock:  unitblockviewmodel.New(unitBlock),
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
