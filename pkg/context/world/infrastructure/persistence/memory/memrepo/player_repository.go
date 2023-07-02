package memrepo

import (
	"fmt"
	"sync"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/playermodel"
)

type playerRepo struct {
	mutex                 sync.RWMutex
	domainEventDispatcher domain.DomainEventDispatcher
}

var worldPlayerMap = make(map[sharedkernelmodel.WorldId]map[playermodel.PlayerId]playermodel.Player)

func NewPlayerRepo(domainEventDispatcher domain.DomainEventDispatcher) (repository playermodel.PlayerRepo) {
	return &playerRepo{
		mutex:                 sync.RWMutex{},
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *playerRepo) Add(player playermodel.Player) error {
	// repo.mutex.Lock()
	// defer repo.mutex.Unlock()

	_, found := worldPlayerMap[player.GetWorldId()]
	if !found {
		worldPlayerMap[player.GetWorldId()] = make(map[playermodel.PlayerId]playermodel.Player, 0)
	}

	if _, exists := worldPlayerMap[player.GetWorldId()][player.GetId()]; exists {
		return fmt.Errorf("player already exists")
	}

	worldPlayerMap[player.GetWorldId()][player.GetId()] = player
	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Update(player playermodel.Player) error {
	// repo.mutex.Lock()
	// defer repo.mutex.Unlock()

	if _, exists := worldPlayerMap[player.GetWorldId()][player.GetId()]; !exists {
		return fmt.Errorf("player does not exists")
	}

	worldPlayerMap[player.GetWorldId()][player.GetId()] = player
	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Delete(player playermodel.Player) error {
	// repo.mutex.Lock()
	// defer repo.mutex.Unlock()

	if _, exists := worldPlayerMap[player.GetWorldId()][player.GetId()]; !exists {
		return fmt.Errorf("player does not exists")
	}

	delete(worldPlayerMap[player.GetWorldId()], player.GetId())
	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Get(playerId playermodel.PlayerId) (player playermodel.Player, err error) {
	// repo.mutex.RLock()
	// defer repo.mutex.RUnlock()

	for _, worldPlayers := range worldPlayerMap {
		for _, player := range worldPlayers {
			if player.GetId().IsEqual(playerId) {
				return player, nil
			}
		}
	}
	return player, fmt.Errorf("player not found")
}

func (repo *playerRepo) GetPlayersAt(worldId sharedkernelmodel.WorldId, position commonmodel.Position) ([]playermodel.Player, error) {
	// repo.mutex.RLock()
	// defer repo.mutex.RUnlock()

	playerMap, found := worldPlayerMap[worldId]
	if !found {
		return []playermodel.Player{}, nil
	}

	playersAtPosition := make([]playermodel.Player, 0)
	for _, player := range playerMap {
		if player.GetPosition().IsEqual(position) {
			playersAtPosition = append(playersAtPosition, player)
		}
	}
	return playersAtPosition, nil
}

func (repo *playerRepo) GetPlayersOfWorld(worldId sharedkernelmodel.WorldId) ([]playermodel.Player, error) {
	// repo.mutex.RLock()
	// defer repo.mutex.RUnlock()

	playerMap, found := worldPlayerMap[worldId]
	if !found {
		return []playermodel.Player{}, nil
	}

	playersOfWorld := make([]playermodel.Player, 0)
	for _, player := range playerMap {
		playersOfWorld = append(playersOfWorld, player)
	}
	return playersOfWorld, nil
}

func (repo *playerRepo) GetAll(worldId sharedkernelmodel.WorldId) []playermodel.Player {
	// repo.mutex.RLock()
	// defer repo.mutex.RUnlock()

	allPlayers := make([]playermodel.Player, 0)
	for _, worldPlayers := range worldPlayerMap {
		for _, player := range worldPlayers {
			allPlayers = append(allPlayers, player)
		}
	}
	return allPlayers
}
