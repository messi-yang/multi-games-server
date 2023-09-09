package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
)

func RegisterEvents() {
	domainEventRegister := memdomainevent.NewRegister()
	domainEventRegister.Register(worldmodel.WorldCreated{}, NewWorldCreatedHandler())
	domainEventRegister.Register(worldmodel.WorldDeleted{}, NewWorldDeletedHandler())

	domainEventRegister.Register(portalunitmodel.PortalUnitCreated{}, NewPortalUnitCreatedHandler())
	domainEventRegister.Register(portalunitmodel.PortalUnitUpdated{}, NewPortalUnitUpdatedHandler())
	domainEventRegister.Register(portalunitmodel.PortalUnitDeleted{}, NewPortalUnitDeletedHandler())

	domainEventRegister.Register(staticunitmodel.StaticUnitCreated{}, NewStaticUnitCreatedHandler())
	domainEventRegister.Register(staticunitmodel.StaticUnitUpdated{}, NewStaticUnitUpdatedHandler())
	domainEventRegister.Register(staticunitmodel.StaticUnitDeleted{}, NewStaticUnitDeletedHandler())
}
