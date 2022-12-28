package main

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver"
)

func main() {
	var app string
	if app = os.Getenv("APP"); app == "" {
		panic("You must set the 'APP'")
	}

	if app == "apiserver" {
		apiserver.Start()
	} else if app == "gameserver" {
		gameserver.Start()
	}
}
