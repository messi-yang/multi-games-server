package clirouter

import (
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/cli/clihandler/seedclihandler"
)

func Run(args []string) {
	switch args[0] {
	case "db-seed":
		seedCliHandler := seedclihandler.NewHandler()
		seedCliHandler.Exec()
		os.Exit(0)
	}
}
