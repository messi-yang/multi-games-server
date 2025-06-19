package usecase

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/service"
	game_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/roomaccessmodel"
	iam_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type DeleteRoomUseCase struct {
	roomService    service.RoomService
	roomMemberRepo roomaccessmodel.RoomMemberRepo
}

func NewDeleteRoomUseCase(roomService service.RoomService, roomMemberRepo roomaccessmodel.RoomMemberRepo) DeleteRoomUseCase {
	return DeleteRoomUseCase{roomService, roomMemberRepo}
}

func ProvideDeleteRoomUseCase(uow pguow.Uow) DeleteRoomUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	gameAccountRepo := game_pgrepo.NewGameAccountRepo(uow, domainEventDispatcher)
	roomRepo := game_pgrepo.NewRoomRepo(uow, domainEventDispatcher)
	roomService := service.NewRoomService(gameAccountRepo, roomRepo)

	roomMemberRepo := iam_pgrepo.NewRoomMemberRepo(uow, domainEventDispatcher)

	return NewDeleteRoomUseCase(roomService, roomMemberRepo)
}

func (useCase *DeleteRoomUseCase) Execute(useIdDto uuid.UUID, roomIdDto uuid.UUID) (err error) {
	roomId := globalcommonmodel.NewRoomId(roomIdDto)
	userId := globalcommonmodel.NewUserId(useIdDto)
	roomMember, err := useCase.roomMemberRepo.GetRoomMemberOfUser(roomId, userId)
	if err != nil {
		return err
	}

	if roomMember == nil {
		return fmt.Errorf("you're not permitted to do this")
	}

	roomPermission := roomaccessmodel.NewRoomPermission(roomMember.GetRole())
	if !roomPermission.CanDeleteRoom() {
		return fmt.Errorf("you're not permitted to do this")
	}

	// TODO - handle this side effects by using integration events
	roomMembersInRoom, err := useCase.roomMemberRepo.GetRoomMembersInRoom(roomId)
	if err != nil {
		return err
	}
	for _, roomMember := range roomMembersInRoom {
		if err = useCase.roomMemberRepo.Delete(roomMember); err != nil {
			return err
		}
	}

	err = useCase.roomService.DeleteRoom(roomId)
	if err != nil {
		return err
	}

	return nil
}
