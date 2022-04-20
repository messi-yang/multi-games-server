package gameentity

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/ggol"
)

var game ggol.Game[GameUnit]
var gameTickerTimeout chan bool
var gameBlockChangeEventListners []gameBlockChangeEventListner

func Initialize() {
	size := ggol.Size{
		Width:  config.Config.GAME_SIZE,
		Height: config.Config.GAME_SIZE,
	}
	initialGameUnit := GameUnit{
		Alive: false,
	}
	newGame, err := ggol.NewGame(&size, &initialGameUnit)
	if err != nil {
		panic(err)
	}
	newGame.SetNextUnitGenerator(gameOfLifeNextUnitGenerator)
	newGame.IterateUnits(func(coord *ggol.Coordinate, _ *GameUnit) {
		newGame.SetUnit(coord, &GameUnit{Alive: rand.Intn(2) == 0})
	})
	game = newGame
	gameBlockChangeEventListners = []gameBlockChangeEventListner{}
}

func Start() {
	gameTicker := time.NewTicker(time.Millisecond * 1000)
	gameTickerTimeout = make(chan bool, 1)
	go func() {
		for {
			select {
			case <-gameTicker.C:
				game.GenerateNextUnits()
				publishGameBlockChangeEvent()
			case <-gameTickerTimeout:
				gameTicker.Stop()
				return
			}
		}
	}()
}

func SubscribeGameBlockChangeEvent(
	fromX int,
	fromY int,
	toX int,
	toY int,
	callback gameBlockChangeEventCallback,
) chan bool {
	stopChannel := make(chan bool)
	newListener := gameBlockChangeEventListner{
		fromX:            fromX,
		fromY:            fromY,
		toX:              toX,
		toY:              toY,
		gameUnitsChannel: make(chan gameUnits),
		stopChannel:      stopChannel,
	}
	gameBlockChangeEventListners = append(
		gameBlockChangeEventListners,
		newListener,
	)

	unsubscribe := func() {
		for listenerIdx, listener := range gameBlockChangeEventListners {
			if listener == newListener {
				gameBlockChangeEventListners =
					append(
						gameBlockChangeEventListners[:listenerIdx],
						gameBlockChangeEventListners[listenerIdx+1:]...,
					)
			}
		}
	}

	go func() {
		for {
			select {
			case gameUnits := <-newListener.gameUnitsChannel:
				callback(gameUnits)
			case <-newListener.stopChannel:
				unsubscribe()
				return
			}
		}
	}()

	return stopChannel
}

func Stop() {
	gameTickerTimeout <- true
}

func publishGameBlockChangeEvent() {
	for _, listner := range gameBlockChangeEventListners {
		gameUnits, err := game.GetUnitsInArea(&ggol.Area{
			From: ggol.Coordinate{X: listner.fromX, Y: listner.fromY},
			To:   ggol.Coordinate{X: listner.toX, Y: listner.toY},
		})
		if err != nil {
			fmt.Println(err.Error())
			listner.stopChannel <- true
			return
		}
		listner.gameUnitsChannel <- gameUnits
	}
}
