package usecase

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	game_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/roomaccessmodel"
	iam_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type UpdateRoomUseCase struct {
	roomRepo       roommodel.RoomRepo
	roomMemberRepo roomaccessmodel.RoomMemberRepo
}

func NewUpdateRoomUseCase(roomRepo roommodel.RoomRepo, roomMemberRepo roomaccessmodel.RoomMemberRepo) UpdateRoomUseCase {
	return UpdateRoomUseCase{roomRepo, roomMemberRepo}
}

func ProvideUpdateRoomUseCase(uow pguow.Uow) UpdateRoomUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	roomRepo := game_pgrepo.NewRoomRepo(uow, domainEventDispatcher)

	roomMemberRepo := iam_pgrepo.NewRoomMemberRepo(uow, domainEventDispatcher)

	return NewUpdateRoomUseCase(roomRepo, roomMemberRepo)
}

func (useCase *UpdateRoomUseCase) Execute(useIdDto uuid.UUID, roomIdDto uuid.UUID, roomName string) (
	updatedRoomDto dto.RoomDto, err error,
) {
	roomId := globalcommonmodel.NewRoomId(roomIdDto)
	userId := globalcommonmodel.NewUserId(useIdDto)

	roomMember, err := useCase.roomMemberRepo.GetRoomMemberOfUser(roomId, userId)
	if err != nil {
		return updatedRoomDto, err
	}

	if roomMember == nil {
		return updatedRoomDto, fmt.Errorf("you're not permitted to do this")
	}

	roomPermission := roomaccessmodel.NewRoomPermission(roomMember.GetRole())
	if !roomPermission.CanUpdateRoom() {
		return updatedRoomDto, fmt.Errorf("you're not permitted to do this")
	}

	room, err := useCase.roomRepo.Get(roomId)
	if err != nil {
		return updatedRoomDto, err
	}
	room.ChangeName(roomName)
	if err = useCase.roomRepo.Update(room); err != nil {
		return updatedRoomDto, err
	}

	updatedRoom, err := useCase.roomRepo.Get(roomId)
	if err != nil {
		return updatedRoomDto, err
	}

	return dto.NewRoomDto(updatedRoom), nil
}
