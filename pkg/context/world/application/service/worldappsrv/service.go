package worldappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
)

var (
	ErrNotPermitted = fmt.Errorf("not permitted to perform this action")
)

type Service interface {
	GetWorld(GetWorldQuery) (dto.WorldDto, error)
}

type serve struct {
	worldRepo worldmodel.WorldRepo
}

func NewService(
	worldRepo worldmodel.WorldRepo,
) Service {
	return &serve{
		worldRepo: worldRepo,
	}
}

func (serve *serve) GetWorld(query GetWorldQuery) (worldDto dto.WorldDto, err error) {
	worldId := globalcommonmodel.NewWorldId(query.WorldId)
	world, err := serve.worldRepo.Get(worldId)
	if err != nil {
		return worldDto, err
	}
	return dto.NewWorldDto(world), nil
}
