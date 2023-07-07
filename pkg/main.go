package main

import (
	"flag"

	iam_mem_domain_event_handler "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/event/mem/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/messaging/redis/redisclient"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgclient"
	world_mem_domain_event_handler "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/event/mem/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/cli/clirouter"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httprouter"
)

func main() {
	pgclient.Connect()
	redisclient.Connect()

	world_mem_domain_event_handler.RegisterEvents()
	iam_mem_domain_event_handler.RegisterEvents()

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
