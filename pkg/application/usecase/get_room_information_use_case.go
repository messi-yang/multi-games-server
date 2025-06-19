package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/redisrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type GetRoomInformationUseCase struct {
	roomRepo   roommodel.RoomRepo
	playerRepo playermodel.PlayerRepo
}

func NewGetRoomInformationUseCase(roomRepo roommodel.RoomRepo, playerRepo playermodel.PlayerRepo) GetRoomInformationUseCase {
	return GetRoomInformationUseCase{roomRepo, playerRepo}
}

func ProvideGetRoomInformationUseCase(uow pguow.Uow) GetRoomInformationUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	roomRepo := pgrepo.NewRoomRepo(uow, domainEventDispatcher)
	playerRepo := redisrepo.NewPlayerRepo(domainEventDispatcher)

	return NewGetRoomInformationUseCase(roomRepo, playerRepo)
}

func (useCase *GetRoomInformationUseCase) Execute(roomIdDto uuid.UUID) (
	roomDto dto.RoomDto, playerDtos []dto.PlayerDto, err error) {
	roomId := globalcommonmodel.NewRoomId(roomIdDto)

	room, err := useCase.roomRepo.Get(roomId)
	if err != nil {
		return roomDto, playerDtos, err
	}
	roomDto = dto.NewRoomDto(room)

	players, err := useCase.playerRepo.GetPlayersOfRoom(roomId)
	if err != nil {
		return roomDto, playerDtos, err
	}
	playerDtos = lo.Map(players, func(_player playermodel.Player, _ int) dto.PlayerDto {
		return dto.NewPlayerDto(_player)
	})

	return roomDto, playerDtos, nil
}
