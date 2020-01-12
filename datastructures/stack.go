package datastructures

import (
	"strconv"
	"strings"
)

// Stack represents a stack data stcture using LinkedList
type Stack struct {
	ll *LinkedList
}

func newStack() *Stack {
	return &Stack{newLinkedList()}
}

func (stack *Stack) push(value int) *Stack {
	stack.ll = stack.ll.addHead(value)
	return stack
}

func (stack *Stack) pop() int {
	result := stack.ll.first.value
	stack.ll = stack.ll.removeHead()
	return result
}

func (stack *Stack) peek() int {
	return stack.ll.first.value
}

func (stack *Stack) toString() string {
	var sb strings.Builder

	currentElement := stack.ll.first
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
