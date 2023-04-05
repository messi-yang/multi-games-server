package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/postgres"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/cli/seedclicontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/assethttpcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/gamesocketcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/itemhttpcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/worldhttpcontroller"
	"github.com/google/uuid"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		switch args[0] {
		case "db-seed":
			seedclicontroller.Exec()
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
	userIdDto, _ := uuid.Parse("d169faa5-c078-42c2-8a42-cd1d43558c7b")

	userId := usermodel.NewUserIdVo(userIdDto)
	newUser := usermodel.NewUserAgg(userId, "dumdumgenius@gmail.com", "DumDumGenius")
	err = userRepository.Add(newUser)
	if err != nil {
		fmt.Println(err)
	}

	assethttpcontroller.Setup(router)
	gamesocketcontroller.Setup(router)
	worldhttpcontroller.Setup(router)
	itemhttpcontroller.Setup(router)
	err = router.Run()
	if err != nil {
		panic(err)
	}
}
