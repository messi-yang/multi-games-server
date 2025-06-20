package redisrepo

import (
	"fmt"
	"sync"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/cache/rediscache"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commandmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/jsonutil"
)

func getCommandCacheKey(gameId gamemodel.GameId, commandId commandmodel.CommandId) string {
	return fmt.Sprintf("game:%s:command:%s", gameId.Uuid(), commandId.Uuid())
}

func getGameCommandsScanPattern(gameId gamemodel.GameId) string {
	return fmt.Sprintf("game:%s:command:*", gameId.Uuid())
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
		getCommandCacheKey(command.GetGameId(), command.GetId()),
		commandDtoBytes,
		60*time.Minute,
	)
}

func (repo *commandRepo) GetCommandsOfGame(gameId gamemodel.GameId) (commands []commandmodel.Command, err error) {
	commandDtoStrings, err := repo.redisCache.Scan(getGameCommandsScanPattern(gameId))
	if err != nil {
		return commands, err
	}

	commands = make([]commandmodel.Command, 0)

	for _, commandDtoString := range commandDtoStrings {
		commandDto, err := jsonutil.Unmarshal[dto.CommandDto]([]byte(commandDtoString))
		if err != nil {
			return nil, err
		}
		command, err := dto.ParseCommandDto(commandDto)
		if err != nil {
			return nil, err
		}
		commands = append(commands, command)
	}

	return commands, nil
}
