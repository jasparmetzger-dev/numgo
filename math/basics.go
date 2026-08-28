package math

import (
	"github.com/jasparmetzger-dev/numgo/ndarray"

)

func Mult[T ndarray.Scalar](array *ndarray.OneDArray[T]) ndarray.ScalarLike[T] {
	res := ndarray.ScalarLike[T]{
		Value: T(1.0),
		IsNaN: false,
	}

	for _, data := range array.Data {
		if data.IsNaN == true {
			res.IsNaN = true
			break
		}
		res.Value *= res.Value
	}
	return res
}

func Sum[T ndarray.Scalar](array *ndarray.OneDArray[T]) ndarray.ScalarLike[T] {
	res := ndarray.ScalarLike[T]{
		Value: T(0),
		IsNaN: false,
	}

	for _, data := range array.Data {
		if data.IsNaN == true {
			res.IsNaN = true
			break
		}
		res.Value += res.Value
	}
	return res
}
