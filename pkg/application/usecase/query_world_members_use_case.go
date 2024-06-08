package usecase

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/worldaccessmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type QueryWorldMembersUseCase struct {
	worldMemberRepo worldaccessmodel.WorldMemberRepo
	userRepo        usermodel.UserRepo
}

func NewQueryWorldMembersUseCase(worldMemberRepo worldaccessmodel.WorldMemberRepo, userRepo usermodel.UserRepo) QueryWorldMembersUseCase {
	return QueryWorldMembersUseCase{worldMemberRepo, userRepo}
}

func ProvideQueryWorldMembersUseCase(uow pguow.Uow) QueryWorldMembersUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	worldMemberRepo := pgrepo.NewWorldMemberRepo(uow, domainEventDispatcher)
	userRepo := pgrepo.NewUserRepo(uow, domainEventDispatcher)

	return NewQueryWorldMembersUseCase(worldMemberRepo, userRepo)
}

func (useCase *QueryWorldMembersUseCase) Execute(worldIdDto uuid.UUID, userIdDto uuid.UUID) (worldMemberDtos []dto.WorldMemberDto, err error) {
	worldId := globalcommonmodel.NewWorldId(worldIdDto)
	userId := globalcommonmodel.NewUserId(userIdDto)
	worldMember, err := useCase.worldMemberRepo.GetWorldMemberOfUser(worldId, userId)
	if err != nil {
		return worldMemberDtos, err
	}

	if worldMember == nil {
		return worldMemberDtos, fmt.Errorf("you're not permitted to do this")
	}

	worldPermission := worldaccessmodel.NewWorldPermission(worldMember.GetRole())
	if !worldPermission.CanGetWorldMembers() {
		return worldMemberDtos, fmt.Errorf("you're not permitted to do this")
	}

	worldMembers, err := useCase.worldMemberRepo.GetWorldMembersInWorld(worldId)
	if err != nil {
		return worldMemberDtos, err
	}

	userIds := lo.Map(worldMembers, func(worldMember worldaccessmodel.WorldMember, _ int) globalcommonmodel.UserId {
		return worldMember.GeUserId()
	})

	users, err := useCase.userRepo.GetUsersOfIds(userIds)
	if err != nil {
		return worldMemberDtos, err
	}

	userMap := lo.KeyBy(users, func(user usermodel.User) globalcommonmodel.UserId {
		return user.GetId()
	})

	return lo.Map(worldMembers, func(worldMember worldaccessmodel.WorldMember, _ int) dto.WorldMemberDto {
		return dto.NewWorldMemberDto(worldMember, userMap[worldMember.GeUserId()])
	}), nil
}
