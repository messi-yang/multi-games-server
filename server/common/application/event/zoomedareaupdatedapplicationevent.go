package event

import (
	"fmt"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type ZoomedAreaUpdatedApplicationEvent struct {
	LiveGameId commonjsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   commonjsondto.PlayerIdJsonDto   `json:"playerId"`
	Area       commonjsondto.AreaJsonDto       `json:"area"`
	UnitBlock  commonjsondto.UnitBlockJsonDto  `json:"unitBlock"`
}

func NewZoomedAreaUpdatedApplicationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area, unitBlock gamecommonmodel.UnitBlock) ZoomedAreaUpdatedApplicationEvent {
	return ZoomedAreaUpdatedApplicationEvent{
		LiveGameId: commonjsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   commonjsondto.NewPlayerIdJsonDto(playerId),
		Area:       commonjsondto.NewAreaJsonDto(area),
		UnitBlock:  commonjsondto.NewUnitBlockJsonDto(unitBlock),
	}
}

func NewZoomedAreaUpdatedApplicationEventChannel(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", liveGameId.GetId().String(), playerId.GetId().String())
}

func (event *ZoomedAreaUpdatedApplicationEvent) GetArea() (gamecommonmodel.Area, error) {
	area, err := event.Area.ToValueObject()
	if err != nil {
		return gamecommonmodel.Area{}, err
	}
	return area, nil
}

func (event *ZoomedAreaUpdatedApplicationEvent) GetUnitBlock() (gamecommonmodel.UnitBlock, error) {
	unitBlock, err := event.UnitBlock.ToValueObject()
	if err != nil {
		return gamecommonmodel.UnitBlock{}, err
	}
	return unitBlock, nil
}
