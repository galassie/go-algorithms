package datastructures

import "testing"

func TestToString(t *testing.T) {
	cases := []struct {
		ll       *LinkedListNode
		expected string
	}{
		{&LinkedListNode{7, nil}, "7 -> EOLL"},
		{&LinkedListNode{10, &LinkedListNode{25, &LinkedListNode{41, nil}}}, "10 -> 25 -> 41 -> EOLL"},
		{&LinkedListNode{3, &LinkedListNode{4, nil}}, "3 -> 4 -> EOLL"},
	}

	for _, c := range cases {
		actual := c.ll.toString()

		if actual != c.expected {
			t.Errorf("LinkedList toString was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func TestAddHeadWithNull(t *testing.T) {
	ll := LinkedListNode{3, nil}

	result := ll.addHead(nil)

	expected := "3 -> EOLL"
	actual := result.toString()
	if actual != expected {
		t.Errorf("LinkedList addHead with nil was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestAddHead(t *testing.T) {
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

func TestAddTailWithNull(t *testing.T) {
	ll := LinkedListNode{3, nil}

	result := ll.addTail(nil)

	expected := "3 -> EOLL"
	actual := result.toString()
	if actual != expected {
		t.Errorf("LinkedList addTail with nil was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestAddTail(t *testing.T) {
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
