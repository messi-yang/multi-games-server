package gamedao

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
)

func convertGameSizeToGameSize(gameSizeModel *valueobject.GameSize) *GameSize {
	return &GameSize{
		Width:  gameSizeModel.Width,
		Height: gameSizeModel.Height,
	}
}

func convertGameUnitMatrixToGameUnits(gameUnitsEntity *valueobject.GameUnitMatrix) *GameUnits {
	width := len(*gameUnitsEntity)
	gameField := make(GameUnits, width)
	for i := 0; i < width; i += 1 {
		height := len((*gameUnitsEntity)[i])
		gameField[i] = make([]GameUnit, height)
		for j := 0; j < height; j += 1 {

			gameField[i][j] = GameUnit{
				Alive: (*gameUnitsEntity)[i][j].Alive,
				Age:   (*gameUnitsEntity)[i][j].Age,
			}
		}
	}

	return &gameField
}
