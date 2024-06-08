package playerappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

var (
	ErrPosHasNoUnits       = fmt.Errorf("the position has no units")
	ErrUnitIsNotPortalType = fmt.Errorf("the unit is not in type of portal")
)

type Service interface {
	ChangePlayerAction(ChangePlayerActionCommand) error
	SendPlayerIntoPortal(SendPlayerIntoPortalCommand) error
	ChangePlayerHeldItem(ChangePlayerHeldItemCommand) error
}

type serve struct {
	playerRepo     playermodel.PlayerRepo
	unitRepo       unitmodel.UnitRepo
	portalUnitRepo portalunitmodel.PortalUnitRepo
}

func NewService(
	playerRepo playermodel.PlayerRepo,
	unitRepo unitmodel.UnitRepo,
	portalUnitRepo portalunitmodel.PortalUnitRepo,
) Service {
	return &serve{
		playerRepo:     playerRepo,
		unitRepo:       unitRepo,
		portalUnitRepo: portalUnitRepo,
	}
}

func (serve *serve) ChangePlayerAction(command ChangePlayerActionCommand) error {
	action, err := dto.ParsePlayerActionDto(command.Action)
	if err != nil {
		return err
	}

	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	playerId := playermodel.NewPlayerId(command.PlayerId)

	player, err := serve.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}

	player.ChangeAction(action)

	return serve.playerRepo.Update(player)

}

func (serve *serve) SendPlayerIntoPortal(command SendPlayerIntoPortalCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	playerId := playermodel.NewPlayerId(command.PlayerId)

	player, err := serve.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}

	portalUnit, err := serve.portalUnitRepo.Get(portalunitmodel.NewPortalUnitId(command.UnitId))
	if err != nil {
		return err
	}
	targetPortalUnitId := portalUnit.GetTargetUnitId()
	if targetPortalUnitId == nil {
		return fmt.Errorf("the portal unit has no target portal unit")
	}

	targetPortalUnit, err := serve.portalUnitRepo.Get(*targetPortalUnitId)
	if err != nil {
		return err
	}

	targetPosition := targetPortalUnit.GetPosition()
	player.Teleport(targetPosition.PrecisePosition())

	return serve.playerRepo.Update(player)
}

func (serve *serve) ChangePlayerHeldItem(command ChangePlayerHeldItemCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	playerId := playermodel.NewPlayerId(command.PlayerId)
	itemId := worldcommonmodel.NewItemId(command.ItemId)

	player, err := serve.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}

	player.ChangeHeldItem(itemId)
	return serve.playerRepo.Update(player)
}
