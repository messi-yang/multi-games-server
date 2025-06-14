package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
)

type GetPortalUnitTargetUnitUseCase struct {
	portalUnitRepo portalunitmodel.PortalUnitRepo
}

func NewGetPortalUnitTargetUnitUseCase(portalUnitRepo portalunitmodel.PortalUnitRepo) GetPortalUnitTargetUnitUseCase {
	return GetPortalUnitTargetUnitUseCase{portalUnitRepo}
}

func ProvideGetPortalUnitTargetUnitUseCase(uow pguow.Uow) GetPortalUnitTargetUnitUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	portalUnitRepo := pgrepo.NewPortalUnitRepo(uow, domainEventDispatcher)
	return NewGetPortalUnitTargetUnitUseCase(portalUnitRepo)
}

func (useCase *GetPortalUnitTargetUnitUseCase) Execute(idDto uuid.UUID) (targetPortalUnitDto *dto.PortalUnitDto, err error) {
	portalUnit, err := useCase.portalUnitRepo.Get(portalunitmodel.NewPortalUnitId(idDto))
	if err != nil {
		return nil, err
	}

	targetPortalUnitId := portalUnit.GetTargetUnitId()
	if targetPortalUnitId == nil {
		return nil, err
	}

	targetPortalUnit, err := useCase.portalUnitRepo.Get(*targetPortalUnitId)
	if err != nil {
		return nil, err
	}

	return commonutil.ToPointer(dto.NewPortalUnitDto(targetPortalUnit)), nil
}
