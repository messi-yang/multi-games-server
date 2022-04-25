package gamedao

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gamemodel"
)

type GameDAO interface {
	InjectGameModel(gamemodel.GameModel)
	GetGameUnits() (*GameUnits, error)
	GetGameSize() (*GameSize, error)
}

type gamDAOImplement struct {
	gameModel gamemodel.GameModel
}

var gameDAO GameDAO

func GetGameDAO() GameDAO {
	if gameDAO == nil {
		gameDAO = &gamDAOImplement{}
	}

	return gameDAO
}

func (gdi *gamDAOImplement) InjectGameModel(gameModel gamemodel.GameModel) {
	gdi.gameModel = gameModel
}

func (gdi *gamDAOImplement) GetGameUnits() (*GameUnits, error) {
	gameUnitsEntity := gdi.gameModel.GetGameUnitsModel()
	gameField := convertGameUnitsModelToGameUnits(gameUnitsEntity)

	return gameField, nil
}

func (gdi *gamDAOImplement) GetGameSize() (*GameSize, error) {
	gameSizeModel := gdi.gameModel.GetGameSizeModel()
	gameFieldSize := convertGameSizeModelToGameSize(gameSizeModel)

	return gameFieldSize, nil
}
