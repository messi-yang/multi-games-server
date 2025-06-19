package pgrepo

import (
	"errors"

	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/roomaccessmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
)

type roomMemberRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewRoomMemberRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository roomaccessmodel.RoomMemberRepo) {
	return &roomMemberRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *roomMemberRepo) Add(roomMember roomaccessmodel.RoomMember) error {
	roomMemberModel := pgmodel.NewRoomMemberModel(roomMember)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&roomMemberModel).Error
	})
}

func (repo *roomMemberRepo) Get(roomMemberId roomaccessmodel.RoomMemberId) (roomMember roomaccessmodel.RoomMember, err error) {
	roomMemberModel := pgmodel.RoomMemberModel{Id: roomMemberId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&roomMemberModel).Error
	}); err != nil {
		return roomMember, err
	}
	roomMember, err = pgmodel.ParseRoomMemberModel(roomMemberModel)
	if err != nil {
		return roomMember, err
	}

	return roomMember, nil
}

func (repo *roomMemberRepo) Delete(roomMember roomaccessmodel.RoomMember) error {
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Delete(&pgmodel.RoomMemberModel{}, roomMember.GetId().Uuid()).Error
	})
}

func (repo *roomMemberRepo) GetRoomMemberOfUser(
	roomId globalcommonmodel.RoomId,
	userId globalcommonmodel.UserId,
) (*roomaccessmodel.RoomMember, error) {
	roomMemberModel := pgmodel.RoomMemberModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"room_id = ? AND user_id = ?",
			roomId.Uuid(),
			userId.Uuid(),
		).First(&roomMemberModel).Error
	}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	roomMember, err := pgmodel.ParseRoomMemberModel(roomMemberModel)
	if err != nil {
		return nil, err
	}
	return &roomMember, nil
}

func (repo *roomMemberRepo) GetRoomMembersInRoom(roomId globalcommonmodel.RoomId) (roomMembers []roomaccessmodel.RoomMember, err error) {
	roomMemberModels := []pgmodel.RoomMemberModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"room_id = ?",
			roomId.Uuid(),
		).Find(&roomMemberModels, pgmodel.RoomMemberModel{}).Error
	}); err != nil {
		return roomMembers, err
	}

	roomMembers, err = commonutil.MapWithError(roomMemberModels, func(_ int, roomMemberModel pgmodel.RoomMemberModel) (roomMember roomaccessmodel.RoomMember, err error) {
		return pgmodel.ParseRoomMemberModel(roomMemberModel)
	})
	if err != nil {
		return roomMembers, err
	}
	return roomMembers, nil
}
