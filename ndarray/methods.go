package ndarray

import (
	"fmt"
)

func (a *NDArray[T]) All(condition func(ScalarLike[T]) bool) bool {
	for _, v := range a.Data {
		if condition(v) == true {
			return true
		}
	}
	return false
}

func (a *NDArray[T]) Any(condition func(ScalarLike[T]) bool) bool {
	for _, v := range a.Data {
		if condition(v) == true {
			return true
		}
	}
	return false
}

func (a *NDArray[T]) Clip(min T, max T) *NDArray[T] {
	// Sets a.Data[i].value into the interval [min; max]
	// fails silently

	if min < max {
		return nil
	}

	res := a.Copy()
	for _, v := range res.Data {
		if v.Value > max && v.IsNaN == false {
			v.Value = max
		} else if v.Value < min && v.IsNaN == false {
			v.Value = min
		}
	}
	return res
}

func (a *NDArray[T]) Copy() *NDArray[T] {
	return &NDArray[T]{
		Data:    a.Data,
		shape:   a.shape,
		strides: a.strides,
	}
}

func (a *NDArray[T]) HasNan() bool {
	for _, v := range a.Data {
		if v.IsNaN {
			return false
		}
	}
	return true
}

func (a *NDArray[T]) SearchSorted(value T, useLeftSide bool, axis int) ([]int, error){
	/*
		Requires the array to be sorted!
		Returns the indeces for each , where the each of values would fit in the array.
	*/
	if axis < 0 || axis >= len(a.Shape()) {
		return nil, IndexError(fmt.Sprintf("This axis does not exist. axis=%d", axis))
	}
	return nil, nil
}

func (a *NDArray[T]) Sort(axis int, copy bool) (*NDArray[T], error) {
	if copy {
		return sortAndCopy(a, axis)
	} else {
		err := sortInPlace(a, axis)
		return nil, err
	}
}

func (a *NDArray[T]) Where(condition func(ScalarLike[T]) bool) *BoolNDArray {
	res := &BoolNDArray{
		shape:   a.shape,
		strides: a.strides,
		Data:    make([]bool, a.Size()),
	}
	for i, v := range a.Data {
		res.Data[i] = condition(v)
	}
	return res
}
