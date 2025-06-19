package usecase

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/roomaccessmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type QueryRoomMembersUseCase struct {
	roomMemberRepo roomaccessmodel.RoomMemberRepo
	userRepo       usermodel.UserRepo
}

func NewQueryRoomMembersUseCase(roomMemberRepo roomaccessmodel.RoomMemberRepo, userRepo usermodel.UserRepo) QueryRoomMembersUseCase {
	return QueryRoomMembersUseCase{roomMemberRepo, userRepo}
}

func ProvideQueryRoomMembersUseCase(uow pguow.Uow) QueryRoomMembersUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	roomMemberRepo := pgrepo.NewRoomMemberRepo(uow, domainEventDispatcher)
	userRepo := pgrepo.NewUserRepo(uow, domainEventDispatcher)

	return NewQueryRoomMembersUseCase(roomMemberRepo, userRepo)
}

func (useCase *QueryRoomMembersUseCase) Execute(roomIdDto uuid.UUID, userIdDto uuid.UUID) (roomMemberDtos []dto.RoomMemberDto, err error) {
	roomId := globalcommonmodel.NewRoomId(roomIdDto)
	userId := globalcommonmodel.NewUserId(userIdDto)
	roomMember, err := useCase.roomMemberRepo.GetRoomMemberOfUser(roomId, userId)
	if err != nil {
		return roomMemberDtos, err
	}

	if roomMember == nil {
		return roomMemberDtos, fmt.Errorf("you're not permitted to do this")
	}

	roomPermission := roomaccessmodel.NewRoomPermission(roomMember.GetRole())
	if !roomPermission.CanGetRoomMembers() {
		return roomMemberDtos, fmt.Errorf("you're not permitted to do this")
	}

	roomMembers, err := useCase.roomMemberRepo.GetRoomMembersInRoom(roomId)
	if err != nil {
		return roomMemberDtos, err
	}

	userIds := lo.Map(roomMembers, func(roomMember roomaccessmodel.RoomMember, _ int) globalcommonmodel.UserId {
		return roomMember.GeUserId()
	})

	users, err := useCase.userRepo.GetUsersOfIds(userIds)
	if err != nil {
		return roomMemberDtos, err
	}

	userMap := lo.KeyBy(users, func(user usermodel.User) globalcommonmodel.UserId {
		return user.GetId()
	})

	return lo.Map(roomMembers, func(roomMember roomaccessmodel.RoomMember, _ int) dto.RoomMemberDto {
		return dto.NewRoomMemberDto(roomMember, userMap[roomMember.GeUserId()])
	}), nil
}
