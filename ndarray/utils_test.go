package ndarray

import (
	"fmt"
	"slices"
	"testing"
)

func TestSortingFunctions(t *testing.T) {
	// NOT IMPLEMENTED
	tests := []struct{
		name string
	}{

	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_ = tc.name
		})
	}
}

func TestCalcStrides(t *testing.T) {
	tests := []struct {
		name string
		shape []int
		strides []int
		raises bool
	} {
		{
			name: "scalar ndarray",
			shape: []int{1},
			strides: []int{1},
			raises: false,
		},
		{
			name: "normal strides",
			shape: []int{2, 3},
			strides: []int{},
			raises: false,
		},
		{
			name: "invalid shape",
			shape: []int{},
			strides: nil,
			raises: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotStrides := calcStrides(tc.shape)
			if !slices.Equal(gotStrides, tc.strides) && !tc.raises {
				t.Errorf(fmt.Sprintf("Error in %v: Strides should be %v, is %v", tc.name, tc.strides, gotStrides))
			}
			if slices.Equal(gotStrides, tc.strides) && tc.raises {
				t.Errorf(fmt.Sprintf("No Error in %v: Strides should be %v, is %v", tc.name, tc.strides, gotStrides))
			}
		})
	}
}
