package datastructures

import "testing"

func Test_Queue_New(t *testing.T) {
	actual := newQueue()

	if actual == nil {
		t.Errorf("NewQueue was incorrect, got nil.")
	}
}

func Test_Queue_Enqueue(t *testing.T) {
	cases := []struct {
		q            *Queue
		valuesToPush []int
		expected     string
	}{
		{newQueue(), []int{3}, "3 - EOQ"},
		{newQueue(), []int{25, 41, 90}, "25 - 41 - 90 - EOQ"},
		{newQueue(), []int{4, 35}, "4 - 35 - EOQ"},
	}

	for _, c := range cases {
		result := c.q
		for _, v := range c.valuesToPush {
			result = result.enqueue(v)
		}
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("Queue enqueue was incorrect, got: %s, want: %s.", actual, c.expected)
		}
	}
}

func Test_Queue_Dequeue(t *testing.T) {
	cases := []struct {
		q                        *Queue
		timesToDequeue           int
		expectedLastValueDequeue int
		expected                 string
	}{
		{newQueue().enqueue(7).enqueue(10), 1, 7, "10 - EOQ"},
		{newQueue().enqueue(10).enqueue(25).enqueue(41), 2, 25, "41 - EOQ"},
		{newQueue().enqueue(4), 1, 4, "EOQ"},
	}

	for _, c := range cases {
		result := c.q
		actualLastValueDequeue := 0
		for i := 0; i < c.timesToDequeue; i++ {
			actualLastValueDequeue = result.dequeue()
		}
		actual := result.toString()

		if actual != c.expected {
			t.Errorf("Queue dequeue was incorrect, got: %s, want: %s.", actual, c.expected)
		}
		if actualLastValueDequeue != c.expectedLastValueDequeue {
			t.Errorf("Queue dequeue last value was incorrect, got: %d, want: %d.", actualLastValueDequeue, c.expectedLastValueDequeue)
		}
	}
}
