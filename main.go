package main

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/commonserver"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver"
)

func main() {
	var app string
	if app = os.Getenv("APP"); app == "" {
		panic("You must set the 'APP'")
	}

	if app == "common-server" {
		commonserver.Start()
	} else if app == "live-game-server" {
		livegameserver.Start()
	}
}
