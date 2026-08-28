package ndarray

import (
	"errors"
	"fmt"
)

type ShapeNotBroadcastableError struct {
	shape1 []int
	shape2 []int
}

func (e *ShapeNotBroadcastableError) Error() string {
	return fmt.Sprintf("Shapes %v and %v are not broadcastable", e.shape1, e.shape2)
}

func ShapeError(err string) error {
	return errors.New(err)
}
func IndexError(err string) error {
	return errors.New(err)
}
