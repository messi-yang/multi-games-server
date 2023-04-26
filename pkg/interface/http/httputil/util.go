package httputil

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserId(c *gin.Context) uuid.UUID {
	userIdDto, _ := c.Get("userId")
	return userIdDto.(uuid.UUID)
}

func SetUserId(c *gin.Context, userId uuid.UUID) {
	c.Set("userId", userId)
}
