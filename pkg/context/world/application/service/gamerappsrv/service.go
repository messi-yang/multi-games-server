package gamerappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/gamermodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	GetGamerByUserId(GetGamerByUserIdQuery) (dto.GamerDto, error)
	CreateGamer(CreateGamerCommand) (gamerIdDto uuid.UUID, err error)
	QueryGamers(QueryGamersQuery) ([]dto.GamerDto, error)
}

type serve struct {
	gamerRepo gamermodel.GamerRepo
}

func NewService(
	gamerRepo gamermodel.GamerRepo,
) Service {
	return &serve{
		gamerRepo: gamerRepo,
	}
}

func (serve *serve) GetGamerByUserId(query GetGamerByUserIdQuery) (gamerDto dto.GamerDto, err error) {
	userId := sharedkernelmodel.NewUserId(query.UserId)
	gamer, err := serve.gamerRepo.GetGamerByUserId(userId)
	if err != nil {
		return gamerDto, err
	}

	return dto.NewGamerDto(gamer), nil
}

func (serve *serve) CreateGamer(command CreateGamerCommand) (gamerIdDto uuid.UUID, err error) {
	userId := sharedkernelmodel.NewUserId(command.UserId)
	_, gamerFound, err := serve.gamerRepo.FindGamerByUserId(userId)
	if err != nil {
		return gamerIdDto, err
	}
	if gamerFound {
		return gamerIdDto, fmt.Errorf("already has a gamer with userId of %s", userId.Uuid().String())
	}
	newGamer := gamermodel.NewGamer(userId, 0, 1)
	if err = serve.gamerRepo.Add(newGamer); err != nil {
		return gamerIdDto, err
	}
	return newGamer.GetId().Uuid(), nil
}

func (serve *serve) QueryGamers(query QueryGamersQuery) (itemDtos []dto.GamerDto, err error) {
	gamers, err := serve.gamerRepo.GetAll()
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(gamers, func(gamer gamermodel.Gamer, _ int) dto.GamerDto {
		return dto.NewGamerDto(gamer)
	}), nil
}
