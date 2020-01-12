package datastructures

import "testing"

func Test_LinkedListNode_New(t *testing.T) {
	actual := newLinkedListNode(3)

	expectedValue := 3
	if actual == nil {
		t.Errorf("NewLinkedListNode was incorrect, got nil.")
	}
	if actual.value != expectedValue {
		t.Errorf("NewLinkedListNode was incorrect, for value got: %d, want: %d.", actual.value, expectedValue)
	}
	if actual.next != nil {
		t.Errorf("NewLinkedListNode was incorrect, for next got: %v, want nil.", actual.next)
	}
}

func Test_LinkedList_New(t *testing.T) {
	actual := newLinkedList()

	if actual == nil {
		t.Errorf("NewLinkedList was incorrect, got nil.")
	}
}

func Test_LinkedList_AddHead(t *testing.T) {
	cases := []struct {
		ll          *LinkedList
		valuesToAdd []int
		expected    string
	}{
		{newLinkedList(), []int{7, 3}, "3 -> 7 -> EOLL"},
		{newLinkedList(), []int{10, 25, 41, 90}, "90 -> 41 -> 25 -> 10 -> EOLL"},
		{newLinkedList(), []int{3, 4, 35}, "35 -> 4 -> 3 -> EOLL"},
	}

	for _, c := range cases {
		result := c.ll
		for _, v := range c.valuesToAdd {
			result.addHead(v)
		}
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("LinkedList addHead was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_LinkedList_RemoveHead(t *testing.T) {
	cases := []struct {
		ll       *LinkedList
		expected string
	}{
		{newLinkedList().addTail(10).addTail(25), "25 -> EOLL"},
		{newLinkedList().addTail(3).addTail(4).addTail(33), "4 -> 33 -> EOLL"},
		{newLinkedList().addTail(3).addTail(4).addTail(33).addTail(92), "4 -> 33 -> 92 -> EOLL"},
	}

	for _, c := range cases {
		result := c.ll.removeHead()
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("LinkedList removeHead was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_LinkedList_AddTail(t *testing.T) {
	cases := []struct {
		ll         *LinkedList
		valueToAdd []int
		expected   string
	}{
		{newLinkedList(), []int{7, 3}, "7 -> 3 -> EOLL"},
		{newLinkedList(), []int{10, 25, 41, 90}, "10 -> 25 -> 41 -> 90 -> EOLL"},
		{newLinkedList(), []int{3, 4, 35}, "3 -> 4 -> 35 -> EOLL"},
	}

	for _, c := range cases {
		result := c.ll
		for _, v := range c.valueToAdd {
			result = result.addTail(v)
		}
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("LinkedList addTail was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_LinkedList_RemoveTail(t *testing.T) {
	cases := []struct {
		ll       *LinkedList
		expected string
	}{
		{newLinkedList().addTail(2), "EOLL"},
		{newLinkedList().addTail(10).addTail(25), "10 -> EOLL"},
		{newLinkedList().addTail(3).addTail(4).addTail(33), "3 -> 4 -> EOLL"},
		{newLinkedList().addTail(3).addTail(4).addTail(33).addTail(92), "3 -> 4 -> 33 -> EOLL"},
	}

	for _, c := range cases {
		result := c.ll.removeTail()
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("LinkedList removeTail was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_LinkedList_ToString(t *testing.T) {
	cases := []struct {
		ll       *LinkedList
		expected string
	}{
		{newLinkedList().addHead(7), "7 -> EOLL"},
		{newLinkedList().addHead(10).addTail(25).addTail(41), "10 -> 25 -> 41 -> EOLL"},
		{newLinkedList().addHead(3).addTail(4), "3 -> 4 -> EOLL"},
	}

	for _, c := range cases {
		actual := c.ll.toString()

		if actual != c.expected {
			t.Errorf("LinkedList toString was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}
