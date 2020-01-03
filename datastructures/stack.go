package datastructures

// Stack represents a stack data stcture using LinkedList
type Stack struct {
	head *LinkedListNode
}

func newStack() *Stack {
	return &Stack{nil}
}
