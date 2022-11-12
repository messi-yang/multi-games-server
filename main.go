package main

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver"
	"github.com/dum-dum-genius/game-of-liberty-computer/mainserver"
)

func main() {
	var app string
	if app = os.Getenv("APP"); app == "" {
		panic("You must set the 'APP'")
	}

	if app == "main-server" {
		mainserver.Start()
	} else if app == "live-game-server" {
		livegameserver.Start()
	}
}
