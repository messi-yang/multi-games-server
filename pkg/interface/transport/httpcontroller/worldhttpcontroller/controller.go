package worldhttpcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/worldappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/transport/httpcontroller"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (controller *Controller) QueryHandler(c *gin.Context) {
	presenter := httpcontroller.NewPresenter(c)
	worldRepository, err := postgres.NewWorldRepository()
	if err != nil {
		presenter.OnError(err)
	}
	worldAppService := worldappservice.NewService(worldRepository, presenter)

	worldAppService.QueryWorlds()
}
