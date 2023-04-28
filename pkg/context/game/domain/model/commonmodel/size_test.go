package commonmodel

import (
	"testing"
)

func Test_NewSize(t *testing.T) {
	_, err := NewSize(-1, -1)
	if err == nil {
		t.Errorf("NewSize should return error when receiving negative width or height")
	}
}

func Test_Size_IsEqual(t *testing.T) {
	size1, _ := NewSize(10, 10)
	size2, _ := NewSize(10, 10)
	size3, _ := NewSize(10, 12)

	if !size1.IsEqual(size2) {
		t.Errorf("size1 is expected to be equal to size2")
	}
	if size1.IsEqual(size3) {
		t.Errorf("size1 is expected to be not equal to size3")
	}
}
