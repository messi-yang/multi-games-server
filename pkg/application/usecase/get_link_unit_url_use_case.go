package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type GetLinkUnitUrlUseCase struct {
	linkUnitRepo linkunitmodel.LinkUnitRepo
}

func NewGetLinkUnitUrlUseCase(linkUnitRepo linkunitmodel.LinkUnitRepo) GetLinkUnitUrlUseCase {
	return GetLinkUnitUrlUseCase{linkUnitRepo}
}

func ProvideGetLinkUnitUrlUseCase(uow pguow.Uow) GetLinkUnitUrlUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	linkUnitRepo := pgrepo.NewLinkUnitRepo(uow, domainEventDispatcher)
	return NewGetLinkUnitUrlUseCase(linkUnitRepo)
}

func (useCase *GetLinkUnitUrlUseCase) Execute(idDto uuid.UUID) (urlDto string, err error) {
	linkUnit, err := useCase.linkUnitRepo.Get(linkunitmodel.NewLinkUnitId(idDto))
	if err != nil {
		return "", err
	}

	return linkUnit.GetUrl().String(), nil
}
