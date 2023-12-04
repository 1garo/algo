package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 8}
	ok := BinarySearch(items, 4)

	if ok == -1 {
		t.Fatalf("expected %d, received: %d", 3, ok)
	}
}
