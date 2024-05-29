package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/embedunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type GetEmbedUnitEmbedCodeUseCase struct {
	embedUnitRepo embedunitmodel.EmbedUnitRepo
}

func NewGetEmbedUnitEmbedCodeUseCase(embedUnitRepo embedunitmodel.EmbedUnitRepo) GetEmbedUnitEmbedCodeUseCase {
	return GetEmbedUnitEmbedCodeUseCase{embedUnitRepo}
}

func ProvideGetEmbedUnitEmbedCodeUseCase(uow pguow.Uow) GetEmbedUnitEmbedCodeUseCase {
	domainEventDispatcher := memdomainevent.NewDispatcher(uow)
	embedUnitRepo := pgrepo.NewEmbedUnitRepo(uow, domainEventDispatcher)
	return NewGetEmbedUnitEmbedCodeUseCase(embedUnitRepo)
}

func (useCase *GetEmbedUnitEmbedCodeUseCase) Execute(idDto uuid.UUID) (embedCodeDto string, err error) {
	embedUnit, err := useCase.embedUnitRepo.Get(embedunitmodel.NewEmbedUnitId(idDto))
	if err != nil {
		return "", err
	}

	return embedUnit.GetEmbedCode().String(), nil
}
