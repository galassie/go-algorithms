package datastructures

import (
	"errors"
	"testing"
)

func Test_DoublyLinkedListNode_New(t *testing.T) {
	actual := newDoublyLinkedListNode(14)

	expectedValue := 14
	if actual == nil {
		t.Errorf("NewDoublyLinkedListNode was incorrect, got nil.")
	}
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

func Test_DoublyLinkedList_New(t *testing.T) {
	actual := newDoublyLinkedList()

	if actual == nil {
		t.Errorf("NewDoublyLinkedList was incorrect, got nil.")
	}
}

func Test_DoublyLinkedList_ElementAt(t *testing.T) {
	cases := []struct {
		dll           *DoublyLinkedList
		index         int
		expectedValue int
		expectedError error
	}{
		{newDoublyLinkedList().addTail(10).addTail(25), 0, 10, nil},
		{newDoublyLinkedList().addTail(3).addTail(4).addTail(33), 2, 33, nil},
		{newDoublyLinkedList().addTail(3).addTail(4).addTail(33).addTail(92), 1, 4, nil},
		{newDoublyLinkedList(), 0, -1, errors.New("DoublyLinkedList is empty")},
		{newDoublyLinkedList().addTail(10).addTail(25), 3, -1, errors.New("Index is out of range")},
	}

	for _, c := range cases {
		actualValue, actualError := c.dll.elementAt(c.index)

		if actualValue != c.expectedValue {
			t.Errorf("DoublyLinkedList elementAt was incorrect for value, got: %d, want: %d.", actualValue, c.expectedValue)
		}

		if actualError == nil && c.expectedError != nil {
			t.Errorf("DoublyLinkedList elementAt was incorrect for error, expected an error but actual error is nil")
		} else if actualError != nil && c.expectedError == nil {
			t.Errorf("DoublyLinkedList elementAt was incorrect for error, got an error but wasn't expected")
		} else if actualError != nil && c.expectedError != nil && actualError.Error() != c.expectedError.Error() {
			t.Errorf("DoublyLinkedList elementAt was incorrect for error, got: \"%s\", want: \"%s\".", actualError, c.expectedError)
		}
	}
}

func Test_DoublyLinkedList_AddHead(t *testing.T) {
	cases := []struct {
		dll        *DoublyLinkedList
		valueToAdd []int
		expected   string
	}{
		{newDoublyLinkedList(), []int{7, 3}, "nil<~3 -- 3<~7 -- EODLL"},
		{newDoublyLinkedList(), []int{10, 25, 41, 90}, "nil<~90 -- 90<~41 -- 41<~25 -- 25<~10 -- EODLL"},
		{newDoublyLinkedList(), []int{3, 4, 35}, "nil<~35 -- 35<~4 -- 4<~3 -- EODLL"},
	}

	for _, c := range cases {
		result := c.dll
		for _, v := range c.valueToAdd {
			result = result.addHead(v)
		}
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("DoublyLinkedList addHead was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_DoublyLinkedList_RemoveHead(t *testing.T) {
	cases := []struct {
		dll      *DoublyLinkedList
		expected string
	}{
		{newDoublyLinkedList().addTail(10).addTail(25), "nil<~25 -- EODLL"},
		{newDoublyLinkedList().addTail(7).addTail(4).addTail(33), "nil<~4 -- 4<~33 -- EODLL"},
		{newDoublyLinkedList().addTail(21).addTail(4).addTail(33).addTail(92), "nil<~4 -- 4<~33 -- 33<~92 -- EODLL"},
	}

	for _, c := range cases {
		result := c.dll.removeHead()
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("DoublyLinkedList removeHead was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_DoublyLinkedList_AddTail(t *testing.T) {
	cases := []struct {
		dll        *DoublyLinkedList
		valueToAdd []int
		expected   string
	}{
		{newDoublyLinkedList(), []int{7, 3}, "nil<~7 -- 7<~3 -- EODLL"},
		{newDoublyLinkedList(), []int{10, 25, 41, 90}, "nil<~10 -- 10<~25 -- 25<~41 -- 41<~90 -- EODLL"},
		{newDoublyLinkedList(), []int{3, 4, 35}, "nil<~3 -- 3<~4 -- 4<~35 -- EODLL"},
	}

	for _, c := range cases {
		result := c.dll
		for _, v := range c.valueToAdd {
			result = result.addTail(v)
		}
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("DoublyLinkedList addTail was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_DoublyLinkedList_RemoveTail(t *testing.T) {
	cases := []struct {
		dll      *DoublyLinkedList
		expected string
	}{
		{newDoublyLinkedList().addHead(7).addHead(5), "nil<~5 -- EODLL"},
		{newDoublyLinkedList().addHead(10).addHead(23).addHead(4), "nil<~4 -- 4<~23 -- EODLL"},
		{newDoublyLinkedList(), "EODLL"},
		{newDoublyLinkedList().addHead(40), "EODLL"},
	}

	for _, c := range cases {
		result := c.dll.removeTail()
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("DoublyLinkedList removeTail was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}
func Test_DoublyLinkedList_ToString(t *testing.T) {
	cases := []struct {
		dll      *DoublyLinkedList
		expected string
	}{
		{newDoublyLinkedList().addHead(3), "nil<~3 -- EODLL"},
		{newDoublyLinkedList().addHead(35).addHead(31).addHead(17), "nil<~17 -- 17<~31 -- 31<~35 -- EODLL"},
		{newDoublyLinkedList().addHead(4).addHead(3), "nil<~3 -- 3<~4 -- EODLL"},
	}

	for _, c := range cases {
		actual := c.dll.toString()

		if actual != c.expected {
			t.Errorf("DoublyLinkedList toString was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}
