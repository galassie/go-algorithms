package datastructures

import "testing"

func TestBasicLinkedList(t *testing.T) {
	cases := []struct {
		ll       LinkedListNode
		expected string
	}{
		{LinkedListNode{7, nil}, "7 -> EOLL"},
		{LinkedListNode{10, &LinkedListNode{25, &LinkedListNode{41, nil}}}, "10 -> 25 -> 41 -> EOLL"},
		{LinkedListNode{3, &LinkedListNode{4, nil}}, "3 -> 4 -> EOLL"},
	}

	for _, c := range cases {
		actual := c.ll.toString()
		if actual != c.expected {
			t.Errorf("LinkedList toString was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}
