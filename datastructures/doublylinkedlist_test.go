package datastructures

import "testing"

func Test_DoublyLinkedList_New(t *testing.T) {
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
func Test_DoublyLinkedList_AddHeadWithNil(t *testing.T) {
	dll := newDoublyLinkedListNode(3)

	result := dll.addHead(nil)

	expected := "nil<~3 -- EODLL"
	actual := result.toString()
	if actual != expected {
		t.Errorf("DoublyLinkedList addHead with nil was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func Test_DoublyLinkedList_AddHead(t *testing.T) {
	cases := []struct {
		dll        	*DoublyLinkedListNode
		valueToAdd 	[]int
		expected   	string
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

func Test_DoublyLinkedList_RemoveHeadWithNextNil(t *testing.T) {
	dll := newDoublyLinkedListNode(3)

	actual := dll.removeHead()

	if actual != nil {
		t.Errorf("DoublyLinkedList removeHead with next nil was incorrect, got: %v, want: nil", actual)
	}
}

func Test_DoublyLinkedList_RemoveHead(t *testing.T) {
	cases := []struct {
		dll         *DoublyLinkedListNode
		expected   	string
	}{
		{newDoublyLinkedListNode(10).addTail(newDoublyLinkedListNode(25)), "nil<~25 -- EODLL"},
		{newDoublyLinkedListNode(3).addTail(newDoublyLinkedListNode(4)).addTail(newDoublyLinkedListNode(33)), "nil<~4 -- 4<~33 -- EODLL"},
		{newDoublyLinkedListNode(3).addTail(newDoublyLinkedListNode(4)).addTail(newDoublyLinkedListNode(33)).addTail(newDoublyLinkedListNode(92)), "nil<~4 -- 4<~33 -- 33<~92 -- EODLL"},
	}

	for _, c := range cases {
		result := c.dll.removeHead()
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("DoublyLinkedList removeHead was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_DoublyLinkedList_RemoveHeadFromInnerNode(t *testing.T) {
	node1 := newDoublyLinkedListNode(3)
	node2 := newDoublyLinkedListNode(4)
	node3 := newDoublyLinkedListNode(33)
	node4 := newDoublyLinkedListNode(92)
	
	node1.addTail(node2).addTail(node3).addTail(node4)

	result := node3.removeHead()

	actual := result.toString()
	expected := "nil<~4 -- 4<~33 -- 33<~92 -- EODLL"
	if actual != expected {
		t.Errorf("DoublyLinkedList removeHead was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func Test_DoublyLinkedList_AddTailWithNil(t *testing.T) {
	dll := newDoublyLinkedListNode(3)

	result := dll.addTail(nil)

	expected := "nil<~3 -- EODLL"
	actual := result.toString()
	if actual != expected {
		t.Errorf("DoublyLinkedList addTail with nil was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func Test_DoublyLinkedList_AddTail(t *testing.T) {
	cases := []struct {
		dll         *DoublyLinkedListNode
		valueToAdd 	[]int
		expected   	string
	}{
		{newDoublyLinkedListNode(7), []int{3}, "nil<~7 -- 7<~3 -- EODLL"},
		{newDoublyLinkedListNode(10), []int{25, 41, 90}, "nil<~10 -- 10<~25 -- 25<~41 -- 41<~90 -- EODLL"},
		{newDoublyLinkedListNode(3), []int{4, 35}, "nil<~3 -- 3<~4 -- 4<~35 -- EODLL"},
	}

	for _, c := range cases {
		result := c.dll
		for _, v := range c.valueToAdd {
			result = result.addTail(newDoublyLinkedListNode(v))
		}
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("DoublyLinkedList addTail was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_DoublyLinkedList_ToString(t *testing.T) {
	cases := []struct {
		dll      	*DoublyLinkedListNode
		expected 	string
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
