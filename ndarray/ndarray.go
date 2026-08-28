package ndarray

import (
	"reflect"
)

// --- Struct and basic methods ---

type NDArray[T Scalar] struct {
	data    []T
	shape   []int
	strides []int
}

func (a *NDArray[T]) Shape() []int { return a.shape }

func (a *NDArray[T]) Size() int { return size(a.shape) }

func (a *NDArray[T]) DType() reflect.Type {
	var zero T
	return reflect.TypeOf(zero)
}

func (a *NDArray[T]) At(index []int, val ...T) (T, error) {
	var dummy T
	i, err := a.index(index)
	if err != nil {
		return dummy, err
	}

	arrVal := a.data[i]
	if len(val) == 1 {
		a.data[i] = val[0]
	}
	return arrVal, nil
}

// --- simple initiating ---

func Full[T Scalar](shape []int, value T) (*NDArray[T], error) {
	if len(shape) == 0 {
		return &NDArray[T]{}, ShapeError("shape cannot be empty")
	}

	size := size(shape)
	if size <= 0 {
		return nil, ShapeError("invalid dimension sizes")
	}

	data := make([]T, size)
	for i := range data {
		data[i] = value
	}

	arr := NDArray[T]{
		data:    data,
		shape:   shape,
		strides: calcStrides(shape),
	}
	return &arr, nil
}

func Zeroes[T _ScalarLike](shape []int) (*NDArray[T], error) {
	var zero T
	return Full(shape, zero)
}
func Ones[T _ScalarLike](shape []int) (*NDArray[T], error) {
	var one T = T(1)
	return Full(shape, one)
}

// --- helpers ---

func size(shape []int) int {
	size := 1
	for _, dim := range shape {
		size *= dim
	}
	return size
}

func (a *NDArray[T]) index(indices []int, val ...int) (int, error) {
	if len(a.shape) != len(indices) {
		return 0, ShapeError("Index must have Dimensions of ndarray.shape")
	}
	if len(val) > 1 {
		return 0, ShapeError("Can only set one scalar at a time")
	}

	offset := 0
	for dim, idx := range indices {
		offset += idx * a.strides[dim]
	}
	return offset, nil
}
