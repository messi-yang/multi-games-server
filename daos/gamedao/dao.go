package gamedao

type GameDAO interface {
	GetGameField() (*GameField, error)
	GetGameFieldSize() (*GameFieldSize, error)
}

type gamDAOImplement struct {
}

var DAO GameDAO = &gamDAOImplement{}

func (d *gamDAOImplement) GetGameField() (*GameField, error) {
	gameFieldEntity := getGameFieldEntityFromStorage()
	gameField := convertGameFieldEntityToGameField(gameFieldEntity)

	return gameField, nil
}

func (d *gamDAOImplement) GetGameFieldSize() (*GameFieldSize, error) {
	gameFieldSizeEntity := getGameFieldSizeEntityFromStorage()
	gameFieldSize := convertGameFieldSizeEntityToGameFieldSize(gameFieldSizeEntity)

	return gameFieldSize, nil
}
