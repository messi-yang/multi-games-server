package gamedao

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gamemodel"
)

func convertGameFieldSizeEntityToGameFieldSize(gameFieldSizeEntity *gamemodel.GameFieldSizeEntity) *GameFieldSize {
	return &GameFieldSize{
		Width:  gameFieldSizeEntity.Width,
		Height: gameFieldSizeEntity.Height,
	}
}

func convertGameFieldEntityToGameField(gameFieldEntity *gamemodel.GameFieldEntity) *GameField {
	width := len(*gameFieldEntity)
	gameField := make(GameField, width)
	for i := 0; i < width; i += 1 {
		height := len((*gameFieldEntity)[i])
		gameField[i] = make([]GameFieldUnit, height)
		for j := 0; j < height; j += 1 {

			gameField[i][j] = GameFieldUnit{
				Alive: (*gameFieldEntity)[i][j].Alive,
				Age:   (*gameFieldEntity)[i][j].Age,
			}
		}
	}

	return &gameField
}
