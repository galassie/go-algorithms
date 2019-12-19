package datastructures

import (
	"strconv"
	"strings"
)

type DoublyLinkedListNode struct {
	value    int
	previous *DoublyLinkedListNode
	next     *DoublyLinkedListNode
}

func newDoublyLinkedListNode(value int) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{value, nil, nil}
}

func (node *DoublyLinkedListNode) addHead(newNode *DoublyLinkedListNode) *DoublyLinkedListNode {
	if newNode == nil {
		return node
	}

	newNode.next = node
	node.previous = newNode
	return newNode
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
