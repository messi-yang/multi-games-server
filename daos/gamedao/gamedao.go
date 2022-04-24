package gamedao

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gamemodel"
)

type GameDAO interface {
	GetGameField() (*GameField, error)
	GetGameFieldSize() (*GameFieldSize, error)
}

type gamDAOImplement struct {
	gameModel gamemodel.GameModel
}

func CreateGameDAO(gameModel gamemodel.GameModel) GameDAO {
	return &gamDAOImplement{
		gameModel: gameModel,
	}
}

var DAO GameDAO = &gamDAOImplement{}

func (gdi *gamDAOImplement) GetGameField() (*GameField, error) {
	gameUnitsEntity := gdi.gameModel.GetGameUnitsModel()
	gameField := convertGameUnitsModelToGameField(gameUnitsEntity)

	return gameField, nil
}

func (gdi *gamDAOImplement) GetGameFieldSize() (*GameFieldSize, error) {
	gameSizeModel := gdi.gameModel.GetGameSizeModel()
	gameFieldSize := convertGameSizeModelToGameFieldSize(gameSizeModel)

	return gameFieldSize, nil
}
