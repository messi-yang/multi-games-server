package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	game_pg_repo "github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/redisrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	iam_pg_repo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type CreatePlayerUseCase struct {
	playerRepo playermodel.PlayerRepo
	userRepo   usermodel.UserRepo
	roomRepo   roommodel.RoomRepo
}

func NewCreatePlayerUseCase(playerRepo playermodel.PlayerRepo,
	userRepo usermodel.UserRepo, roomRepo roommodel.RoomRepo) CreatePlayerUseCase {
	return CreatePlayerUseCase{playerRepo, userRepo, roomRepo}
}

func ProvideCreatePlayerUseCase(uow pguow.Uow) CreatePlayerUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	playerRepo := redisrepo.NewPlayerRepo(domainEventDispatcher)
	userRepo := iam_pg_repo.NewUserRepo(uow, domainEventDispatcher)
	roomRepo := game_pg_repo.NewRoomRepo(uow, domainEventDispatcher)

	return NewCreatePlayerUseCase(playerRepo, userRepo, roomRepo)
}

func (useCase *CreatePlayerUseCase) Execute(roomIdDto uuid.UUID, userIdDto *uuid.UUID) (newPlayerDto dto.PlayerDto, err error) {
	roomId := globalcommonmodel.NewRoomId(roomIdDto)
	_, err = useCase.roomRepo.Get(roomId)
	if err != nil {
		return newPlayerDto, err
	}

	newPlayer := playermodel.NewPlayer(
		roomId,
		"Guest",
	)

	if userIdDto != nil {
		user, err := useCase.userRepo.Get(globalcommonmodel.NewUserId(*userIdDto))
		if err != nil {
			return newPlayerDto, err
		}
		newPlayer.UpdateName(user.GetFriendlyName().String())
	}

	if err = useCase.playerRepo.Add(newPlayer); err != nil {
		return newPlayerDto, err
	}

	return dto.NewPlayerDto(newPlayer), nil
}
