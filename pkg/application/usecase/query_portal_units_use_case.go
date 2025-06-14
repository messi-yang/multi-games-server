package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	world_pgrepo "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type QueryPortalUnitsUseCase struct {
	portalUnitRepo portalunitmodel.PortalUnitRepo
}

func NewQueryPortalUnitsUseCase(portalUnitRepo portalunitmodel.PortalUnitRepo) QueryPortalUnitsUseCase {
	return QueryPortalUnitsUseCase{portalUnitRepo}
}

func ProvideQueryPortalUnitsUseCase(uow pguow.Uow) QueryPortalUnitsUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	portalUnitRepo := world_pgrepo.NewPortalUnitRepo(uow, domainEventDispatcher)

	return NewQueryPortalUnitsUseCase(portalUnitRepo)
}

func (useCase *QueryPortalUnitsUseCase) Execute(worldIdDto uuid.UUID, limit int, offset int) (portalUnitsDto []dto.PortalUnitDto, err error) {
	worldId := globalcommonmodel.NewWorldId(worldIdDto)

	portalUnits, err := useCase.portalUnitRepo.Query(worldId, limit, offset)
	if err != nil {
		return nil, err
	}

	portalUnitsDto = make([]dto.PortalUnitDto, len(portalUnits))
	for i, portalUnit := range portalUnits {
		portalUnitsDto[i] = dto.NewPortalUnitDto(portalUnit)
	}

	return portalUnitsDto, nil
}
