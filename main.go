package main

import (
	"math/rand"

	"github.com/DumDumGeniuss/ggol"
	"github.com/gin-gonic/gin"
)

type CgolCell struct {
	Alive bool `json:"alive"`
}

func conwaysGameOfLifeNextAreaGenerator(
	coord *ggol.Coordinate,
	area *CgolCell,
	getAdjacentArea ggol.AdjacentAreaGetter[CgolCell],
) (nextArea *CgolCell) {
	newArea := *area

	var aliveAdjacentCellsCount int = 0
	for i := -1; i < 2; i += 1 {
		for j := -1; j < 2; j += 1 {
			if !(i == 0 && j == 0) {
				adjArea, _ := getAdjacentArea(coord, &ggol.Coordinate{X: i, Y: j})
				if adjArea.Alive {
					aliveAdjacentCellsCount += 1
				}
			}
		}
	}
	if newArea.Alive {
		if aliveAdjacentCellsCount != 2 && aliveAdjacentCellsCount != 3 {
			newArea.Alive = false
			return &newArea
		} else {
			newArea.Alive = true
			return &newArea
		}
	} else {
		if aliveAdjacentCellsCount == 3 {
			newArea.Alive = true
			return &newArea
		} else {
			newArea.Alive = false
			return &newArea
		}
	}
}

func main() {
	size := ggol.FieldSize{
		Width:  100,
		Height: 100,
	}
	initialCgolCell := CgolCell{
		Alive: false,
	}
	game, _ := ggol.NewGame(&size, &initialCgolCell)
	game.SetNextAreaGenerator(conwaysGameOfLifeNextAreaGenerator)

	for i := 0; i < size.Width; i += 1 {
		for j := 0; j < size.Height; j += 1 {
			game.SetArea(&ggol.Coordinate{X: i, Y: j}, &CgolCell{Alive: rand.Intn(2) == 0})
		}
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		game.GenerateNextField()
		c.JSON(200, game.GetField())
	})
	r.Run()
}
