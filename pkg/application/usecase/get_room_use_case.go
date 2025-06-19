package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type GetRoomUseCase struct {
	roomRepo roommodel.RoomRepo
}

func NewGetRoomUseCase(roomRepo roommodel.RoomRepo) GetRoomUseCase {
	return GetRoomUseCase{roomRepo}
}

func ProvideGetRoomUseCase(uow pguow.Uow) GetRoomUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	roomRepo := pgrepo.NewRoomRepo(uow, domainEventDispatcher)

	return NewGetRoomUseCase(roomRepo)
}

func (useCase *GetRoomUseCase) Execute(roomIdDto uuid.UUID) (roomDto dto.RoomDto, err error) {
	roomId := globalcommonmodel.NewRoomId(roomIdDto)
	room, err := useCase.roomRepo.Get(roomId)
	if err != nil {
		return roomDto, err
	}

	return dto.NewRoomDto(room), nil
}
