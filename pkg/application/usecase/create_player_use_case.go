package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	iam_pg_repo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	world_pg_repo "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/redisrepo"
	"github.com/google/uuid"
)

type CreatePlayerUseCase struct {
	playerRepo playermodel.PlayerRepo
	userRepo   usermodel.UserRepo
	worldRepo  worldmodel.WorldRepo
}

func NewCreatePlayerUseCase(playerRepo playermodel.PlayerRepo,
	userRepo usermodel.UserRepo, worldRepo worldmodel.WorldRepo) CreatePlayerUseCase {
	return CreatePlayerUseCase{playerRepo, userRepo, worldRepo}
}

func ProvideCreatePlayerUseCase(uow pguow.Uow) CreatePlayerUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	playerRepo := redisrepo.NewPlayerRepo(domainEventDispatcher)
	userRepo := iam_pg_repo.NewUserRepo(uow, domainEventDispatcher)
	worldRepo := world_pg_repo.NewWorldRepo(uow, domainEventDispatcher)

	return NewCreatePlayerUseCase(playerRepo, userRepo, worldRepo)
}

func (useCase *CreatePlayerUseCase) Execute(worldIdDto uuid.UUID, userIdDto *uuid.UUID) (newPlayerDto dto.PlayerDto, err error) {
	worldId := globalcommonmodel.NewWorldId(worldIdDto)
	_, err = useCase.worldRepo.Get(worldId)
	if err != nil {
		return newPlayerDto, err
	}

	direction := worldcommonmodel.NewDownDirection()
	newPlayer := playermodel.NewPlayer(
		worldId,
		"Guest",
		direction,
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
