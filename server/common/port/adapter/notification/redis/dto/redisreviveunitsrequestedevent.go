package dto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type RedisReviveUnitsRequestedEvent struct {
	LiveGameId  commonjsondto.LiveGameIdJsonDto `json:"liveGameId"`
	Coordinates []jsondto.CoordinateJsonDto     `json:"coordinates"`
}

func NewRedisReviveUnitsRequestedEvent(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) RedisReviveUnitsRequestedEvent {
	return RedisReviveUnitsRequestedEvent{
		LiveGameId:  commonjsondto.NewLiveGameIdJsonDto(liveGameId),
		Coordinates: commonjsondto.NewCoordinateJsonDtos(coordinates),
	}
}

func NewRedisReviveUnitsRequestedEventChannel() string {
	return "revive-units-requested"
}
