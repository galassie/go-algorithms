package algorithms

import (
	"testing"
)

func Test_SelectionSort(t *testing.T) {
	cases := []struct {
		input          []int
		expectedOutput []int
	}{
		{[]int{0, 5, 3, 2, 2}, []int{0, 2, 2, 3, 5}},
		{[]int{64, 25, 12, 22, 11}, []int{11, 12, 22, 25, 64}},
		{[]int{-4, -50, 9, 2}, []int{-50, -4, 2, 9}},
	}

	for _, c := range cases {
		actual := SelectionSort(c.input)

		if !Equal(actual, c.expectedOutput) {
			t.Errorf("SelectionSort was incorrect, got: %v, want: %v.", actual, c.expectedOutput)
		}
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
