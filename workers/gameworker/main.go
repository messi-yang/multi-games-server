package gameworker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/ggol"
)

var game ggol.Game[GameUnit]
var gameTickerTimeout chan bool
var gameBlockChangeEventSubscribers []gameBlockChangeEventSubscriber

func Initialize() {
	if game != nil {
		return
	}

	// gameDao := gamedao.GetGameDao()
	// gameField, err := gameDao.GetGameField()
	// gameFieldSize, err := gameDao.GetGameFieldSize()

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
	gameBlockChangeEventSubscribers = []gameBlockChangeEventSubscriber{}
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
	key string,
	gameBlockArea GameBlockArea,
	callback gameBlockChangeEventCallback,
) {
	UnsubscribeGameBlockChangeEvent(key)

	newSubscriber := gameBlockChangeEventSubscriber{
		key:              key,
		gameBlockArea:    gameBlockArea,
		gameUnitsChannel: make(chan gameUnits),
	}
	gameBlockChangeEventSubscribers = append(
		gameBlockChangeEventSubscribers,
		newSubscriber,
	)

	go func() {
		for {
			select {
			case gameUnits := <-newSubscriber.gameUnitsChannel:
				callback(gameUnits)
			}
		}
	}()
}

func UnsubscribeGameBlockChangeEvent(key string) {
	for listenerIdx, listener := range gameBlockChangeEventSubscribers {
		if listener.key == key {
			gameBlockChangeEventSubscribers =
				append(
					gameBlockChangeEventSubscribers[:listenerIdx],
					gameBlockChangeEventSubscribers[listenerIdx+1:]...,
				)
		}
	}
}

func publishGameBlockChangeEvent() {
	for _, subscriber := range gameBlockChangeEventSubscribers {
		gameUnits, err := game.GetUnitsInArea(&ggol.Area{
			From: ggol.Coordinate{X: subscriber.gameBlockArea.FromX, Y: subscriber.gameBlockArea.FromY},
			To:   ggol.Coordinate{X: subscriber.gameBlockArea.ToX, Y: subscriber.gameBlockArea.ToY},
		})
		if err != nil {
			fmt.Println(err.Error())
			UnsubscribeGameBlockChangeEvent(subscriber.key)
			return
		}
		subscriber.gameUnitsChannel <- gameUnits
	}
}

func Stop() {
	gameTickerTimeout <- true
}
