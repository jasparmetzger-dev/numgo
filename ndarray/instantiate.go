package ndarray

func Array[T Scalar](data any) (*NDArray[T], error) {
	return nil, nil
}

func Full[T Scalar](shape []int, value T) (*NDArray[T], error) {
	if len(shape) == 0 {
		return &NDArray[T]{}, ShapeError("shape cannot be empty")
	}

	size := size(shape)
	if size <= 0 {
		return nil, ShapeError("invalid dimension sizes")
	}

	data := make([]ScalarLike[T], size)
	for i := range data {
		data[i].Value = value
		data[i].IsNaN = false
	}

	arr := NDArray[T]{
		Data:    data,
		shape:   shape,
		strides: calcStrides(shape),
	}
	return &arr, nil
}

func Zeroes[T Scalar](shape []int) (*NDArray[T], error) {
	var zero T
	return Full(shape, zero)
}

func Ones[T Scalar](shape []int) (*NDArray[T], error) {
	var one T = T(1)
	return Full(shape, one)
}
