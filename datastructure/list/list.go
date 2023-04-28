package list

type LinkedListNode[T any] struct {
	Value T
	next  *LinkedListNode[T]
	prev  *LinkedListNode[T]
}

type LinkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
}

// Front return first node value
func (s *LinkedList[T]) Front() T {
	return s.head.Value
}

// Back return last node value
func (s *LinkedList[T]) Back() T {
	return s.tail.Value
}

// Begin return first node
func (s *LinkedList[T]) Begin() *LinkedListNode[T] {
	return s.head
}

// End return last node
func (s *LinkedList[T]) End() *LinkedListNode[T] {
	return s.tail
}

// PushFront push value to front of list
func (s *LinkedList[T]) PushFront(value T) {
	node := &LinkedListNode[T]{Value: value}

	// If list is empty set node as head
	if s.head == nil {
		s.head = node
		s.tail = node
		return
	}

	s.head.prev = node
	node.next = s.head
	s.head = node
}

// PushBack push value to back of list
func (s *LinkedList[T]) PushBack(value T) {
	node := &LinkedListNode[T]{Value: value}

	// If list is empty set node as head
	if s.head == nil {
		s.head = node
		s.tail = node
		return
	}

	s.tail.next = node
	node.prev = s.tail
	s.tail = node
}

// PopFront pop value from front of list
func (s *LinkedList[T]) PopFront() {
	// Check if list is empty
	if s.head == nil {
		return
	}

	// If list has only one node
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
		return
	}

	s.head = s.head.next
	s.head.prev = nil
}

// PopBack pop value from back of list
func (s *LinkedList[T]) PopBack() {
	// Check if list is empty
	if s.tail == nil {
		return
	}

	// If list has only one node
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
		return
	}

	s.tail = s.tail.prev
	s.tail.next = nil
}

// Insert value to front of node
func (s *LinkedList[T]) Insert(pos *LinkedListNode[T], value T) *LinkedListNode[T] {
	node := &LinkedListNode[T]{Value: value}

	// if list is empty set node as head
	if pos == nil {
		s.head = node
		s.tail = node
		return node
	}

	node.next = pos.next
	pos.next.prev = node
	node.prev = pos
	pos.next = node

	return node
}

// Erase node from list
func (s *LinkedList[T]) Erase(node *LinkedListNode[T]) {
	node.prev.next = node.next
	node.next.prev = node.prev

	if node == s.head {
		s.head = node.next
	}

	if node == s.tail {
		s.tail = node.prev
	}
}

// Reverse list
func (s *LinkedList[T]) Reverse() {
	for curr := s.head; curr != nil; curr = curr.prev {
		curr.prev, curr.next = curr.next, curr.prev
	}

	s.head, s.tail = s.tail, s.head
}

// Size return size of list
func (s *LinkedList[T]) Size() int {
	size := 0

	for curr := s.head; curr != nil; curr = curr.next {
		size++
	}

	return size
}

// RemoveIf remove node from list if findFunc return true
func (s *LinkedList[T]) RemoveIf(findFunc func(T) bool) {
	for curr := s.head; curr != nil; curr = curr.next {
		if findFunc(curr.Value) {
			s.Erase(curr)
			return
		}
	}
}

// Sort list
func (s *LinkedList[T]) Sort(sortFunc func(T, T) bool) {
	// Bubble sort
	for curr := s.head; curr != nil; curr = curr.next {
		for next := curr.next; next != nil; next = next.next {
			if sortFunc(curr.Value, next.Value) {
				curr.Value, next.Value = next.Value, curr.Value
			}
		}
	}
}

// Swap two list
func (s *LinkedList[T]) Swap(other *LinkedList[T]) {
	s.head, other.head = other.head, s.head
	s.tail, other.tail = other.tail, s.tail
}

func (n *LinkedListNode[T]) Next() *LinkedListNode[T] {
	return n.next
}

func (n *LinkedListNode[T]) Prev() *LinkedListNode[T] {
	return n.prev
}
