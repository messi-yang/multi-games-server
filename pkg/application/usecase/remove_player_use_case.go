package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/redisrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
)

type RemovePlayerUseCase struct {
	playerRepo playermodel.PlayerRepo
}

func NewRemovePlayerUseCase(playerRepo playermodel.PlayerRepo) RemovePlayerUseCase {
	return RemovePlayerUseCase{playerRepo}
}

func ProvideRemovePlayerUseCase(uow pguow.Uow) RemovePlayerUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	playerRepo := redisrepo.NewPlayerRepo(domainEventDispatcher)

	return NewRemovePlayerUseCase(playerRepo)
}

func (useCase *RemovePlayerUseCase) Execute(roomIdDto uuid.UUID, playerIdDto uuid.UUID) (err error) {
	roomId := globalcommonmodel.NewRoomId(roomIdDto)
	playerId := playermodel.NewPlayerId(playerIdDto)

	player, err := useCase.playerRepo.Get(roomId, playerId)
	if err != nil {
		return err
	}
	return useCase.playerRepo.Delete(player)
}
