package ndarray

func calcStrides(shape []int) []int {
	/*
		C-Contigous Stride Calculation.
		Implemented through a reverse for loop. O(N)

		For shape d_0, ... d_{N-1}:

		S_k = \prod_{j=(k+1)}^{N-1} (d_j)
		with S_{N-1} = 1

	*/

	N := len(shape)
	if N == 0 {
		return nil
	}

	S_k := make([]int, N)
	stride := 1

	for k := N - 1; k >= 0; k-- {
		S_k[k] = stride
		stride *= shape[k]
	}

	return S_k
}

func sortAndCopy[T Scalar](array *NDArray[T], axis int) (*NDArray[T], error) {
	return nil, nil
}
func sortInPlace[T Scalar](array *NDArray[T], axis int) error {
	return nil
}
