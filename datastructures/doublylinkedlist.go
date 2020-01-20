package datastructures

import (
	"errors"
	"strconv"
	"strings"
)

// DoublyLinkedListNode represents an integer-value node in a DoublyLinkedList
type DoublyLinkedListNode struct {
	value    int
	previous *DoublyLinkedListNode
	next     *DoublyLinkedListNode
}

// DoublyLinkedList represents an integer-value DoublyLinkedList
type DoublyLinkedList struct {
	first *DoublyLinkedListNode
	last  *DoublyLinkedListNode
}

func newDoublyLinkedListNode(value int) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{value, nil, nil}
}

func newDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil}
}

func (doublyLinkedList *DoublyLinkedList) elementAt(index int) (int, error) {
	elementToEvaluate := doublyLinkedList.first
	if elementToEvaluate == nil {
		return -1, errors.New("DoublyLinkedList is empty")
	}

	for i := 0; i < index; i++ {
		elementToEvaluate = elementToEvaluate.next
		if elementToEvaluate == nil {
			return -1, errors.New("Index is out of range")
		}
	}

	return elementToEvaluate.value, nil
}

func (doublyLinkedList *DoublyLinkedList) addHead(value int) *DoublyLinkedList {
	node := newDoublyLinkedListNode(value)

	if doublyLinkedList.first != nil {
		node.next = doublyLinkedList.first
		doublyLinkedList.first.previous = node
	} else {
		doublyLinkedList.last = node
	}

	doublyLinkedList.first = node
	return doublyLinkedList
}

func (doublyLinkedList *DoublyLinkedList) removeHead() *DoublyLinkedList {
	if doublyLinkedList.first != nil {
		newFirst := doublyLinkedList.first.next
		doublyLinkedList.first.next = nil
		newFirst.previous = nil
		doublyLinkedList.first = newFirst
	}

	return doublyLinkedList
}

func (doublyLinkedList *DoublyLinkedList) addTail(value int) *DoublyLinkedList {
	node := newDoublyLinkedListNode(value)

	if doublyLinkedList.last != nil {
		node.previous = doublyLinkedList.last
		doublyLinkedList.last.next = node
	} else {
		doublyLinkedList.first = node
	}

	doublyLinkedList.last = node
	return doublyLinkedList
}

func (doublyLinkedList *DoublyLinkedList) removeTail() *DoublyLinkedList {
	if doublyLinkedList.last != nil {
		nodeToRemove := doublyLinkedList.last
		if nodeToRemove.previous != nil {
			newLastNode := nodeToRemove.previous

			nodeToRemove.previous = nil
			newLastNode.next = nil

			doublyLinkedList.last = newLastNode
		} else {
			doublyLinkedList.first = nil
			doublyLinkedList.last = nil
		}
	}

	return doublyLinkedList
}

func (doublyLinkedList *DoublyLinkedList) toString() string {
	var sb strings.Builder

	currentNode := doublyLinkedList.first
	for {
		if currentNode == nil {
			break
		}

		prevValue := "nil"
		if currentNode.previous != nil {
			prevValue = strconv.Itoa(currentNode.previous.value)
		}
		sb.WriteString(prevValue + "<~")
		sb.WriteString(strconv.Itoa(currentNode.value))
		sb.WriteString(" -- ")
		currentNode = currentNode.next
	}

	sb.WriteString("EODLL")
	return sb.String()
}
