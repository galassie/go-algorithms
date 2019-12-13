package main

import "testing"

func TestBasicLinkedList(t *testing.T) {
	linkedList := LinkedListNode { 10, &LinkedListNode { 20, &LinkedListNode { 30, nil } } }
	result := linkedList.toString()
	expected := "10 -> 20 -> 30 -> EOLL"
    if result != expected {
       t.Errorf("LinkedList toString was incorrect, got: %s, want: %s.", result, expected)
    }
}