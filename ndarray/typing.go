package ndarray


type Number interface {
	_IntLike | _FloatLike
}

type _IntLike interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | _UIntLike
}

type _FloatLike interface {
	~float32 | ~float64
}

type _UIntLike interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
