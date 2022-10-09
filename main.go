package main

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer"
)

func main() {
	var app string
	if app = os.Getenv("APP"); app == "" {
		panic("You must set the 'APP'")
	}

	if app == "client" {
		gameclient.Start()
	} else if app == "computer" {
		gamecomputer.Start()
	}
}
