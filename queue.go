package main

type QNode[T any] struct {
	Value T
	next *QNode[T]
}

type Queue[T any] struct {
	Length uint
	head *QNode[T]
	tail *QNode[T]
}

func (qn *Queue[T]) enqueue(item T) {
	node := &QNode[T]{
		Value: item,
	}

	if qn.Length == 0 {
		qn.head = node
		qn.tail = node

		qn.Length++;
		return
	}

	qn.tail.next = node
	qn.tail = node
	qn.Length++;
}

func (qn *Queue[T]) deque() *T {
	if qn.tail == nil {
		return nil
	}

	head := qn.head
	qn.head = qn.head.next
	qn.Length--;

	// free, idk if this is needed since we have garbage collector
	head.next = nil

	return &head.Value
}

func (qn *Queue[T]) peek() *T {
	return &qn.head.Value
}
