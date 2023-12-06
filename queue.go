package main

import "fmt"


type QNode[T any] struct {
	Value T
	next *QNode[T]
}

type Queue[T any] struct {
	Length uint
	head *QNode[T]
	tail *QNode[T]
}

func (qn *Queue[T]) New() *QNode[T] {
	var value T
	return &QNode[T]{
		Value: value,
		next: nil,
	}
}

func (qn *Queue[T]) enqueue(item T) {
	node := &QNode[T]{
		Value: item,
	}

	if qn.Length == 0 {
		qn.Length++;
		qn.head = node
		qn.tail = node
		return
	}
	qn.Length++;

	qn.tail.next = node
	qn.tail = node
}

func (qn *Queue[T]) deque() (*T, error) {
	if qn.tail == nil {
		return nil, fmt.Errorf("cannot deque from an empty queue")
	}

	qn.Length--;
	head := qn.head
	qn.head = qn.head.next

	// free, idk if this is needed since we have garbage collector
	head.next = nil

	return &head.Value, nil
}

func (qn *Queue[T]) peek() *T {
	return &qn.head.Value
}
