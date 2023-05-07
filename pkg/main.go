package main

import (
	"flag"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/memory/memdomaineventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/redis/redisclient"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgclient"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/cli/clirouter"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httprouter"
)

func main() {
	pgclient.Connect()
	redisclient.Connect()

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		clirouter.Run(args)
		return
	}

	memdomaineventhandler.Run()

	err := httprouter.Run()
	if err != nil {
		panic(err)
	}
}
