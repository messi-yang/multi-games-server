package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
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

type AddPlayerToWorldUseCase struct {
	itemRepo   itemmodel.ItemRepo
	playerRepo playermodel.PlayerRepo
	userRepo   usermodel.UserRepo
	worldRepo  worldmodel.WorldRepo
}

func NewAddPlayerToWorldUseCase(itemRepo itemmodel.ItemRepo, playerRepo playermodel.PlayerRepo,
	userRepo usermodel.UserRepo, worldRepo worldmodel.WorldRepo) AddPlayerToWorldUseCase {
	return AddPlayerToWorldUseCase{itemRepo, playerRepo, userRepo, worldRepo}
}

func ProvideAddPlayerToWorldUseCase(uow pguow.Uow) AddPlayerToWorldUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	itemRepo := world_pg_repo.NewItemRepo(uow, domainEventDispatcher)
	playerRepo := redisrepo.NewPlayerRepo(domainEventDispatcher)
	userRepo := iam_pg_repo.NewUserRepo(uow, domainEventDispatcher)
	worldRepo := world_pg_repo.NewWorldRepo(uow, domainEventDispatcher)

	return NewAddPlayerToWorldUseCase(itemRepo, playerRepo, userRepo, worldRepo)
}

func (useCase *AddPlayerToWorldUseCase) Execute(worldIdDto uuid.UUID, userIdDto *uuid.UUID) (newPlayerIdDto uuid.UUID, err error) {
	worldId := globalcommonmodel.NewWorldId(worldIdDto)
	_, err = useCase.worldRepo.Get(worldId)
	if err != nil {
		return newPlayerIdDto, err
	}

	item, err := useCase.itemRepo.GetFirstItem()
	if err != nil {
		return newPlayerIdDto, err
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
			return newPlayerIdDto, err
		}
		newPlayer.UpdateName(user.GetFriendlyName().String())
	}

	if err = useCase.playerRepo.Add(newPlayer); err != nil {
		return newPlayerIdDto, err
	}

	return newPlayer.GetId().Uuid(), nil
}
