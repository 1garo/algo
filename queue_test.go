package main

import (
	"testing"
)


func TestQueue(t *testing.T) {
	queue := &Queue[int]{}
	//queue.New()

	val := queue.deque()
	if val != nil {
		t.Fatalf("expected: %v, received %d", nil, *val)
	}

	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	if queue.Length != 3 {
		t.Fatalf("expected %d, received %d", 3, queue.Length)
	}

	val = queue.peek()

	if *val != 1 {
		t.Fatalf("expected %d, received %d", 1, *val)
	}

	val = queue.deque()
	if queue.Length != 2 {
		t.Fatalf("expected %d, received %d", 2, queue.Length)
	}

	if *val != 1 {
		t.Fatalf("expected %d, received %d", 1, *val)
	}
}
