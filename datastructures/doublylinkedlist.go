package datastructures

import (
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
}

func newDoublyLinkedListNode(value int) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{value, nil, nil}
}

func newDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil}
}

func (node *DoublyLinkedListNode) addHead(newNode *DoublyLinkedListNode) *DoublyLinkedListNode {
	if newNode == nil {
		return node
	}

	newNode.next = node
	node.previous = newNode
	return newNode
}

func (node *DoublyLinkedListNode) removeHead() *DoublyLinkedListNode {
	if node.previous != nil {
		return node.previous.removeHead()
	}
	if node.next == nil {
		return nil
	}

	result := node.next
	node.next = nil
	result.previous = nil
	return result
}

func (node *DoublyLinkedListNode) addTail(newNode *DoublyLinkedListNode) *DoublyLinkedListNode {
	if newNode == nil {
		return node
	}

	if node.next == nil {
		node.next = newNode
		newNode.previous = node
	} else {
		node.next.addTail(newNode)
	}

	return node
}

func (node *DoublyLinkedListNode) toString() string {
	var sb strings.Builder

	currentNode := node
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
