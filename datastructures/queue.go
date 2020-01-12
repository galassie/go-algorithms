package datastructures

import (
	"strconv"
	"strings"
)

// Queue represents a queue data stcture using LinkedList
type Queue struct {
	ll *LinkedList
}

func newQueue() *Queue {
	return &Queue{newLinkedList()}
}

func (queue *Queue) enqueue(value int) *Queue {
	queue.ll.addTail(value)
	return queue
}

func (queue *Queue) dequeue() int {
	result := queue.ll.first.value
	queue.ll = queue.ll.removeHead()
	return result
}

func (queue *Queue) toString() string {
	var sb strings.Builder

	currentElement := queue.ll.first
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
