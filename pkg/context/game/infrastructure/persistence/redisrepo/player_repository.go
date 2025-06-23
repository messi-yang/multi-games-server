package redisrepo

import (
	"fmt"
	"sync"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/cache/rediscache"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/jsonutil"
)

func getPlayerCacheKey(roomId globalcommonmodel.RoomId, playerId playermodel.PlayerId) string {
	return fmt.Sprintf("room:%s:player:%s", roomId.Uuid(), playerId.Uuid())
}

func getRoomPlayersScanPattern(roomId globalcommonmodel.RoomId) string {
	return fmt.Sprintf("room:%s:player:*", roomId.Uuid())
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
	return repo.redisCache.Set(
		getPlayerCacheKey(player.GetRoomId(), player.GetId()),
		playerDtoBytes,
		60*time.Minute,
	)
}

func (repo *playerRepo) Update(player playermodel.Player) error {
	playerDto := dto.NewPlayerDto(player)
	playerDtoBytes := string(jsonutil.Marshal(playerDto))

	return repo.redisCache.Set(
		getPlayerCacheKey(player.GetRoomId(), player.GetId()),
		playerDtoBytes,
		60*time.Minute,
	)
}

func (repo *playerRepo) Delete(player playermodel.Player) error {
	return repo.redisCache.Del(getPlayerCacheKey(player.GetRoomId(), player.GetId()))
}

func (repo *playerRepo) Get(roomId globalcommonmodel.RoomId, playerId playermodel.PlayerId) (player playermodel.Player, err error) {
	playerDtoString, err := repo.redisCache.Get(getPlayerCacheKey(roomId, playerId))
	if err != nil {
		return player, err
	}
	playerDto, err := jsonutil.Unmarshal[dto.PlayerDto]([]byte(playerDtoString))
	if err != nil {
		return player, err
	}

	return dto.ParsePlayerDto(playerDto)
}

func (repo *playerRepo) GetPlayerOfUser(roomId globalcommonmodel.RoomId, userId globalcommonmodel.UserId) (player *playermodel.Player, err error) {
	playerDtoStrings, err := repo.redisCache.Scan(getRoomPlayersScanPattern(roomId))
	if err != nil {
		return player, err
	}

	for _, playerDtoString := range playerDtoStrings {
		playerDto, err := jsonutil.Unmarshal[dto.PlayerDto]([]byte(playerDtoString))
		if err != nil {
			return player, err
		}
		if playerDto.UserId != nil && *playerDto.UserId == userId.Uuid() {
			parsedPlayer, err := dto.ParsePlayerDto(playerDto)
			if err != nil {
				return player, err
			}
			return &parsedPlayer, nil
		}
	}

	return player, nil
}

func (repo *playerRepo) GetPlayersOfRoom(roomId globalcommonmodel.RoomId) (players []playermodel.Player, err error) {
	playerDtoStrings, err := repo.redisCache.Scan(getRoomPlayersScanPattern(roomId))
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
