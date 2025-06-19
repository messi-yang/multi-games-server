package pgrepo

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/samber/lo"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

type roomRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewRoomRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository roommodel.RoomRepo) {
	return &roomRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *roomRepo) Add(room roommodel.Room) error {
	roomModel := pgmodel.NewRoomModel(room)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&roomModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&room)
}

func (repo *roomRepo) Update(room roommodel.Room) error {
	roomModel := pgmodel.NewRoomModel(room)
	roomModel.UpdatedAt = time.Now()
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Model(&pgmodel.RoomModel{}).Where(
			"id = ?",
			room.GetId().Uuid(),
		).Select("*").Updates(&roomModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&room)
}

func (repo *roomRepo) Delete(room roommodel.Room) error {
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Delete(&pgmodel.RoomModel{}, room.GetId().Uuid()).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&room)
}

func (repo *roomRepo) Get(roomId globalcommonmodel.RoomId) (room roommodel.Room, err error) {
	roomModel := pgmodel.RoomModel{Id: roomId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&roomModel).Error
	}); err != nil {
		return room, err
	}
	return pgmodel.ParseRoomModel(roomModel), nil
}

func (repo *roomRepo) Query(limit int, offset int) (rooms []roommodel.Room, err error) {
	var roomModels []pgmodel.RoomModel

	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Limit(limit).Offset(offset).Order("created_at desc").Find(&roomModels).Error
	}); err != nil {
		return rooms, err
	}

	return lo.Map(roomModels, func(roomModel pgmodel.RoomModel, _ int) roommodel.Room {
		return pgmodel.ParseRoomModel(roomModel)
	}), nil
}

func (repo *roomRepo) GetRoomsOfUser(userId globalcommonmodel.UserId) (rooms []roommodel.Room, err error) {
	var roomModels []pgmodel.RoomModel

	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Order("created_at desc").Find(
			&roomModels,
			pgmodel.RoomModel{
				UserId: userId.Uuid(),
			},
		).Error
	}); err != nil {
		return rooms, err
	}

	return commonutil.MapWithError[pgmodel.RoomModel](roomModels, func(_ int, roomModel pgmodel.RoomModel) (roommodel.Room, error) {
		return pgmodel.ParseRoomModel(roomModel), nil
	})
}
