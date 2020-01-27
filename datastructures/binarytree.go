package datastructures

// BinaryTreeNode represents an integer-value node in a BinaryTree
type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

// BinaryTree represents an integer-value BinaryTree
type BinaryTree struct {
	root *BinaryTreeNode
}

func newBinaryTreeNode(value int) *BinaryTreeNode {
	return &BinaryTreeNode{value, nil, nil}
}
