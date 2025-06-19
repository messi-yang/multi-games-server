package redisrepo

import (
	"fmt"
	"sync"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/cache/rediscache"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commandmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/jsonutil"
)

func getCommandCacheKey(commandId commandmodel.CommandId) string {
	return fmt.Sprintf("command:%s", commandId.Uuid())
}

type commandRepo struct {
	redisCache            rediscache.RedisCacher
	mutex                 sync.RWMutex
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewCommandRepo(domainEventDispatcher domain.DomainEventDispatcher) (repository commandmodel.CommandRepo) {
	return &commandRepo{
		redisCache:            rediscache.NewRedisCacher(),
		mutex:                 sync.RWMutex{},
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *commandRepo) Add(command commandmodel.Command) error {
	commandDto := dto.NewCommandDto(command)
	commandDtoBytes := string(jsonutil.Marshal(commandDto))
	fmt.Println(commandDto)
	return repo.redisCache.Set(
		getCommandCacheKey(command.GetId()),
		commandDtoBytes,
		60*time.Minute,
	)
}
