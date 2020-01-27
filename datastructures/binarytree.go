package datastructures

// BinaryTreeNode represents an integer-value node in a BinaryTree
type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func newBinaryTreeNode(value int) *BinaryTreeNode {
	return &BinaryTreeNode{value, nil, nil}
}

func (binaryTreeNode *BinaryTreeNode) getValue() int {
	return binaryTreeNode.value
}

func (binaryTreeNode *BinaryTreeNode) hasLeft() bool {
	return binaryTreeNode.left != nil
}

func (binaryTreeNode *BinaryTreeNode) hasRight() bool {
	return binaryTreeNode.right != nil
}

func (binaryTreeNode *BinaryTreeNode) getLeft() *BinaryTreeNode {
	return binaryTreeNode.left
}

func (binaryTreeNode *BinaryTreeNode) getRight() *BinaryTreeNode {
	return binaryTreeNode.right
}

func (binaryTreeNode *BinaryTreeNode) setLeft(value int) *BinaryTreeNode {
	binaryTreeNode.left = newBinaryTreeNode(value)
	return binaryTreeNode
}

func (binaryTreeNode *BinaryTreeNode) setRight(value int) *BinaryTreeNode {
	binaryTreeNode.left = newBinaryTreeNode(value)
	return binaryTreeNode
}
