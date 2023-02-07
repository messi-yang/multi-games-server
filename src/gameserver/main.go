package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/commonmemrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/redispub"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/infrastructure/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/interface/inteventcontroller"
)

func main() {
	itemRepo := commonmemrepo.NewItemMemRepo()
	liveGameRepo := memrepo.NewLiveGameMemRepo()
	intEventPublisher := redispub.New()
	liveGameAppService := appservice.NewLiveGameAppService(
		liveGameRepo,
		itemRepo,
		intEventPublisher,
	)

	mapSize, _ := commonmodel.NewSizeVo(200, 200)
	liveGameAppService.LoadGame(viewmodel.NewSizeVm(mapSize), "20716447-6514-4eac-bd05-e558ca72bf3c")

	inteventcontroller.NewLiveGameIntEventController(liveGameAppService)
}
