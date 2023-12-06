package main

type SNode[T any] struct {
	Value T
	prev  *SNode[T]
}

type Stack[T any] struct {
	Length uint
	head   *SNode[T]
}

func (qn *Stack[T]) push(item T) {
	node := &SNode[T]{
		Value: item,
	}

	qn.Length++;
	if qn.head == nil {
		qn.head = node
		return
	}

	node.prev = qn.head
	qn.head = node
}

func (qn *Stack[T]) pop() *T {
	if qn.head == nil || qn.Length <= 0{
		return nil
	}

	qn.Length--;
	head := qn.head
	qn.head = qn.head.prev

	// free?! Maybe it's not needed
	head.prev = nil

	return &head.Value
}

func (qn *Stack[T]) peek() *T {
	return &qn.head.Value
}
