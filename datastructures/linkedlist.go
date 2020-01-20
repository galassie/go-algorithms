package datastructures

import (
	"errors"
	"strconv"
	"strings"
)

// LinkedListNode represents an integer-value node in a LinkedList
type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

// LinkedList represents an integer-value LinkedList
type LinkedList struct {
	first *LinkedListNode
}

func newLinkedListNode(value int) *LinkedListNode {
	return &LinkedListNode{value, nil}
}

func newLinkedList() *LinkedList {
	return &LinkedList{nil}
}

func (linkedList *LinkedList) elementAt(index int) (int, error) {
	elementToEvaluate := linkedList.first
	if elementToEvaluate == nil {
		return -1, errors.New("LinkedList is empty")
	}

	for i := 0; i < index; i++ {
		elementToEvaluate = elementToEvaluate.next
		if elementToEvaluate == nil {
			return -1, errors.New("Index is out of range")
		}
	}

	return elementToEvaluate.value, nil
}

func (linkedList *LinkedList) addHead(value int) *LinkedList {
	node := newLinkedListNode(value)

	if linkedList.first != nil {
		node.next = linkedList.first
	}

	linkedList.first = node
	return linkedList
}

func (linkedList *LinkedList) removeHead() *LinkedList {
	if linkedList.first != nil {
		toRemove := linkedList.first
		linkedList.first = toRemove.next
		toRemove.next = nil
	}

	return linkedList
}

func (linkedList *LinkedList) addTail(value int) *LinkedList {
	node := newLinkedListNode(value)

	if linkedList.first == nil {
		linkedList.first = node
	} else {
		nodeToEval := linkedList.first
		for nodeToEval.next != nil {
			nodeToEval = nodeToEval.next
		}
		nodeToEval.next = node
	}

	return linkedList
}

func (linkedList *LinkedList) removeTail() *LinkedList {
	if linkedList.first != nil && linkedList.first.next == nil {
		linkedList.first = nil
		return linkedList
	} else if linkedList.first != nil && linkedList.first.next != nil {
		possibleNewLast := linkedList.first
		possibleToRemove := possibleNewLast.next
		for possibleToRemove.next != nil {
			possibleNewLast = possibleToRemove
			possibleToRemove = possibleNewLast.next
		}
		possibleNewLast.next = nil
	}

	return linkedList
}

func (linkedList *LinkedList) toString() string {
	var sb strings.Builder

	currentNode := linkedList.first
	for {
		if currentNode == nil {
			break
		}

		sb.WriteString(strconv.Itoa(currentNode.value))
		sb.WriteString(" -> ")
		currentNode = currentNode.next
	}

	sb.WriteString("EOLL")
	return sb.String()
}
