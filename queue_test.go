package main

import (
	"testing"
)


func TestQueue(t *testing.T) {
	queue := &Queue[int]{}
	queue.New()

	_, err := queue.deque()
	if err == nil {
		t.Fatalf("expected: %s, received %v", err.Error(), nil)
	}

	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	if queue.Length != 3 {
		t.Fatalf("expected %d, received %d", 3, queue.Length)
	}

	val := queue.peek()

	if *val != 1 {
		t.Fatalf("expected %d, received %d", 1, *val)
	}

	val, _ = queue.deque()
	if queue.Length != 2 {
		t.Fatalf("expected %d, received %d", 2, queue.Length)
	}

	if *val != 1 {
		t.Fatalf("expected %d, received %d", 1, *val)
	}
}
