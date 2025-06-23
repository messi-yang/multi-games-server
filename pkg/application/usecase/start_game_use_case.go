package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/service"
	game_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	gameaccount_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	room_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type StartGameUseCase struct {
	roomService service.RoomService
}

func NewStartGameUseCase(roomService service.RoomService) StartGameUseCase {
	return StartGameUseCase{roomService}
}

func ProvideStartGameUseCase(uow pguow.Uow) StartGameUseCase {
	gameRepo := game_pgrepo.NewGameRepo(uow, memdomaineventhandler.NewDispatcher(uow))
	roomRepo := room_pgrepo.NewRoomRepo(uow, memdomaineventhandler.NewDispatcher(uow))
	gameAccountRepo := gameaccount_pgrepo.NewGameAccountRepo(uow, memdomaineventhandler.NewDispatcher(uow))
	roomService := service.NewRoomService(gameAccountRepo, roomRepo, gameRepo)

	return NewStartGameUseCase(roomService)
}

func (useCase *StartGameUseCase) Execute(roomId uuid.UUID, gameId uuid.UUID, gameState map[string]interface{}) (gameDto dto.GameDto, err error) {
	game, err := useCase.roomService.StartGame(globalcommonmodel.NewRoomId(roomId), gamemodel.NewGameId(gameId), gameState)
	if err != nil {
		return gameDto, err
	}
	return dto.NewGameDto(game), nil
}
