package memrepo

import (
	"fmt"
	"sync"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/playermodel"
)

var locker = sync.RWMutex{}
var worldPlayerMap = make(map[sharedkernelmodel.WorldId]map[playermodel.PlayerId]playermodel.Player)

type playerRepo struct {
	mutex                 sync.RWMutex
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewPlayerRepo(domainEventDispatcher domain.DomainEventDispatcher) (repository playermodel.PlayerRepo) {
	return &playerRepo{
		mutex:                 sync.RWMutex{},
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *playerRepo) Add(player playermodel.Player) error {
	locker.Lock()

	_, found := worldPlayerMap[player.GetWorldId()]
	if !found {
		worldPlayerMap[player.GetWorldId()] = make(map[playermodel.PlayerId]playermodel.Player, 0)
	}

	if _, exists := worldPlayerMap[player.GetWorldId()][player.GetId()]; exists {
		locker.Unlock()
		return fmt.Errorf("player already exists")
	}

	worldPlayerMap[player.GetWorldId()][player.GetId()] = player

	locker.Unlock()

	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Update(player playermodel.Player) error {
	locker.Lock()

	if _, exists := worldPlayerMap[player.GetWorldId()][player.GetId()]; !exists {
		locker.Unlock()
		return fmt.Errorf("player does not exists")
	}

	worldPlayerMap[player.GetWorldId()][player.GetId()] = player

	locker.Unlock()

	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Delete(player playermodel.Player) error {
	locker.Lock()

	if _, exists := worldPlayerMap[player.GetWorldId()][player.GetId()]; !exists {
		locker.Unlock()
		return fmt.Errorf("player does not exists")
	}

	delete(worldPlayerMap[player.GetWorldId()], player.GetId())

	locker.Unlock()

	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Get(worldId sharedkernelmodel.WorldId, playerId playermodel.PlayerId) (player playermodel.Player, err error) {
	locker.RLock()

	playerMap, found := worldPlayerMap[worldId]
	if !found {
		return player, fmt.Errorf("player not found")
	}

	player, found = playerMap[playerId]
	if !found {
		return player, fmt.Errorf("player not found")
	}

	locker.RUnlock()

	return player, nil
}

func (repo *playerRepo) GetPlayersOfWorld(worldId sharedkernelmodel.WorldId) ([]playermodel.Player, error) {
	locker.RLock()

	playerMap, found := worldPlayerMap[worldId]
	if !found {
		locker.RUnlock()
		return []playermodel.Player{}, nil
	}

	playersOfWorld := make([]playermodel.Player, 0)
	for _, player := range playerMap {
		playersOfWorld = append(playersOfWorld, player)
	}

	locker.RUnlock()

	return playersOfWorld, nil
}
