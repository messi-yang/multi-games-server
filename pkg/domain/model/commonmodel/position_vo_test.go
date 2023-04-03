package commonmodel_test

import (
	"testing"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
)

func Test_PositionVo_IsEqual(t *testing.T) {
	pos1 := commonmodel.NewPositionVo(0, 0)
	pos2 := commonmodel.NewPositionVo(0, 0)
	pos3 := commonmodel.NewPositionVo(1, 1)

	if !pos1.IsEqual(pos2) {
		t.Errorf("pos1 is expected to be equal to pos2")
	}
	if pos1.IsEqual(pos3) {
		t.Errorf("pos1 is expected to be not equal to pos3")
	}
}
