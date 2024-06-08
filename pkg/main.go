package main

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/application/domaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pgclient"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/redisclient"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httprouter"
	"github.com/google/uuid"
)

func main() {
	pgclient.Connect()
	redisclient.Connect()

	domaineventhandler.RegisterEvents()

	fmt.Println(uuid.New())

	err := httprouter.Run()
	if err != nil {
		panic(err)
	}
}
