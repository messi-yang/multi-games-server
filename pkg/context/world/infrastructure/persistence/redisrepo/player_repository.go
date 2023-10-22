package redisrepo

import (
	"fmt"
	"sync"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/cache/rediscache"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/jsonutil"
)

func getPlayerCacheKey(worldId globalcommonmodel.WorldId, playerId playermodel.PlayerId) string {
	return fmt.Sprintf("world:%s:player:%s", worldId.Uuid(), playerId.Uuid())
}

func getWorldPlayersScanPattern(worldId globalcommonmodel.WorldId) string {
	return fmt.Sprintf("world:%s:player:*", worldId.Uuid())
}

type playerRepo struct {
	redisCache            rediscache.RedisCacher
	mutex                 sync.RWMutex
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewPlayerRepo(domainEventDispatcher domain.DomainEventDispatcher) (repository playermodel.PlayerRepo) {
	return &playerRepo{
		redisCache:            rediscache.NewRedisCacher(),
		mutex:                 sync.RWMutex{},
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *playerRepo) Add(player playermodel.Player) error {
	playerDto := dto.NewPlayerDto(player)
	playerDtoBytes := string(jsonutil.Marshal(playerDto))
	if err := repo.redisCache.Set(
		getPlayerCacheKey(player.GetWorldId(), player.GetId()),
		playerDtoBytes,
		5*time.Minute,
	); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Update(player playermodel.Player) error {
	playerDto := dto.NewPlayerDto(player)
	playerDtoBytes := string(jsonutil.Marshal(playerDto))

	if err := repo.redisCache.Set(
		getPlayerCacheKey(player.GetWorldId(), player.GetId()),
		playerDtoBytes,
		5*time.Minute,
	); err != nil {
		return err
	}

	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Delete(player playermodel.Player) error {
	if err := repo.redisCache.Del(getPlayerCacheKey(player.GetWorldId(), player.GetId())); err != nil {
		return err
	}

	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Get(worldId globalcommonmodel.WorldId, playerId playermodel.PlayerId) (player playermodel.Player, err error) {
	playerDtoString, err := repo.redisCache.Get(getPlayerCacheKey(worldId, playerId))
	if err != nil {
		return player, err
	}
	playerDto, err := jsonutil.Unmarshal[dto.PlayerDto]([]byte(playerDtoString))
	if err != nil {
		return player, err
	}

	return dto.ParsePlayerDto(playerDto)
}

func (repo *playerRepo) GetPlayersOfWorld(worldId globalcommonmodel.WorldId) (players []playermodel.Player, err error) {
	playerDtoStrings, err := repo.redisCache.Scan(getWorldPlayersScanPattern(worldId))
	if err != nil {
		return players, err
	}

	players = make([]playermodel.Player, 0)

	for _, playerDtoString := range playerDtoStrings {
		playerDto, err := jsonutil.Unmarshal[dto.PlayerDto]([]byte(playerDtoString))
		if err != nil {
			return players, err
		}
		player, err := dto.ParsePlayerDto(playerDto)
		if err != nil {
			return players, err
		}
		players = append(players, player)
	}

	return players, nil
}
