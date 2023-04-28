package queue

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Push(value T) {
	q.items = append(q.items, value)
}

func (q *Queue[T]) Front() T {
	return q.items[0]
}

func (q *Queue[T]) Pop() {
	q.items = q.items[1:]
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) Empty() bool {
	return len(q.items) == 0
}
