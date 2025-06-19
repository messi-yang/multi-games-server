package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/service"
	game_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/roomaccessmodel"
	iam_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type CreateRoomUseCase struct {
	roomService    service.RoomService
	roomMemberRepo roomaccessmodel.RoomMemberRepo
}

func NewCreateRoomUseCase(roomService service.RoomService, roomMemberRepo roomaccessmodel.RoomMemberRepo) CreateRoomUseCase {
	return CreateRoomUseCase{roomService, roomMemberRepo}
}

func ProvideCreateRoomUseCase(uow pguow.Uow) CreateRoomUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	gameAccountRepo := game_pgrepo.NewGameAccountRepo(uow, domainEventDispatcher)
	roomRepo := game_pgrepo.NewRoomRepo(uow, domainEventDispatcher)
	roomService := service.NewRoomService(gameAccountRepo, roomRepo)

	roomMemberRepo := iam_pgrepo.NewRoomMemberRepo(uow, domainEventDispatcher)

	return NewCreateRoomUseCase(roomService, roomMemberRepo)
}

func (useCase *CreateRoomUseCase) Execute(useIdDto uuid.UUID, name string) (roomDto dto.RoomDto, err error) {
	userId := globalcommonmodel.NewUserId(useIdDto)
	newRoom, err := useCase.roomService.CreateRoom(userId, name)
	if err != nil {
		return roomDto, err
	}

	// TODO - Please do add new role in WOORLD_CREATED domain event handler
	roomRole, err := globalcommonmodel.NewRoomRole("owner")
	if err != nil {
		return roomDto, err
	}
	newRoomMember := roomaccessmodel.NewRoomMember(
		newRoom.GetId(),
		userId,
		roomRole,
	)
	if err := useCase.roomMemberRepo.Add(newRoomMember); err != nil {
		return roomDto, err
	}

	return dto.NewRoomDto(newRoom), nil
}
