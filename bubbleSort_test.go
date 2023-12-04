package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	items := []int{0, 2, 1, 6, 3, 8, 4}
	sortedItems := BubbleSort(items)
	if !reflect.DeepEqual(items, sortedItems) {
		t.Fatalf("expected %v: received %v", items, sortedItems)
	}
}

