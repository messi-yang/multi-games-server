package routers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func InitializeRouters() {
	router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, game.GetCells())
	// })
	router.Run(":" + os.Getenv("PORT"))
}
