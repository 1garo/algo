package main

import (
	"testing"
)


func TestQueue(t *testing.T) {
	queue := &Queue[int]{}
	queue.New()

	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	val := queue.peek()

	if *val != 1 {
		t.Fatalf("expected %d, received %d", 1, *val)
	}

	queue.deque()
}
