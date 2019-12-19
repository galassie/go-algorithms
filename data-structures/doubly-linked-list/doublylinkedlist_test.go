package datastructures

import "testing"

func TestNewDoublyLinkedListNode(t *testing.T) {
	actual := newDoublyLinkedListNode(14)

	expectedValue := 14
	if actual.value != expectedValue {
		t.Errorf("NewDoublyLinkedListNode was incorrect, for value got: %d, want: %d.", actual.value, expectedValue)
	}
	if actual.previous != nil {
		t.Errorf("NewDoublyLinkedListNode was incorrect, for previous got: %v, want nil.", actual.previous)
	}
	if actual.next != nil {
		t.Errorf("NewDoublyLinkedListNode was incorrect, for next got: %v, want nil.", actual.next)
	}
}
func TestAddHeadWithNull(t *testing.T) {
	dll := newDoublyLinkedListNode(3)

	result := dll.addHead(nil)

	expected := "nil<~3 -- EODLL"
	actual := result.toString()
	if actual != expected {
		t.Errorf("DoublyLinkedList addHead with nil was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestAddHead(t *testing.T) {
	cases := []struct {
		dll        *DoublyLinkedListNode
		valueToAdd []int
		expected   string
	}{
		{newDoublyLinkedListNode(7), []int{3}, "nil<~3 -- 3<~7 -- EODLL"},
		{newDoublyLinkedListNode(10), []int{25, 41, 90}, "nil<~90 -- 90<~41 -- 41<~25 -- 25<~10 -- EODLL"},
		{newDoublyLinkedListNode(3), []int{4, 35}, "nil<~35 -- 35<~4 -- 4<~3 -- EODLL"},
	}

	for _, c := range cases {
		result := c.dll
		for _, v := range c.valueToAdd {
			result = result.addHead(newDoublyLinkedListNode(v))
		}
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("DoublyLinkedList addHead was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func TestAddTailWithNull(t *testing.T) {
	dll := newDoublyLinkedListNode(3)

	result := dll.addTail(nil)

	expected := "nil<~3 -- EODLL"
	actual := result.toString()
	if actual != expected {
		t.Errorf("DoublyLinkedList addTail with nil was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestAddTail(t *testing.T) {
	cases := []struct {
		ll         *DoublyLinkedListNode
		valueToAdd []int
		expected   string
	}{
		{newDoublyLinkedListNode(7), []int{3}, "nil<~7 -- 7<~3 -- EODLL"},
		{newDoublyLinkedListNode(10), []int{25, 41, 90}, "nil<~10 -- 10<~25 -- 25<~41 -- 41<~90 -- EODLL"},
		{newDoublyLinkedListNode(3), []int{4, 35}, "nil<~3 -- 3<~4 -- 4<~35 -- EODLL"},
	}

	for _, c := range cases {
		result := c.ll
		for _, v := range c.valueToAdd {
			result = result.addTail(newDoublyLinkedListNode(v))
		}
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("DoublyLinkedList addTail was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func TestToString(t *testing.T) {
	cases := []struct {
		dll      *DoublyLinkedListNode
		expected string
	}{
		{newDoublyLinkedListNode(3), "nil<~3 -- EODLL"},
		{newDoublyLinkedListNode(17).addTail(newDoublyLinkedListNode(31)).addTail(newDoublyLinkedListNode(35)), "nil<~17 -- 17<~31 -- 31<~35 -- EODLL"},
		{newDoublyLinkedListNode(3).addTail(newDoublyLinkedListNode(4)), "nil<~3 -- 3<~4 -- EODLL"},
	}

	for _, c := range cases {
		actual := c.dll.toString()

		if actual != c.expected {
			t.Errorf("DoublyLinkedList toString was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}
