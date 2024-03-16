package main

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pgclient"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/redisclient"
	world_mem_domain_event_handler "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/domainevent/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httprouter"
	"github.com/google/uuid"
)

func main() {
	pgclient.Connect()
	redisclient.Connect()

	world_mem_domain_event_handler.RegisterEvents()

	fmt.Println(uuid.New())

	err := httprouter.Run()
	if err != nil {
		panic(err)
	}
}
