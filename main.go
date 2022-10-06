package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer"
)

func main() {
	go gamecomputer.Start()
	gameclientcommunicator.Start()
}
