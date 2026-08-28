package ndarray

// Types for Values in the array

type DEFAULTT float64

type Scalar interface {
	_ScalarLike | NaN | None
}

type NaN struct{}
type None struct{}

type _ScalarLike interface {
	_IntLike | _FloatLike
}

type _IntLike interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | _UIntLike
}

type _FloatLike interface {
	~float32 | ~float64
}

type _UIntLike interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// --- retyping the array ---
