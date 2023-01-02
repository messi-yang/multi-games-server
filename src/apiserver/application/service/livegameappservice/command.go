package livegameappservice

import (
	"encoding/json"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/locationviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/maprangeviewmodel"
)

type CommanType string

const (
	NilCommandType         CommanType = ""
	ZoomMapRangeCommanType CommanType = "ZOOM_MAP_RANGE"
	BuildItemCommanType    CommanType = "BUILD_ITEM"
	DestroyItemCommanType  CommanType = "DESTROY_ITEM"
)

type GenericCommand struct {
	Type CommanType `json:"type"`
}

func ParseCommandType(message []byte) (CommanType, error) {
	var command GenericCommand
	err := json.Unmarshal(message, &command)
	if err != nil {
		return NilCommandType, err
	}

	return command.Type, nil
}

func ParseCommand[T any](message []byte) (T, error) {
	var command T
	err := json.Unmarshal(message, &command)
	if err != nil {
		return command, err
	}

	return command, nil
}

type BuildItemCommand struct {
	Type    CommanType `json:"type"`
	Payload struct {
		Location   locationviewmodel.ViewModel `json:"location"`
		ItemId     string                      `json:"itemId"`
		ActionedAt time.Time                   `json:"actionedAt"`
	} `json:"payload"`
}

type DestroyItemCommand struct {
	Type    CommanType `json:"type"`
	Payload struct {
		Location   locationviewmodel.ViewModel `json:"location"`
		ActionedAt time.Time                   `json:"actionedAt"`
	} `json:"payload"`
}

type ZoomMapRangeCommand struct {
	Type    CommanType `json:"type"`
	Payload struct {
		MapRange   maprangeviewmodel.ViewModel `json:"mapRange"`
		ActionedAt time.Time                   `json:"actionedAt"`
	} `json:"payload"`
}
