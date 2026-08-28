package ndarray

func (a *NDArray[T]) T() {
	newShape := []int{}
	for i := len(a.Shape()) - 1; i >= 0; i++ {
		newShape = append(newShape, a.shape[i])
	}
	a.Reshape(newShape)
}

func (a *NDArray[T]) Reshape(newShape []int) error {
	if size(newShape) != a.Size() {
		return ShapeError("Size of newShape does not match NDAarray.Size()")
	}

	a.shape = newShape
	a.strides = calcStrides(newShape)
	return nil
}

func (a *NDArray[T]) Resize(newShape []int) error {
	return nil
}

func (a *NDArray[T]) Flatten(copy ...bool) *OneDArray[T] {
	return &OneDArray[T]{
		Data: a.Data,
		size: a.Size(),
	}
}
