package gamedao

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/repository"
)

type GameDAO interface {
	InjectGameRepository(repository.GameRepository)
	GetGameUnits() (*GameUnits, error)
	GetGameSize() (*GameSize, error)
}

type gamDAOImplement struct {
	gameModel repository.GameRepository
}

var gameDAO GameDAO

func GetGameDAO() GameDAO {
	if gameDAO == nil {
		gameDAO = &gamDAOImplement{}
	}

	return gameDAO
}

func (gdi *gamDAOImplement) InjectGameRepository(gameModel repository.GameRepository) {
	gdi.gameModel = gameModel
}

func (gdi *gamDAOImplement) GetGameUnits() (*GameUnits, error) {
	gameUnitsEntity := gdi.gameModel.GetGameUnitMatrix()
	gameField := convertGameUnitMatrixToGameUnits(gameUnitsEntity)

	return gameField, nil
}

func (gdi *gamDAOImplement) GetGameSize() (*GameSize, error) {
	gameSizeModel := gdi.gameModel.GetGameSize()
	gameFieldSize := convertGameSizeToGameSize(gameSizeModel)

	return gameFieldSize, nil
}
