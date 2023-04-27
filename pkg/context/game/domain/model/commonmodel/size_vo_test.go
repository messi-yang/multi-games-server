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

func Test_SizeVo_IsEqual(t *testing.T) {
	size1, _ := NewSizeVo(10, 10)
	size2, _ := NewSizeVo(10, 10)
	size3, _ := NewSizeVo(10, 12)

	if !size1.IsEqual(size2) {
		t.Errorf("size1 is expected to be equal to size2")
	}
	if size1.IsEqual(size3) {
		t.Errorf("size1 is expected to be not equal to size3")
	}
}
