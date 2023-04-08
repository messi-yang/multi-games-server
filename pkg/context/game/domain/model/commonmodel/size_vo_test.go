package commonmodel_test

import (
	"testing"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
)

func Test_NewSizeVo(t *testing.T) {
	_, err := commonmodel.NewSizeVo(-1, -1)
	if err == nil {
		t.Errorf("NewSizeVo should return error when receiving negative width or height")
	}
}
