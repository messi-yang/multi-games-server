package gamedao

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gamemodel"
)

type GameDAO interface {
	GetGameField() (*GameField, error)
	GetGameSize() (*GameSize, error)
}

type gamDAOImplement struct {
	gameModel gamemodel.GameModel
}

var gameDAO GameDAO

func GetGameDAO(gameModel gamemodel.GameModel) GameDAO {
	if gameDAO == nil {
		gameDAO = &gamDAOImplement{
			gameModel: gameModel,
		}
		return gameDAO
	} else {
		return gameDAO
	}
}

func (gdi *gamDAOImplement) GetGameField() (*GameField, error) {
	gameUnitsEntity := gdi.gameModel.GetGameUnitsModel()
	gameField := convertGameUnitsModelToGameField(gameUnitsEntity)

	return gameField, nil
}

func (gdi *gamDAOImplement) GetGameSize() (*GameSize, error) {
	gameSizeModel := gdi.gameModel.GetGameSizeModel()
	gameFieldSize := convertGameSizeModelToGameFieldSize(gameSizeModel)

	return gameFieldSize, nil
}
