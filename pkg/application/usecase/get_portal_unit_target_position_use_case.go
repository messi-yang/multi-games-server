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

type GetPortalUnitTargetPositionUseCase struct {
	portalUnitRepo portalunitmodel.PortalUnitRepo
}

func NewGetPortalUnitTargetPositionUseCase(portalUnitRepo portalunitmodel.PortalUnitRepo) GetPortalUnitTargetPositionUseCase {
	return GetPortalUnitTargetPositionUseCase{portalUnitRepo}
}

func ProvideGetPortalUnitTargetPositionUseCase(uow pguow.Uow) GetPortalUnitTargetPositionUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	portalUnitRepo := pgrepo.NewPortalUnitRepo(uow, domainEventDispatcher)
	return NewGetPortalUnitTargetPositionUseCase(portalUnitRepo)
}

func (useCase *GetPortalUnitTargetPositionUseCase) Execute(idDto uuid.UUID) (targetPosition *dto.PositionDto, err error) {
	portalUnit, err := useCase.portalUnitRepo.Get(portalunitmodel.NewPortalUnitId(idDto))
	if err != nil {
		return targetPosition, err
	}

	targetPortalUnitId := portalUnit.GetTargetUnitId()
	if targetPortalUnitId == nil {
		return nil, err
	}

	targetPortalUnit, err := useCase.portalUnitRepo.Get(*targetPortalUnitId)
	if err != nil {
		return targetPosition, err
	}

	return commonutil.ToPointer(dto.NewPositionDto(targetPortalUnit.GetPosition())), nil
}
