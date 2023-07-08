package main

import (
	"flag"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pgclient"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/redisclient"
	iam_mem_domain_event_handler "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/domainevent/memdomaineventhandler"
	world_mem_domain_event_handler "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/domainevent/memdomaineventhandler"
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
