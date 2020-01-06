package datastructures

import (
	"strconv"
	"strings"
)

// Stack represents a stack data stcture using LinkedList
type Stack struct {
	head *LinkedListNode
}

func newStack() *Stack {
	return &Stack{nil}
}

func (stack *Stack) push(value int) *Stack {
	newNode := newLinkedListNode(value)
	stack.head = stack.head.addHead(newNode)
	return stack
}

func (stack *Stack) pop() int {
	result := stack.head.value
	stack.head = stack.head.removeHead()
	return result
}

func (stack *Stack) toString() string {
	var sb strings.Builder

	currentElement := stack.head
	for {
		if currentElement == nil {
			break
		}

		sb.WriteString(strconv.Itoa(currentElement.value))
		sb.WriteString(" | ")
		currentElement = currentElement.next
	}

	sb.WriteString("EOS")
	return sb.String()
}
