package appservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainservice"
)

type GameAppService interface{}

type gameAppServe struct {
	gameDomainService domainservice.GameDomainService
}

func NewGameAppService(gameDomainService domainservice.GameDomainService) GameAppService {
	return &gameAppServe{gameDomainService: gameDomainService}
}
