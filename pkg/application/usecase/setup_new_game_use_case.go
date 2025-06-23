package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/service"
	game_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type SetupNewGameUseCase struct {
	roomService service.RoomService
}

func NewSetupNewGameUseCase(roomService service.RoomService) SetupNewGameUseCase {
	return SetupNewGameUseCase{roomService}
}

func ProvideSetupNewGameUseCase(uow pguow.Uow) SetupNewGameUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	gameAccountRepo := game_pgrepo.NewGameAccountRepo(uow, domainEventDispatcher)
	roomRepo := game_pgrepo.NewRoomRepo(uow, domainEventDispatcher)
	gameRepo := game_pgrepo.NewGameRepo(uow, domainEventDispatcher)
	roomService := service.NewRoomService(gameAccountRepo, roomRepo, gameRepo)

	return NewSetupNewGameUseCase(roomService)
}

func (useCase *SetupNewGameUseCase) Execute(roomId uuid.UUID, gameName string) (gameDto dto.GameDto, err error) {
	game, err := useCase.roomService.SetupNewGame(globalcommonmodel.NewRoomId(roomId), gameName)
	if err != nil {
		return gameDto, err
	}
	return dto.NewGameDto(game), nil
}
