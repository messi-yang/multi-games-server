package applicationservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/service"
	liveGameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
)

type LiveGameApplicationService struct {
	LiveGameService *service.LiveGameService
}

type liveGameApplicationServiceConfiguration func(service *LiveGameApplicationService) error

func NewLiveGameApplicationService(cfgs ...liveGameApplicationServiceConfiguration) (*LiveGameApplicationService, error) {
	service := &LiveGameApplicationService{}
	for _, cfg := range cfgs {
		err := cfg(service)
		if err != nil {
			return nil, err
		}
	}
	return service, nil
}

func WithLiveGameService() liveGameApplicationServiceConfiguration {
	liveGameService, _ := service.NewLiveGameService()
	return func(service *LiveGameApplicationService) error {
		service.LiveGameService = liveGameService
		return nil
	}
}

func (service *LiveGameApplicationService) GetAllLiveGameIds() []liveGameValueObject.LiveGameId {
	return service.LiveGameService.GetAllLiveGameIds()
}
