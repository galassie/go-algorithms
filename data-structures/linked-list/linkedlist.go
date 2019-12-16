package datastructures

import (
	"strconv"
	"strings"
)

type LinkedListNode struct {
	value int
	next  *LinkedListNode
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

func (node *LinkedListNode) addHead(newNode *LinkedListNode) LinkedListNode {
	if newNode == nil {
		return *node
	}

	newNode.next = node
	return *newNode
}

func (node *LinkedListNode) addTail(newNode *LinkedListNode) LinkedListNode {
	if newNode == nil {
		return *node
	}

	if node.next == nil {
		node.next = newNode
	} else {
		node.next.addTail(newNode)
	}

	return *node
}
