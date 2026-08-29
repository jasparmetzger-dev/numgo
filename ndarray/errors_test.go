package ndarray

import (
	"errors"
	"fmt"
	"testing"
)

func TestShapeNotBroadcastableError(t *testing.T) {
	shape1 := []int{2, 3, 4}
	shape2 := []int{3, 4, 5}

	err := ShapeNotBroadcastableError{
		shape1: shape1,
		shape2: shape2,
	}
	if err.Error() != fmt.Sprintf("Shapes %v and %v are not broadcastable", shape1, shape2) {
		t.Errorf("ShapeNotBroadcastableError has some error")
	}
}

func TestShapeAndIndexError(t *testing.T) {
	strs := []string{"a", "abc", "Sjsddpfj"}
	for _, s := range strs {
		if errors.New(s) != ShapeError(s) {
			t.Errorf("Error with ShapeError")
		}
		if errors.New(s) != IndexError(s) {
			t.Errorf("Error with IndexError")
		}
	}
}
