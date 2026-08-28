package ndarray_test

import (
	"testing"
)

func TestNewNDArray(t *testing.T) {
	// 2x3 matrix filled with 0.0
	shape := []int{2, 3}
	arr := New[float64](shape, 0.0)

	if arr == nil {
		t.Fatal("expected non-nil NDArray")
	}

	// Verify initial element access
	val := arr.At(1, 2)
	if val != 0.0 {
		t.Errorf("expected 0.0 at (1,2), got %f", val)
	}
}
