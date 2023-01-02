package livegameappservice

import (
	"encoding/json"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/areaviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/locationviewmodel"
)

type CommanType string

const (
	NilCommandType        CommanType = ""
	ZoomAreaCommanType    CommanType = "ZOOM_AREA"
	BuildItemCommanType   CommanType = "BUILD_ITEM"
	DestroyItemCommanType CommanType = "DESTROY_ITEM"
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

type ZoomAreaCommand struct {
	Type    CommanType `json:"type"`
	Payload struct {
		Area       areaviewmodel.ViewModel `json:"area"`
		ActionedAt time.Time               `json:"actionedAt"`
	} `json:"payload"`
}
