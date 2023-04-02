package gameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type ChangeHeldItemCommand struct {
	WorldId  worldmodel.WorldIdVo
	PlayerId playermodel.PlayerIdVo
	ItemId   itemmodel.ItemIdVo
}

type PlaceItemCommand struct {
	WorldId  worldmodel.WorldIdVo
	PlayerId playermodel.PlayerIdVo
}

type DestroyItemCommand struct {
	WorldId  worldmodel.WorldIdVo
	PlayerId playermodel.PlayerIdVo
}

type AddPlayerCommand struct {
	WorldId  worldmodel.WorldIdVo
	PlayerId playermodel.PlayerIdVo
}

type MovePlayerCommand struct {
	WorldId   worldmodel.WorldIdVo
	PlayerId  playermodel.PlayerIdVo
	Direction commonmodel.DirectionVo
}

type RemovePlayerCommand struct {
	WorldId  worldmodel.WorldIdVo
	PlayerId playermodel.PlayerIdVo
}
