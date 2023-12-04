package main

import "testing"

func TestBinaryTree(t *testing.T) {
	tree := &Tree{}
	tree.insert(50).
		insert(30).
		insert(20).
		insert(40).
		insert(70).
		insert(60).
		insert(80)

	if tree.root.left.data != 30 {
		t.Fatalf("expected %d, received %d", 30, tree.root.left.data)
	}

	if tree.root.right.data != 70 {
		t.Fatalf("expected %d, received %d", 70, tree.root.right.data)
	}
}

