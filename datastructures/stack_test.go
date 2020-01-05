package datastructures

import "testing"

func Test_Stack_New(t *testing.T) {
	actual := newStack()

	if actual == nil {
		t.Errorf("NewStack was incorrect, got nil.")
	}
}

func Test_Stack_Push(t *testing.T) {
	cases := []struct {
		ll           *Stack
		valuesToPush []int
		expected     string
	}{
		{newStack(), []int{3}, "3 | EOS"},
		{newStack(), []int{25, 41, 90}, "90 | 41 | 25 | EOS"},
		{newStack(), []int{4, 35}, "35 | 4 | EOS"},
	}

	for _, c := range cases {
		result := c.ll
		for _, v := range c.valuesToPush {
			result = result.push(v)
		}
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("Stack push was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_Stack_ToString(t *testing.T) {
	cases := []struct {
		s        *Stack
		expected string
	}{
		{newStack().push(7).push(10), "10 | 7 | EOS"},
		{newStack().push(10).push(25).push(41), "41 | 25 | 10 | EOS"},
		{newStack().push(4), "4 | EOS"},
		{newStack(), "EOS"},
	}

	for _, c := range cases {
		actual := c.s.toString()

		if actual != c.expected {
			t.Errorf("Stack toString was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}
