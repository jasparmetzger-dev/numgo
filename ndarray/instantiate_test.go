package ndarray

import (
	"errors"
	"testing"
)

func TestArray(t *testing.T) {
	//Array() not implemented
}

func TestFull(t *testing.T) {
	tests := []struct {
		name   string
		shape  []int
		size   int
		val    any
		passes bool
	}{
		{
			name:   "Valid Float64",
			shape:  []int{4, 2, 5},
			size:   40,
			val:    float64(1.2),
			passes: true,
		},
		{
			name:   "Invalid Type (string)",
			shape:  []int{4, 2, 5},
			size:   40,
			val:    "invalid_type",
			passes: false,
		},
		{
			name:   "Empty Shape",
			shape:  []int{},
			size:   0,
			val:    1,
			passes: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var err error

			// Type Switch zur Behandlung des generic calls
			switch v := tc.val.(type) {
			case float64:
				_, err = Full(tc.shape, v)
			case int:
				_, err = Full(tc.shape, v)
			default:
				// Simuliert den Compile/Type-Error für ungültige Typen
				err = errors.New("unsupported type")
			}

			if (err == nil) != tc.passes {
				t.Errorf("Error in Full(): got err=%v, expected passes=%v (testcase: %v)", err, tc.passes, tc)
			}
		})
	}
}

func TestZeroesOnes(t *testing.T) {
	tests := []struct {
		name   string
		shape  []int
		size   int
		passes bool
		dtype  any
	}{
		{
			name:   "Valid",
			shape:  []int{4, 2, 5},
			size:   40,
			passes: true,
		},
		{
			name:   "Empty Shape",
			shape:  []int{},
			size:   0,
			passes: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			_, err = Zeroes[Default_Scalar](tc.shape)
			if (err == nil) != tc.passes {
				t.Errorf("Error in ZerosOnes(): got err=%v, expected passes=%v (testcase: %v)", err, tc.passes, tc)
			}
			_, err = Ones[Default_Scalar](tc.shape)
			if (err == nil) != tc.passes {
				t.Errorf("Error in ZerosOnes(): got err=%v, expected passes=%v (testcase: %v)", err, tc.passes, tc)
			}
		})
	}
}
