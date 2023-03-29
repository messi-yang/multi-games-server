package worldappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
)

type CreateWorldCommand struct {
	UserId usermodel.UserIdVo
}
