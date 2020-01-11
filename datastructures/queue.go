package datastructures

import (
	"strconv"
	"strings"
)

// Queue represents a queue data stcture using LinkedList
type Queue struct {
	first *LinkedListNode
}

func newQueue() *Queue {
	return &Queue{nil}
}

func (queue *Queue) enqueue(value int) *Queue {
	newNode := newLinkedListNode(value)
	queue.first.addTail(newNode)
	return queue
}

func (queue *Queue) dequeue() int {
	result := queue.first.value
	queue.first = queue.first.removeHead()
	return result
}

func (queue *Queue) toString() string {
	var sb strings.Builder

	currentElement := queue.first
	for {
		if currentElement == nil {
			break
		}

		sb.WriteString(strconv.Itoa(currentElement.value))
		sb.WriteString(" - ")
		currentElement = currentElement.next
	}

	sb.WriteString("EOQ")
	return sb.String()
}
