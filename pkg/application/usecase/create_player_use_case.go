package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	iam_pg_repo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	world_pg_repo "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/redisrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
)

type CreatePlayerUseCase struct {
	itemRepo   itemmodel.ItemRepo
	playerRepo playermodel.PlayerRepo
	userRepo   usermodel.UserRepo
	worldRepo  worldmodel.WorldRepo
}

func NewCreatePlayerUseCase(itemRepo itemmodel.ItemRepo, playerRepo playermodel.PlayerRepo,
	userRepo usermodel.UserRepo, worldRepo worldmodel.WorldRepo) CreatePlayerUseCase {
	return CreatePlayerUseCase{itemRepo, playerRepo, userRepo, worldRepo}
}

func ProvideCreatePlayerUseCase(uow pguow.Uow) CreatePlayerUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	itemRepo := world_pg_repo.NewItemRepo(uow, domainEventDispatcher)
	playerRepo := redisrepo.NewPlayerRepo(domainEventDispatcher)
	userRepo := iam_pg_repo.NewUserRepo(uow, domainEventDispatcher)
	worldRepo := world_pg_repo.NewWorldRepo(uow, domainEventDispatcher)

	return NewCreatePlayerUseCase(itemRepo, playerRepo, userRepo, worldRepo)
}

func (useCase *CreatePlayerUseCase) Execute(worldIdDto uuid.UUID, userIdDto *uuid.UUID) (newPlayerDto dto.PlayerDto, err error) {
	worldId := globalcommonmodel.NewWorldId(worldIdDto)
	_, err = useCase.worldRepo.Get(worldId)
	if err != nil {
		return newPlayerDto, err
	}

	item, err := useCase.itemRepo.GetFirstItem()
	if err != nil {
		return newPlayerDto, err
	}

	direction := worldcommonmodel.NewDownDirection()
	newPlayer := playermodel.NewPlayer(
		worldId,
		"Guest",
		direction,
		commonutil.ToPointer(item.GetId()),
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
