package main

import (
	"flag"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/cli/clirouter"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httprouter"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		clirouter.Run(args)
		return
	}

	err := httprouter.Run()
	if err != nil {
		panic(err)
	}
}
