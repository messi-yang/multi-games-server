package gamedao

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/repository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
)

type GameDAO interface {
	InjectGameRepository(repository.GameRepository)
	GetGameUnits() (*valueobject.GameUnitMatrix, error)
	GetGameSize() (*valueobject.GameSize, error)
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

func (gdi *gamDAOImplement) GetGameUnits() (*valueobject.GameUnitMatrix, error) {
	gameUnitsEntity := gdi.gameModel.GetGameUnitMatrix()

	return gameUnitsEntity, nil
}

func (gdi *gamDAOImplement) GetGameSize() (*valueobject.GameSize, error) {
	gameSizeModel := gdi.gameModel.GetGameSize()

	return gameSizeModel, nil
}
