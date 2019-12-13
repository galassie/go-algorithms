package main

import (
	"fmt"
	"strings"
	"strconv"
)

type LinkedListNode struct {
	value int
	next *LinkedListNode
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