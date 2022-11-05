package redisrepository

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/sandbox"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
	"github.com/google/uuid"
)

var (
	ErrGameNotFound      = errors.New("the game room with the id was not found")
	ErrGameAlreadyExists = errors.New("the game room with same id already exists")
)

type UnitModel struct {
	Alive bool `json:"alive"`
}

type sandboxRecord struct {
	Id      uuid.UUID     `json:"id"`
	UnitMap [][]UnitModel `json:"unitMap"`
}

func ConvertUnitMapToUnitModelMatrix(unitMap valueobject.UnitMap) [][]UnitModel {
	unitMatrix := unitMap.ToValueObjectMatrix()
	unitModelMatrix := make([][]UnitModel, 0)
	for colIdx, unitMatrixCol := range unitMatrix {
		unitModelMatrix = append(unitModelMatrix, make([]UnitModel, 0))
		for _, unit := range unitMatrixCol {
			unitModelMatrix[colIdx] = append(unitModelMatrix[colIdx], UnitModel{
				Alive: unit.GetAlive(),
			})
		}
	}

	return unitModelMatrix
}

func ConvertUnitMapMatrixToUnitMap(unitModelMatrix [][]UnitModel) valueobject.UnitMap {
	unitMatrix := make([][]valueobject.Unit, 0)
	for colIdx, unitModelMatrixCol := range unitModelMatrix {
		unitMatrix = append(unitMatrix, make([]valueobject.Unit, 0))
		for _, unitModel := range unitModelMatrixCol {
			unitMatrix[colIdx] = append(unitMatrix[colIdx], valueobject.NewUnit(
				unitModel.Alive,
				uuid.Nil,
			))
		}
	}

	return valueobject.NewUnitMap(unitMatrix)
}

type sandboxRedisRepository struct {
	redisInfrastructureService infrastructureservice.RedisInfrastructureService
}

type SandboxRedisRepositoryConfiguration struct {
	RedisInfrastructureService infrastructureservice.RedisInfrastructureService
}

func NewSandboxRedisRepository(configuration SandboxRedisRepositoryConfiguration) sandbox.Repository {
	return &sandboxRedisRepository{
		redisInfrastructureService: configuration.RedisInfrastructureService,
	}
}

func (repository *sandboxRedisRepository) createKey(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-id-%s", gameId)
}

func (repository *sandboxRedisRepository) Add(game sandbox.Sandbox) error {
	dataKey := repository.createKey(game.GetId())
	newGameRecord := sandboxRecord{
		Id:      game.GetId(),
		UnitMap: ConvertUnitMapToUnitModelMatrix(game.GetUnitMap()),
	}
	newGameRecordInBytes, _ := json.Marshal(newGameRecord)
	repository.redisInfrastructureService.Set(dataKey, newGameRecordInBytes)
	repository.redisInfrastructureService.Set("game-id", []byte(game.GetId().String()))

	return nil
}

func (repository *sandboxRedisRepository) Get(id uuid.UUID) (sandbox.Sandbox, error) {
	dataKey := repository.createKey(id)
	gameFromRedisInBytes, _ := repository.redisInfrastructureService.Get(dataKey)
	var gameFromRedis sandboxRecord
	json.Unmarshal(gameFromRedisInBytes, &gameFromRedis)

	unitMap := ConvertUnitMapMatrixToUnitMap(gameFromRedis.UnitMap)
	game := sandbox.NewSandbox(gameFromRedis.Id, unitMap)

	return game, nil
}

func (repository *sandboxRedisRepository) GetFirstGameId() (uuid.UUID, error) {
	gameIdInBytes, _ := repository.redisInfrastructureService.Get("game-id")
	if len(gameIdInBytes) == 0 {
		return uuid.Nil, ErrGameNotFound
	}
	gameId, _ := uuid.ParseBytes(gameIdInBytes)

	return gameId, nil
}

func (repository *sandboxRedisRepository) Update(id uuid.UUID, game sandbox.Sandbox) error {
	return nil
}
