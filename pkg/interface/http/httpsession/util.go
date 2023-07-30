package httpsession

import (
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAuthorizedUserId(c *gin.Context) *uuid.UUID {
	authorizedUserIdDto, exits := c.Get("authorizedUserId")
	if !exits {
		return nil
	}
	return commonutil.ToPointer(authorizedUserIdDto.(uuid.UUID))
}

func SetAuthrorizedUserId(c *gin.Context, userId uuid.UUID) {
	c.Set("authorizedUserId", userId)
}
