package main

import (
	"flag"

	game_mem_domain_event_handler "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/event/mem/memdomaineventhandler"
	iam_mem_domain_event_handler "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/infrastructure/event/mem/memdomaineventhandler"
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

	game_mem_domain_event_handler.RegisterEvents()
	iam_mem_domain_event_handler.RegisterEvents()

	err := httprouter.Run()
	if err != nil {
		panic(err)
	}
}
