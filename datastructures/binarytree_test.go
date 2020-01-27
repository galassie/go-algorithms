package datastructures

import (
	"testing"
)

func Test_BinaryTreeNode_New(t *testing.T) {
	actual := newBinaryTreeNode(3)

	expectedValue := 3
	if actual == nil {
		t.Errorf("NewBinaryTreeNode was incorrect, got nil.")
	}
	if actual.value != expectedValue {
		t.Errorf("NewBinaryTreeNode was incorrect, for value got: %d, want: %d.", actual.value, expectedValue)
	}
	if actual.right != nil {
		t.Errorf("NewBinaryTreeNode was incorrect, for right got: %v, want nil.", actual.right)
	}
	if actual.left != nil {
		t.Errorf("NewBinaryTreeNode was incorrect, for left got: %v, want nil.", actual.left)
	}
}
