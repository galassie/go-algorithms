package sort

import (
	"testing"
)

func Test_BubleSort(t *testing.T) {
	cases := []struct {
		input          []int
		expectedOutput []int
	}{
		{[]int{0, 5, 3, 2, 2}, []int{0, 2, 2, 3, 5}},
		{[]int{64, 25, 12, 22, 11}, []int{11, 12, 22, 25, 64}},
		{[]int{-4, -50, 9, 2}, []int{-50, -4, 2, 9}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	}

	for _, c := range cases {
		actual := BubleSort(c.input)

		if !Equal(actual, c.expectedOutput) {
			t.Errorf("BubleSort was incorrect, got: %v, want: %v.", actual, c.expectedOutput)
		}
	}
}
