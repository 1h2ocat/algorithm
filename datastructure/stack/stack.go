package stack

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Top() T {
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Pop() {
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) Empty() bool {
	return len(s.items) == 0
}
