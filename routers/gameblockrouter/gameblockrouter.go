package gameblockrouter

import (
	"fmt"
	"strconv"

	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gameblockmodel"
	"github.com/gin-gonic/gin"
)

func SetRouters(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		rowIdx, _ := strconv.Atoi((c.Query("rowIdx")))
		colIdx, _ := strconv.Atoi((c.Query("colIdx")))
		fmt.Println(rowIdx, colIdx)
		gameBlock, err := gameblockmodel.GetGameBlock(rowIdx, colIdx)
		if err != nil {
			c.JSON(404, "Not found.")
		}
		c.JSON(200, gameBlock)
	})
}
