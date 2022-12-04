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

func (event *RedisReviveUnitsRequestedEvent) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := event.LiveGameId.ToValueObject()
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *RedisReviveUnitsRequestedEvent) GetCoordinates() ([]gamecommonmodel.Coordinate, error) {
	coordinates, err := jsondto.ParseCoordinateJsonDtos(event.Coordinates)
	if err != nil {
		return nil, err
	}
	return coordinates, nil
}
