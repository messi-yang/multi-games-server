package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/addplayerrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/buliditemrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/destroyitemrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/removeplayerrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/zoomarearequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/messaging/intgreventpublisher"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/areaviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/coordinateviewmodel"
)

type Service interface {
	RequestToAddPlayer(rawLiveGameId string, rawPlayerId string)
	RequestToZoomArea(rawLiveGameId string, rawPlayerId string, rawArea areaviewmodel.ViewModel)
	RequestToBuildItem(rawLiveGameId string, rawCoordinate coordinateviewmodel.ViewModel, rawItemId string)
	RequestToDestroyItem(rawLiveGameId string, rawCoordinate coordinateviewmodel.ViewModel)
	RequestToRemovePlayer(rawLiveGameId string, rawPlayerId string)
}

type serve struct {
	intgrEventPublisher intgreventpublisher.Publisher
}

func New(intgrEventPublisher intgreventpublisher.Publisher) Service {
	return &serve{intgrEventPublisher: intgrEventPublisher}
}

func (serve *serve) RequestToAddPlayer(rawLiveGameId string, rawPlayerId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		addplayerrequestedintgrevent.New(rawLiveGameId, rawPlayerId).Serialize(),
	)
}

func (serve *serve) RequestToZoomArea(rawLiveGameId string, rawPlayerId string, rawArea areaviewmodel.ViewModel) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		zoomarearequestedintgrevent.New(rawLiveGameId, rawPlayerId, rawArea).Serialize(),
	)
}

func (serve *serve) RequestToBuildItem(rawLiveGameId string, rawCoordinate coordinateviewmodel.ViewModel, rawItemId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		buliditemrequestedintgrevent.New(rawLiveGameId, rawCoordinate, rawItemId).Serialize(),
	)
}

func (serve *serve) RequestToDestroyItem(rawLiveGameId string, rawCoordinate coordinateviewmodel.ViewModel) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		destroyitemrequestedintgrevent.New(rawLiveGameId, rawCoordinate).Serialize(),
	)
}

func (serve *serve) RequestToRemovePlayer(rawLiveGameId string, rawPlayerId string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		removeplayerrequestedintgrevent.New(rawLiveGameId, rawPlayerId).Serialize(),
	)
}
