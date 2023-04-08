package commonmodel

import (
	"testing"
)

func Test_NewSizeVo(t *testing.T) {
	_, err := NewSizeVo(-1, -1)
	if err == nil {
		t.Errorf("NewSizeVo should return error when receiving negative width or height")
	}
}
