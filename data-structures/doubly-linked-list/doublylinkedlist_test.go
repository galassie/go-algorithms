package datastructures

import "testing"

func TestToString(t *testing.T) {
	cases := []struct {
		dll      DoublyLinkedListNode
		expected string
	}{
		{DoublyLinkedListNode{3, nil, nil}, "3 <-> EODLL"},
		{DoublyLinkedListNode{17, nil, &DoublyLinkedListNode{31, nil, &DoublyLinkedListNode{35, nil, nil}}}, "17 <-> 31 <-> 35 <-> EODLL"},
		{DoublyLinkedListNode{3, nil, &DoublyLinkedListNode{4, nil, nil}}, "3 <-> 4 <-> EODLL"},
	}

	for _, c := range cases {
		actual := c.dll.toString()

		if actual != c.expected {
			t.Errorf("DoublyLinkedList toString was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}
