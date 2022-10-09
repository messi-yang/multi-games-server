package redisrepository

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/repository"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
	"github.com/google/uuid"
)

var (
	ErrGameRoomNotFound      = errors.New("the game room with the id was not found")
	ErrGameRoomAlreadyExists = errors.New("the game room with same id already exists")
)

type UnitModel struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

type gameRoomRecord struct {
	Id      uuid.UUID     `json:"id"`
	UnitMap [][]UnitModel `json:"unitMap"`
}

func ConvertUnitMapToModelMatrix(unitMap *valueobject.UnitMap) [][]UnitModel {
	unitMatrix := unitMap.ToValueObjectMatrix()
	unitModelMatrix := make([][]UnitModel, 0)
	for colIdx, unitMatrixCol := range *unitMatrix {
		unitModelMatrix = append(unitModelMatrix, make([]UnitModel, 0))
		for _, unit := range unitMatrixCol {
			unitModelMatrix[colIdx] = append(unitModelMatrix[colIdx], UnitModel{
				Alive: unit.GetAlive(),
				Age:   unit.GetAge(),
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
				unitModel.Age,
			))
		}
	}

	return valueobject.NewUnitMapFromUnitMatrix(&unitMatrix)
}

type gameRoomPersistentRedisRepository struct {
	redisInfrastructureService infrastructureservice.RedisInfrastructureService
}

type GameRoomPersistentRedisRepositoryConfiguration struct {
	RedisInfrastructureService infrastructureservice.RedisInfrastructureService
}

func NewGameRoomPersistentRedisRepository(configuration GameRoomPersistentRedisRepositoryConfiguration) repository.GameRoomPersistentRepository {
	return &gameRoomPersistentRedisRepository{
		redisInfrastructureService: configuration.RedisInfrastructureService,
	}
}

func (repository *gameRoomPersistentRedisRepository) createKey(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-id-%s", gameId)
}

func (repository *gameRoomPersistentRedisRepository) Add(gameRoom aggregate.GameRoom) error {
	redisDataKey := repository.createKey(gameRoom.GetId())
	newGameRoomModel := gameRoomRecord{
		Id:      gameRoom.GetId(),
		UnitMap: ConvertUnitMapToModelMatrix(gameRoom.GetUnitMap()),
	}
	newGameRoomModelInBytes, _ := json.Marshal(newGameRoomModel)
	repository.redisInfrastructureService.Set(redisDataKey, newGameRoomModelInBytes)
	fmt.Println(gameRoom.GetId())

	return nil
}

func (repository *gameRoomPersistentRedisRepository) Get(id uuid.UUID) (aggregate.GameRoom, error) {
	redisDataKey := repository.createKey(id)
	goomModelFromRedisInBytes, _ := repository.redisInfrastructureService.Get(redisDataKey)
	var goomModelFromRedis gameRoomRecord
	json.Unmarshal(goomModelFromRedisInBytes, &goomModelFromRedis)

	unitMap := ConvertUnitMapMatrixToUnitMap(goomModelFromRedis.UnitMap)
	game := entity.LoadGame(goomModelFromRedis.Id, unitMap, 1000)
	gameRoom := aggregate.NewGameRoom(game)

	return gameRoom, nil
}

func (repository *gameRoomPersistentRedisRepository) Update(id uuid.UUID, gameRoom aggregate.GameRoom) error {
	return nil
}
