package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commandmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/colorunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/embedunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/signunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/persistence/pgrepo"
	"github.com/google/uuid"
)

type ExecuteCommandUseCase struct {
	unitRepo          unitmodel.UnitRepo
	staticUnitService service.StaticUnitService
	fenceUnitService  service.FenceUnitService
	linkUnitService   service.LinkUnitService
	portalUnitService service.PortalUnitService
	embedUnitService  service.EmbedUnitService
	colorUnitService  service.ColorUnitService
	signUnitService   service.SignUnitService
	unitService       service.UnitService
}

func NewExecuteCommandUseCase(
	unitRepo unitmodel.UnitRepo,
	staticUnitService service.StaticUnitService, fenceUnitService service.FenceUnitService, linkUnitService service.LinkUnitService,
	portalUnitService service.PortalUnitService, embedUnitService service.EmbedUnitService, colorUnitService service.ColorUnitService,
	signUnitService service.SignUnitService,
	unitService service.UnitService,
) ExecuteCommandUseCase {
	return ExecuteCommandUseCase{
		unitRepo,
		staticUnitService, fenceUnitService, linkUnitService,
		portalUnitService, embedUnitService, colorUnitService,
		signUnitService,
		unitService,
	}
}

func ProvideExecuteCommandUseCase(uow pguow.Uow) ExecuteCommandUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	itemRepo := pgrepo.NewItemRepo(uow, domainEventDispatcher)
	unitRepo := pgrepo.NewUnitRepo(uow, domainEventDispatcher)
	staticUnitRepo := pgrepo.NewStaticUnitRepo(uow, domainEventDispatcher)
	fenceUnitRepo := pgrepo.NewFenceUnitRepo(uow, domainEventDispatcher)
	linkUnitRepo := pgrepo.NewLinkUnitRepo(uow, domainEventDispatcher)
	portalUnitRepo := pgrepo.NewPortalUnitRepo(uow, domainEventDispatcher)
	embedUnitRepo := pgrepo.NewEmbedUnitRepo(uow, domainEventDispatcher)
	colorUnitRepo := pgrepo.NewColorUnitRepo(uow, domainEventDispatcher)
	signUnitRepo := pgrepo.NewSignUnitRepo(uow, domainEventDispatcher)
	staticUnitRepoUnitService := service.NewStaticUnitService(unitRepo, staticUnitRepo, itemRepo)
	fenceUnitRepoUnitService := service.NewFenceUnitService(unitRepo, fenceUnitRepo, itemRepo)
	linkUnitRepoUnitService := service.NewLinkUnitService(unitRepo, linkUnitRepo, itemRepo)
	portalUnitRepoUnitService := service.NewPortalUnitService(unitRepo, portalUnitRepo, itemRepo)
	embedUnitRepoUnitService := service.NewEmbedUnitService(unitRepo, embedUnitRepo, itemRepo)
	colorUnitRepoUnitService := service.NewColorUnitService(unitRepo, colorUnitRepo, itemRepo)
	signUnitRepoUnitService := service.NewSignUnitService(unitRepo, signUnitRepo, itemRepo)
	unitRepoUnitService := service.NewUnitService(unitRepo)
	return NewExecuteCommandUseCase(
		unitRepo,
		staticUnitRepoUnitService, fenceUnitRepoUnitService, linkUnitRepoUnitService,
		portalUnitRepoUnitService, embedUnitRepoUnitService, colorUnitRepoUnitService,
		signUnitRepoUnitService,
		unitRepoUnitService,
	)
}

func (useCase *ExecuteCommandUseCase) Execute(worldIdDto uuid.UUID, commandDto dto.CommandDto) error {
	commandName, err := commandmodel.NewCommandName(commandDto.Name)
	if err != nil {
		return err
	}
	command := commandmodel.CreateCommand(
		commandmodel.NewCommandId(commandDto.Id),
		commandDto.Timestamp,
		commandName,
		commandDto.Payload,
	)

	if commandName.IsCreateStaticUnitCommandName() {
		payload, err := command.GetCreateStaticUnitCommandPayload()
		if err != nil {
			return err
		}

		err = useCase.staticUnitService.CreateStaticUnit(
			staticunitmodel.NewStaticUnitId(payload.UnitId),
			globalcommonmodel.NewWorldId(worldIdDto),
			worldcommonmodel.NewItemId(payload.ItemId),
			payload.UnitPosition.ToValueObject(),
			worldcommonmodel.NewDirection(payload.UnitDirection),
		)
		if err != nil {
			return err
		}
	} else if commandName.IsCreateFenceUnitCommandName() {
		payload, err := command.GetCreateFenceUnitCommandPayload()
		if err != nil {
			return err
		}

		err = useCase.fenceUnitService.CreateFenceUnit(
			fenceunitmodel.NewFenceUnitId(payload.UnitId),
			globalcommonmodel.NewWorldId(worldIdDto),
			worldcommonmodel.NewItemId(payload.ItemId),
			payload.UnitPosition.ToValueObject(),
			worldcommonmodel.NewDirection(payload.UnitDirection),
		)
		if err != nil {
			return err
		}
	} else if commandName.IsCreatePortalUnitCommandName() {
		payload, err := command.GetCreatePortalUnitCommandPayload()
		if err != nil {
			return err
		}

		err = useCase.portalUnitService.CreatePortalUnit(
			portalunitmodel.NewPortalUnitId(payload.UnitId),
			globalcommonmodel.NewWorldId(worldIdDto),
			worldcommonmodel.NewItemId(payload.ItemId),
			payload.UnitPosition.ToValueObject(),
			worldcommonmodel.NewDirection(payload.UnitDirection),
		)
		if err != nil {
			return err
		}
	} else if commandName.IsCreateLinkUnitCommandName() {
		payload, err := command.GetCreateLinkUnitCommandPayload()
		if err != nil {
			return err
		}

		url, err := globalcommonmodel.NewUrl(payload.UnitUrl)
		if err != nil {
			return err
		}
		err = useCase.linkUnitService.CreateLinkUnit(
			linkunitmodel.NewLinkUnitId(payload.UnitId),
			globalcommonmodel.NewWorldId(worldIdDto),
			worldcommonmodel.NewItemId(payload.ItemId),
			payload.UnitPosition.ToValueObject(),
			worldcommonmodel.NewDirection(payload.UnitDirection),
			payload.UnitLabel,
			url,
		)
		if err != nil {
			return err
		}
	} else if commandName.IsCreateEmbedUnitCommandName() {
		payload, err := command.GetCreateEmbedUnitCommandPayload()
		if err != nil {
			return err
		}

		embedCode, err := worldcommonmodel.NewEmbedCode(payload.UnitEmbedCode)
		if err != nil {
			return err
		}

		err = useCase.embedUnitService.CreateEmbedUnit(
			embedunitmodel.NewEmbedUnitId(payload.UnitId),
			globalcommonmodel.NewWorldId(worldIdDto),
			worldcommonmodel.NewItemId(payload.ItemId),
			payload.UnitPosition.ToValueObject(),
			worldcommonmodel.NewDirection(payload.UnitDirection),
			payload.UnitLabel,
			embedCode,
		)
		if err != nil {
			return err
		}
	} else if commandName.IsCreateColorUnitCommandName() {
		payload, err := command.GetCreateColorUnitCommandPayload()
		if err != nil {
			return err
		}

		color, err := globalcommonmodel.NewColorFromHexString(payload.UnitColor)
		if err != nil {
			return err
		}

		err = useCase.colorUnitService.CreateColorUnit(
			colorunitmodel.NewColorUnitId(payload.UnitId),
			globalcommonmodel.NewWorldId(worldIdDto),
			worldcommonmodel.NewItemId(payload.ItemId),
			payload.UnitPosition.ToValueObject(),
			worldcommonmodel.NewDirection(payload.UnitDirection),
			payload.UnitLabel,
			color,
		)
		if err != nil {
			return err
		}
	} else if commandName.IsCreateSignUnitCommandName() {
		payload, err := command.GetCreateSignUnitCommandPayload()
		if err != nil {
			return err
		}

		err = useCase.signUnitService.CreateSignUnit(
			signunitmodel.NewSignUnitId(payload.UnitId),
			globalcommonmodel.NewWorldId(worldIdDto),
			worldcommonmodel.NewItemId(payload.ItemId),
			payload.UnitPosition.ToValueObject(),
			worldcommonmodel.NewDirection(payload.UnitDirection),
			payload.UnitLabel,
		)
		if err != nil {
			return err
		}
	} else if commandName.IsRotateUnitCommandName() {
		payload, err := command.GetRotateUnitCommandPayload()
		if err != nil {
			return err
		}

		unitId := unitmodel.NewUnitId(payload.UnitId)
		unit, err := useCase.unitRepo.Get(unitId)
		if err != nil {
			return err
		}

		if unit.GetType().IsPortal() {
			err = useCase.portalUnitService.RotatePortalUnit(portalunitmodel.NewPortalUnitId(payload.UnitId))
		} else if unit.GetType().IsStatic() {
			err = useCase.staticUnitService.RotateStaticUnit(staticunitmodel.NewStaticUnitId(payload.UnitId))
		} else if unit.GetType().IsFence() {
			err = useCase.fenceUnitService.RotateFenceUnit(fenceunitmodel.NewFenceUnitId(payload.UnitId))
		} else if unit.GetType().IsLink() {
			err = useCase.linkUnitService.RotateLinkUnit(linkunitmodel.NewLinkUnitId(payload.UnitId))
		} else if unit.GetType().IsEmbed() {
			err = useCase.embedUnitService.RotateEmbedUnit(embedunitmodel.NewEmbedUnitId(payload.UnitId))
		} else if unit.GetType().IsColor() {
			err = useCase.colorUnitService.RotateColorUnit(colorunitmodel.NewColorUnitId(payload.UnitId))
		} else if unit.GetType().IsSign() {
			err = useCase.signUnitService.RotateSignUnit(signunitmodel.NewSignUnitId(payload.UnitId))
		}
		if err != nil {
			return err
		}

	} else if commandName.IsMoveUnitCommandName() {
		payload, err := command.GetMoveUnitCommandPayload()
		if err != nil {
			return err
		}

		err = useCase.unitService.MoveUnit(unitmodel.NewUnitId(payload.UnitId), payload.UnitPosition.ToValueObject())
		if err != nil {
			return err
		}
	} else if commandName.IsRemoveStaticUnitCommandName() {
		payload, err := command.GetRemoveStaticUnitCommandPayload()
		if err != nil {
			return err
		}

		err = useCase.staticUnitService.RemoveStaticUnit(staticunitmodel.NewStaticUnitId(payload.UnitId))
		if err != nil {
			return err
		}
	} else if commandName.IsRemoveFenceUnitCommandName() {
		payload, err := command.GetRemoveFenceUnitCommandPayload()
		if err != nil {
			return err
		}

		err = useCase.fenceUnitService.RemoveFenceUnit(fenceunitmodel.NewFenceUnitId(payload.UnitId))
		if err != nil {
			return err
		}
	} else if commandName.IsRemovePortalUnitCommandName() {
		payload, err := command.GetRemovePortalUnitCommandPayload()
		if err != nil {
			return err
		}

		err = useCase.portalUnitService.RemovePortalUnit(portalunitmodel.NewPortalUnitId(payload.UnitId))
		if err != nil {
			return err
		}
	} else if commandName.IsRemoveLinkUnitCommandName() {
		payload, err := command.GetRemoveLinkUnitCommandPayload()
		if err != nil {
			return err
		}

		err = useCase.linkUnitService.RemoveLinkUnit(linkunitmodel.NewLinkUnitId(payload.UnitId))
		if err != nil {
			return err
		}
	} else if commandName.IsRemoveEmbedUnitCommandName() {
		payload, err := command.GetRemoveEmbedUnitCommandPayload()
		if err != nil {
			return err
		}

		err = useCase.embedUnitService.RemoveEmbedUnit(embedunitmodel.NewEmbedUnitId(payload.UnitId))
		if err != nil {
			return err
		}
	} else if commandName.IsRemoveColorUnitCommandName() {
		payload, err := command.GetRemoveColorUnitCommandPayload()
		if err != nil {
			return err
		}

		err = useCase.colorUnitService.RemoveColorUnit(colorunitmodel.NewColorUnitId(payload.UnitId))
		if err != nil {
			return err
		}
	} else if commandName.IsRemoveSignUnitCommandName() {
		payload, err := command.GetRemoveSignUnitCommandPayload()
		if err != nil {
			return err
		}

		err = useCase.signUnitService.RemoveSignUnit(signunitmodel.NewSignUnitId(payload.UnitId))
		if err != nil {
			return err
		}
	}

	return nil
}
