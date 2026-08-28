package ndarray

func (a *NDArray[T]) HasNan() bool {
	for _, v := range a.Data {
		if v.IsNaN {
			return false
		}
	}
	return true
}

func (a *NDArray[T]) Any() bool {
	return false
}

func (a *NDArray[T]) All() bool {
	return false
}

func (a *NDArray[T]) Where() *BoolNDArray {
	return nil
}

func (a *NDArray[T]) Clip(min T, max T) *NDArray[T] {
	//copy
	return nil
}

func (a *NDArray[T]) Copy() *NDArray[T] {
	return nil
}

func (a *NDArray[T]) SearchSorted(val T, useLeftSide bool) []int {
	return nil
}

func (a *NDArray[T]) Sort(axis int) error {
	return nil
}
