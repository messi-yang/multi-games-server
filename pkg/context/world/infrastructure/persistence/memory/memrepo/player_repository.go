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
	repo.mutex.Lock()

	_, found := worldPlayerMap[player.GetWorldId()]
	if !found {
		worldPlayerMap[player.GetWorldId()] = make(map[playermodel.PlayerId]playermodel.Player, 0)
	}

	if _, exists := worldPlayerMap[player.GetWorldId()][player.GetId()]; exists {
		repo.mutex.Unlock()
		return fmt.Errorf("player already exists")
	}

	worldPlayerMap[player.GetWorldId()][player.GetId()] = player

	repo.mutex.Unlock()

	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Update(player playermodel.Player) error {
	repo.mutex.Lock()

	if _, exists := worldPlayerMap[player.GetWorldId()][player.GetId()]; !exists {
		repo.mutex.Unlock()
		return fmt.Errorf("player does not exists")
	}

	worldPlayerMap[player.GetWorldId()][player.GetId()] = player

	repo.mutex.Unlock()

	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Delete(player playermodel.Player) error {
	repo.mutex.Lock()

	if _, exists := worldPlayerMap[player.GetWorldId()][player.GetId()]; !exists {
		repo.mutex.Unlock()
		return fmt.Errorf("player does not exists")
	}

	delete(worldPlayerMap[player.GetWorldId()], player.GetId())

	repo.mutex.Unlock()

	return repo.domainEventDispatcher.Dispatch(&player)
}

func (repo *playerRepo) Get(playerId playermodel.PlayerId) (player playermodel.Player, err error) {
	repo.mutex.RLock()

	for _, worldPlayers := range worldPlayerMap {
		for _, player := range worldPlayers {
			if player.GetId().IsEqual(playerId) {
				repo.mutex.RUnlock()
				return player, nil
			}
		}
	}

	repo.mutex.RUnlock()

	return player, fmt.Errorf("player not found")
}

func (repo *playerRepo) GetPlayersAt(worldId sharedkernelmodel.WorldId, position commonmodel.Position) ([]playermodel.Player, error) {
	repo.mutex.RLock()

	playerMap, found := worldPlayerMap[worldId]
	if !found {
		repo.mutex.RUnlock()
		return []playermodel.Player{}, nil
	}

	playersAtPosition := make([]playermodel.Player, 0)
	for _, player := range playerMap {
		if player.GetPosition().IsEqual(position) {
			playersAtPosition = append(playersAtPosition, player)
		}
	}

	repo.mutex.RUnlock()

	return playersAtPosition, nil
}

func (repo *playerRepo) GetPlayersOfWorld(worldId sharedkernelmodel.WorldId) ([]playermodel.Player, error) {
	repo.mutex.RLock()

	playerMap, found := worldPlayerMap[worldId]
	if !found {
		repo.mutex.RUnlock()
		return []playermodel.Player{}, nil
	}

	playersOfWorld := make([]playermodel.Player, 0)
	for _, player := range playerMap {
		playersOfWorld = append(playersOfWorld, player)
	}

	repo.mutex.RUnlock()

	return playersOfWorld, nil
}

func (repo *playerRepo) GetAll(worldId sharedkernelmodel.WorldId) []playermodel.Player {
	repo.mutex.RLock()

	allPlayers := make([]playermodel.Player, 0)
	for _, worldPlayers := range worldPlayerMap {
		for _, player := range worldPlayers {
			allPlayers = append(allPlayers, player)
		}
	}

	repo.mutex.RUnlock()

	return allPlayers
}
