package datastructures

type Stack struct {
	head *LinkedListNode
}

func newStack() *Stack {
	return &Stack{nil}
}
