package memdomaineventhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domainevent/memdomainevent"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/worldaccessappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type WorldCreatedHandler struct {
}

func NewWorldCreatedHandler() memdomainevent.Handler {
	return &WorldCreatedHandler{}
}

func (handler WorldCreatedHandler) Handle(uow pguow.Uow, domainEvent domain.DomainEvent) error {
	worldCreated := domainEvent.(sharedkernelmodel.WorldCreated)

	worldAccessAppService := providedependency.ProvideWorldAccessAppService(uow)

	return worldAccessAppService.AddWorldMember(worldaccessappsrv.AddWorldMemberCommand{
		UserId:  worldCreated.GetUserId().Uuid(),
		WorldId: worldCreated.GetWorldId().Uuid(),
		Role:    "owner",
	})
}
