package redisrepository

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/repository"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
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

type gameRecord struct {
	Id      uuid.UUID     `json:"id"`
	UnitMap [][]UnitModel `json:"unitMap"`
}

func ConvertUnitMapToUnitModelMatrix(unitMap *valueobject.UnitMap) [][]UnitModel {
	unitMatrix := unitMap.ToValueObjectMatrix()
	unitModelMatrix := make([][]UnitModel, 0)
	for colIdx, unitMatrixCol := range *unitMatrix {
		unitModelMatrix = append(unitModelMatrix, make([]UnitModel, 0))
		for _, unit := range unitMatrixCol {
			unitModelMatrix[colIdx] = append(unitModelMatrix[colIdx], UnitModel{
				Alive: unit.GetAlive(),
			})
		}
	}

	return unitModelMatrix
}

func ConvertUnitMapMatrixToUnitMap(unitModelMatrix [][]UnitModel) *valueobject.UnitMap {
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

	return valueobject.NewUnitMapFromUnitMatrix(&unitMatrix)
}

type gameRedisRepository struct {
	redisInfrastructureService infrastructureservice.RedisInfrastructureService
}

type GameRedisRepositoryConfiguration struct {
	RedisInfrastructureService infrastructureservice.RedisInfrastructureService
}

func NewGameRedisRepository(configuration GameRedisRepositoryConfiguration) repository.GameRepository {
	return &gameRedisRepository{
		redisInfrastructureService: configuration.RedisInfrastructureService,
	}
}

func (repository *gameRedisRepository) createKey(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-id-%s", gameId)
}

func (repository *gameRedisRepository) Add(game entity.Game) error {
	dataKey := repository.createKey(game.GetId())
	newGameRecord := gameRecord{
		Id:      game.GetId(),
		UnitMap: ConvertUnitMapToUnitModelMatrix(game.GetUnitMap()),
	}
	newGameRecordInBytes, _ := json.Marshal(newGameRecord)
	repository.redisInfrastructureService.Set(dataKey, newGameRecordInBytes)
	repository.redisInfrastructureService.Set("game-id", []byte(game.GetId().String()))

	return nil
}

func (repository *gameRedisRepository) Get(id uuid.UUID) (entity.Game, error) {
	dataKey := repository.createKey(id)
	gameFromRedisInBytes, _ := repository.redisInfrastructureService.Get(dataKey)
	var gameFromRedis gameRecord
	json.Unmarshal(gameFromRedisInBytes, &gameFromRedis)

	unitMap := ConvertUnitMapMatrixToUnitMap(gameFromRedis.UnitMap)
	game := entity.NewGame(gameFromRedis.Id, unitMap)

	return game, nil
}

func (repository *gameRedisRepository) GetFirstGameId() (uuid.UUID, error) {
	gameIdInBytes, _ := repository.redisInfrastructureService.Get("game-id")
	if len(gameIdInBytes) == 0 {
		return uuid.Nil, ErrGameNotFound
	}
	gameId, _ := uuid.ParseBytes(gameIdInBytes)

	return gameId, nil
}

func (repository *gameRedisRepository) Update(id uuid.UUID, game entity.Game) error {
	return nil
}
