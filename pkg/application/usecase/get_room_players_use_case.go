package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/redisrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type GetRoomPlayersUseCase struct {
	playerRepo playermodel.PlayerRepo
}

func NewGetRoomPlayersUseCase(playerRepo playermodel.PlayerRepo) GetRoomPlayersUseCase {
	return GetRoomPlayersUseCase{playerRepo}
}

func ProvideGetRoomPlayersUseCase(uow pguow.Uow) GetRoomPlayersUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	playerRepo := redisrepo.NewPlayerRepo(domainEventDispatcher)

	return NewGetRoomPlayersUseCase(playerRepo)
}

func (useCase *GetRoomPlayersUseCase) Execute(roomIdDto uuid.UUID) (
	playerDtos []dto.PlayerDto, err error) {
	roomId := globalcommonmodel.NewRoomId(roomIdDto)

	players, err := useCase.playerRepo.GetPlayersOfRoom(roomId)
	if err != nil {
		return playerDtos, err
	}
	playerDtos = lo.Map(players, func(_player playermodel.Player, _ int) dto.PlayerDto {
		return dto.NewPlayerDto(_player)
	})

	return playerDtos, nil
}
