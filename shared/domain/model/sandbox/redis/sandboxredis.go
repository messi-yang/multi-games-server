package redis

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
	Id        uuid.UUID     `json:"id"`
	UnitBlock [][]UnitModel `json:"unitBlock"`
}

func ConvertUnitBlockToUnitModelMatrix(unitBlock valueobject.UnitBlock) [][]UnitModel {
	unitMatrix := unitBlock.ToValueObjectMatrix()
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

func ConvertUnitBlockMatrixToUnitBlock(unitModelMatrix [][]UnitModel) valueobject.UnitBlock {
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

	return valueobject.NewUnitBlock(unitMatrix)
}

type sandboxRedis struct {
	redisService infrastructureservice.RedisService
}

type sandboxRedisConfiguration func(sandboxRedis *sandboxRedis) error

func NewSandboxRedis(cfgs ...sandboxRedisConfiguration) (*sandboxRedis, error) {
	sandboxRedis := &sandboxRedis{}
	for _, cfg := range cfgs {
		err := cfg(sandboxRedis)
		if err != nil {
			return nil, err
		}
	}
	return sandboxRedis, nil
}

func WithRedisService() sandboxRedisConfiguration {
	redisService := infrastructureservice.NewRedisService()
	return func(sandboxRedis *sandboxRedis) error {
		sandboxRedis.redisService = redisService
		return nil
	}
}

func (repository *sandboxRedis) createKey(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-id-%s", gameId)
}

func (repository *sandboxRedis) Add(game sandbox.Sandbox) error {
	dataKey := repository.createKey(game.GetId())
	newGameRecord := sandboxRecord{
		Id:        game.GetId(),
		UnitBlock: ConvertUnitBlockToUnitModelMatrix(game.GetUnitBlock()),
	}
	newGameRecordInBytes, _ := json.Marshal(newGameRecord)
	repository.redisService.Set(dataKey, newGameRecordInBytes)
	repository.redisService.Set("game-id", []byte(game.GetId().String()))

	return nil
}

func (repository *sandboxRedis) Get(id uuid.UUID) (sandbox.Sandbox, error) {
	dataKey := repository.createKey(id)
	gameFromRedisInBytes, _ := repository.redisService.Get(dataKey)
	var gameFromRedis sandboxRecord
	json.Unmarshal(gameFromRedisInBytes, &gameFromRedis)

	unitBlock := ConvertUnitBlockMatrixToUnitBlock(gameFromRedis.UnitBlock)
	game := sandbox.NewSandbox(gameFromRedis.Id, unitBlock)

	return game, nil
}

func (repository *sandboxRedis) GetFirstGameId() (uuid.UUID, error) {
	gameIdInBytes, _ := repository.redisService.Get("game-id")
	if len(gameIdInBytes) == 0 {
		return uuid.Nil, ErrGameNotFound
	}
	gameId, _ := uuid.ParseBytes(gameIdInBytes)

	return gameId, nil
}

func (repository *sandboxRedis) Update(id uuid.UUID, game sandbox.Sandbox) error {
	return nil
}
