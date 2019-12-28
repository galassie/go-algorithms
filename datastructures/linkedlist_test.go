package datastructures

import "testing"

func Test_LinkedList_New(t *testing.T) {
	actual := newLinkedListNode(3)

	expectedValue := 3
	if actual.value != expectedValue {
		t.Errorf("NewLinkedList was incorrect, for value got: %d, want: %d.", actual.value, expectedValue)
	}
	if actual.next != nil {
		t.Errorf("NewLinkedList was incorrect, for next got: %v, want nil.", actual.next)
	}
}

func Test_LinkedList_AddHeadWithNil(t *testing.T) {
	ll := newLinkedListNode(3)

	result := ll.addHead(nil)

	expected := "3 -> EOLL"
	actual := result.toString()
	if actual != expected {
		t.Errorf("LinkedList addHead with nil was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func Test_LinkedList_AddHead(t *testing.T) {
	cases := []struct {
		ll         *LinkedListNode
		valueToAdd []int
		expected   string
	}{
		{newLinkedListNode(7), []int{3}, "3 -> 7 -> EOLL"},
		{newLinkedListNode(10), []int{25, 41, 90}, "90 -> 41 -> 25 -> 10 -> EOLL"},
		{newLinkedListNode(3), []int{4, 35}, "35 -> 4 -> 3 -> EOLL"},
	}

	for _, c := range cases {
		result := c.ll
		for _, v := range c.valueToAdd {
			result = result.addHead(newLinkedListNode(v))
		}
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("LinkedList addHead was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_LinkedList_RemoveHeadWithNextNil(t *testing.T) {
	ll := newLinkedListNode(3)

	actual := ll.removeHead()

	if actual != nil {
		t.Errorf("LinkedList removeHead with next nil was incorrect, got: %v, want: nil", actual)
	}
}

func Test_LinkedList_RemoveHead(t *testing.T) {
	cases := []struct {
		ll       *LinkedListNode
		expected string
	}{
		{newLinkedListNode(10).addTail(newLinkedListNode(25)), "25 -> EOLL"},
		{newLinkedListNode(3).addTail(newLinkedListNode(4)).addTail(newLinkedListNode(33)), "4 -> 33 -> EOLL"},
		{newLinkedListNode(3).addTail(newLinkedListNode(4)).addTail(newLinkedListNode(33)).addTail(newLinkedListNode(92)), "4 -> 33 -> 92 -> EOLL"},
	}

	for _, c := range cases {
		result := c.ll.removeHead()
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("LinkedList removeHead was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_LinkedList_AddTailWithNil(t *testing.T) {
	ll := newLinkedListNode(3)

	result := ll.addTail(nil)

	expected := "3 -> EOLL"
	actual := result.toString()
	if actual != expected {
		t.Errorf("LinkedList addTail with nil was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func Test_LinkedList_AddTail(t *testing.T) {
	cases := []struct {
		ll         *LinkedListNode
		valueToAdd []int
		expected   string
	}{
		{newLinkedListNode(7), []int{3}, "7 -> 3 -> EOLL"},
		{newLinkedListNode(10), []int{25, 41, 90}, "10 -> 25 -> 41 -> 90 -> EOLL"},
		{newLinkedListNode(3), []int{4, 35}, "3 -> 4 -> 35 -> EOLL"},
	}

	for _, c := range cases {
		result := c.ll
		for _, v := range c.valueToAdd {
			result = result.addTail(newLinkedListNode(v))
		}
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("LinkedList addTail was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_LinkedList_RemoveTailWithNextNil(t *testing.T) {
	ll := newLinkedListNode(3)

	actual := ll.removeTail()

	if actual != nil {
		t.Errorf("LinkedList removeTail with next nil was incorrect, got: %v, want: nil", actual)
	}
}

func Test_LinkedList_RemoveTail(t *testing.T) {
	cases := []struct {
		ll       *LinkedListNode
		expected string
	}{
		{newLinkedListNode(10).addTail(newLinkedListNode(25)), "10 -> EOLL"},
		{newLinkedListNode(3).addTail(newLinkedListNode(4)).addTail(newLinkedListNode(33)), "3 -> 4 -> EOLL"},
		{newLinkedListNode(3).addTail(newLinkedListNode(4)).addTail(newLinkedListNode(33)).addTail(newLinkedListNode(92)), "3 -> 4 -> 33 -> EOLL"},
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
		ll       *LinkedListNode
		expected string
	}{
		{newLinkedListNode(7), "7 -> EOLL"},
		{newLinkedListNode(10).addTail(newLinkedListNode(25)).addTail(newLinkedListNode(41)), "10 -> 25 -> 41 -> EOLL"},
		{newLinkedListNode(3).addTail(newLinkedListNode(4)), "3 -> 4 -> EOLL"},
	}

	for _, c := range cases {
		actual := c.ll.toString()

		if actual != c.expected {
			t.Errorf("LinkedList toString was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}
