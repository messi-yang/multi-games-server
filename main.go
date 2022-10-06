package main

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer"
)

func main() {
	var app string
	if app = os.Getenv("APP"); app == "" {
		panic("You must set the 'APP'")
	}

	if app == "client" {
		gameclientcommunicator.Start()
	} else if app == "computer" {
		gamecomputer.Start()
	}
}
