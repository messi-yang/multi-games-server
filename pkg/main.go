package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/cmd/seedcmd"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/transport/api/gameapi"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/transport/api/worldapi"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		switch args[0] {
		case "db-seed":
			seedcmd.Exec()
			os.Exit(0)
		}
		return
	}

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	userRepository, err := postgres.NewUserRepository()
	if err != nil {
		panic(err)
	}
	userId, _ := usermodel.ParseUserIdVo("d169faa5-c078-42c2-8a42-cd1d43558c7b")
	newUser := usermodel.NewUnitAgg(userId, "dumdumgenius@gmail.com", "DumDumGenius")
	err = userRepository.Add(newUser)
	if err != nil {
		fmt.Println(err)
	}

	router.Static("/asset", "./pkg/interface/transport/asset")

	gameapi.SetRouter(router.Group("/ws/game"))
	worldapi.SetRouter(router.Group("/api/worlds"))
	err = router.Run()
	if err != nil {
		panic(err)
	}
}
