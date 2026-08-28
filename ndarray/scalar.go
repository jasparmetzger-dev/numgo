package ndarray

type Default_Scalar float64

// Datatype of _NDArray
type ScalarLike[T Scalar] struct {
	Value T
	IsNaN bool
}

type Scalar interface {
	_SInt | _UInt | _Float
}

type _Float interface {
	~float32 | ~float64
}
type _SInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type _UInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// ---- functions ----

func UseScalar[T Scalar](scalar T) ScalarLike[T] {
	return ScalarLike[T]{
		Value: scalar,
		IsNaN: false,
	}
}

func Zero[T any]() T {
	var zero T
	return zero
}

func NaN[T Scalar]() ScalarLike[T] {
	return ScalarLike[T]{
		Value: Zero[T](),
		IsNaN: true,
	}
}
