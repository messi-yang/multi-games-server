package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type GetMyRoomsUseCase struct {
	roomRepo roommodel.RoomRepo
}

func NewGetMyRoomsUseCase(roomRepo roommodel.RoomRepo) GetMyRoomsUseCase {
	return GetMyRoomsUseCase{roomRepo}
}

func ProvideGetMyRoomsUseCase(uow pguow.Uow) GetMyRoomsUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	roomRepo := pgrepo.NewRoomRepo(uow, domainEventDispatcher)

	return NewGetMyRoomsUseCase(roomRepo)
}

func (useCase *GetMyRoomsUseCase) Execute(useIdDto uuid.UUID) (roomDtos []dto.RoomDto, err error) {
	userId := globalcommonmodel.NewUserId(useIdDto)
	myRooms, err := useCase.roomRepo.GetRoomsOfUser(userId)
	if err != nil {
		return roomDtos, err
	}

	myRoomDtos := lo.Map(myRooms, func(room roommodel.Room, _ int) dto.RoomDto {
		return dto.NewRoomDto(room)
	})

	return myRoomDtos, nil
}
