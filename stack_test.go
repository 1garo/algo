package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := &Stack[uint]{}

	stack.push(0)
	stack.push(1)
	stack.push(2)

	if stack.Length != 3 {
		t.Fatalf("expected: %d, received: %d", 3, stack.Length)
	}

	val := stack.pop()
	if stack.Length != 2 {
		t.Fatalf("expected: %d, received: %d", 2, stack.Length)
	}

	if *val != 2 {
		t.Fatalf("expected: %d, received: %d", 2, *val)
	}

	peeked := stack.peek()
	if *peeked != 1 {
		t.Fatalf("expected: %d, received: %d", 1, *val)
	}

}
