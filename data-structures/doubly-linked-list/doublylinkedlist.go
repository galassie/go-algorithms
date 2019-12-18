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

func (node *DoublyLinkedListNode) toString() string {
	var sb strings.Builder

	currentNode := node
	for {
		if currentNode == nil {
			break
		}

		sb.WriteString(strconv.Itoa(currentNode.value))
		sb.WriteString(" <-> ")
		currentNode = currentNode.next
	}

	sb.WriteString("EODLL")
	return sb.String()
}
