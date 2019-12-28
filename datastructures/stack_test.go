package datastructures

import "testing"

func Test_Stack_New(t *testing.T) {
	actual := newStack()

	if actual == nil {
		t.Errorf("NewStack was incorrect, got nil.")
	}
}
