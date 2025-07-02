package usecase

import (
	"errors"
	"math/rand/v2"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	game_pg_repo "github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/redisrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/roomaccessmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	iam_pg_repo "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
)

type CreatePlayerUseCase struct {
	playerRepo     playermodel.PlayerRepo
	roomAccessRepo roomaccessmodel.RoomMemberRepo
	userRepo       usermodel.UserRepo
	roomRepo       roommodel.RoomRepo
}

func NewCreatePlayerUseCase(playerRepo playermodel.PlayerRepo,
	roomAccessRepo roomaccessmodel.RoomMemberRepo,
	userRepo usermodel.UserRepo,
	roomRepo roommodel.RoomRepo) CreatePlayerUseCase {
	return CreatePlayerUseCase{playerRepo, roomAccessRepo, userRepo, roomRepo}
}

func ProvideCreatePlayerUseCase(uow pguow.Uow) CreatePlayerUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	playerRepo := redisrepo.NewPlayerRepo(domainEventDispatcher)
	roomAccessRepo := iam_pg_repo.NewRoomMemberRepo(uow, domainEventDispatcher)
	userRepo := iam_pg_repo.NewUserRepo(uow, domainEventDispatcher)
	roomRepo := game_pg_repo.NewRoomRepo(uow, domainEventDispatcher)

	return NewCreatePlayerUseCase(playerRepo, roomAccessRepo, userRepo, roomRepo)
}

func (useCase *CreatePlayerUseCase) Execute(roomIdDto uuid.UUID, userIdDto *uuid.UUID, guestPlayerName string) (newPlayerDto dto.PlayerDto, err error) {
	roomId := globalcommonmodel.NewRoomId(roomIdDto)
	_, err = useCase.roomRepo.Get(roomId)
	if err != nil {
		return newPlayerDto, err
	}

	var name string = guestPlayerName
	var hostPriority float64 = rand.Float64() * 9999999
	var userId *globalcommonmodel.UserId = nil

	if userIdDto != nil {
		playerWithSameUserId, err := useCase.playerRepo.GetPlayerOfUser(roomId, globalcommonmodel.NewUserId(*userIdDto))
		if err != nil {
			return newPlayerDto, err
		}
		if playerWithSameUserId != nil {
			return newPlayerDto, errors.New("user already has a player")
		}

		user, err := useCase.userRepo.Get(globalcommonmodel.NewUserId(*userIdDto))
		if err != nil {
			return newPlayerDto, err
		}
		name = user.GetFriendlyName().String()
		userId = commonutil.ToPointer(user.GetId())

		roomMember, err := useCase.roomAccessRepo.GetRoomMemberOfUser(roomId, user.GetId())
		if err != nil {
			return newPlayerDto, err
		}
		if roomMember.GetRole().IsOwner() {
			hostPriority = -1
		}
	}

	newPlayer := playermodel.NewPlayer(
		roomId,
		userId,
		name,
		hostPriority,
	)

	if err = useCase.playerRepo.Add(newPlayer); err != nil {
		return newPlayerDto, err
	}

	return dto.NewPlayerDto(newPlayer), nil
}
