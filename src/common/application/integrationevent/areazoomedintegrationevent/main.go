package areazoomedintegrationevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/viewmodel/areaviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/viewmodel/unitblockviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type Event struct {
	Name       string                                `json:"name"`
	LiveGameId string                                `json:"liveGameId"`
	PlayerId   string                                `json:"playerId"`
	Area       areaviewmodel.AreaViewModel           `json:"area"`
	UnitBlock  unitblockviewmodel.UnitBlockViewModel `json:"unitBlock"`
}

var EVENT_NAME = "AREA_ZOOMED"

func New(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area, unitBlock commonmodel.UnitBlock) *Event {
	return &Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId.ToString(),
		PlayerId:   playerId.ToString(),
		Area:       areaviewmodel.New(area),
		UnitBlock:  unitblockviewmodel.New(unitBlock),
	}
}

func Deserialize(message []byte) Event {
	var event Event
	json.Unmarshal(message, &event)
	return event
}

func (event *Event) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *Event) GetArea() (commonmodel.Area, error) {
	area, err := event.Area.ToValueObject()
	if err != nil {
		return commonmodel.Area{}, err
	}
	return area, nil
}

func (event *Event) GetUnitBlock() (commonmodel.UnitBlock, error) {
	unitBlock, err := event.UnitBlock.ToValueObject()
	if err != nil {
		return commonmodel.UnitBlock{}, err
	}
	return unitBlock, nil
}
