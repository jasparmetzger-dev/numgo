package ndarray

import (
	"reflect"
)

// --- STRUCT AND BASIC METHODS ---

type NDArray[T Scalar] struct {
	Data    []ScalarLike[T]
	shape   []int
	strides []int
}

type OneDArray[T Scalar] struct {
	Data []ScalarLike[T]
	size int
}

type BoolNDArray struct {
	Data    []bool
	shape   []int
	strides []int
}

func (a *NDArray[T]) Shape() []int { return a.shape }

func (a *NDArray[T]) Size() int { return size(a.shape) }

func (a *NDArray[T]) DType() reflect.Type {
	var zero T
	return reflect.TypeOf(zero)
}

func (a *NDArray[T]) At(index []int, val ...ScalarLike[T]) (ScalarLike[T], error) {
	var dummy ScalarLike[T]
	i, err := a.index(index)
	if err != nil {
		return dummy, err
	}

	arrVal := a.Data[i]
	if len(val) == 1 {
		a.Data[i] = val[0]
	}
	return arrVal, nil
}

// --- HELPERS ---

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
