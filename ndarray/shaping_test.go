package ndarray

import (
	"testing"
	"slices"
)
func TestTReshapeFlatten(t *testing.T) {
	tests := []struct{
		name string
		arr *NDArray[Default_Scalar]
		newShape []int
		passes bool
	}{
		{
			name: "only shapes",
			passes: true,
			newShape: []int{4},
		},
		{
			name: "shapes and values",
			arr: &NDArray[Default_Scalar]{
				Data: []ScalarLike[Default_Scalar] {
					{
						Value: Default_Scalar(1),
					},
					{
						Value: Default_Scalar(2),
					},
					{
						Value: Default_Scalar(3),
					},
					{
						Value: Default_Scalar(4),
					},
				},
				shape: []int{2, 2},
				strides: calcStrides([]int{2, 2}),
			},
			passes: true,
			newShape: []int{2, 2},

		},

	}
	arr0, _ := Full([]int{4, 12, 2}, Default_Scalar(4))
	tests[0].arr = arr0

	for _, tc := range(tests) {
		t.Run(tc.name, func(t *testing.T) {
			shape := tc.arr.Shape()
			shapeT := []int{}
			for i := len(shape) -1; i >= 0; i-- {
				shapeT = append(shapeT, shape[i])
			}

			if !slices.Equal(shapeT, tc.arr.T().Shape()) {}
		})
	}
}

func TestResize(t *testing.T) {
	// NOT IMPLEMENTED
}
