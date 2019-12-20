package datastructures

import (
	"strconv"
	"strings"
)

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

func newLinkedListNode(value int) *LinkedListNode {
	return &LinkedListNode{value, nil}
}

func (node *LinkedListNode) addHead(newNode *LinkedListNode) *LinkedListNode {
	if newNode == nil {
		return node
	}

	newNode.next = node
	return newNode
}

func (node *LinkedListNode) removeHead() *LinkedListNode {
	if node.next == nil {
		return nil
	}

	result := node.next
	node.next = nil
	return result
}

func (node *LinkedListNode) addTail(newNode *LinkedListNode) *LinkedListNode {
	if newNode == nil {
		return node
	}

	if node.next == nil {
		node.next = newNode
	} else {
		node.next.addTail(newNode)
	}

	return node
}

func (node *LinkedListNode) removeTail() *LinkedListNode {
	if node.next == nil {
		return nil
	}
	if node.next.next == nil {
		node.next = nil
		return node
	}

	node.next.removeTail()
	return node
}


func (node *LinkedListNode) toString() string {
	var sb strings.Builder

	currentNode := node
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
