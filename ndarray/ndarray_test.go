package ndarray

import (
	"reflect"
	"slices"
	"testing"
)

func TestNDArray(t *testing.T) {
	t.Run("TestShapeSizeDType", func(t *testing.T) {
		tests := []struct {
			name   string
			shape  []int
			size   int
			val    float64
			dtype  reflect.Type
			passes bool
		}{
			{
				name:   "Valid 3D Float64",
				shape:  []int{4, 2, 5},
				size:   40,
				val:    float64(1.2),
				dtype:  reflect.TypeOf(float64(0)),
				passes: true,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				arr, err := Full(tc.shape, tc.val)
				if err != nil {
					if tc.passes {
						t.Fatalf("unexpected error: %v", err)
					}
					return
				}

				if tc.dtype != arr.DType() {
					t.Errorf("DType(): has %v, should %v", arr.DType(), tc.dtype)
				}
				if !slices.Equal(tc.shape, arr.Shape()) {
					t.Errorf("Shape(): has %v, should %v", arr.Shape(), tc.shape)
				}
				if tc.size != arr.Size() {
					t.Errorf("Shape(): has %v, should %v", arr.Size(), tc.size)
				}
				if tc.size != len(arr.Data) {
					t.Errorf("Size(): has %v, should %v", len(arr.Data), tc.size)
				}
			})
		}
	})
}
