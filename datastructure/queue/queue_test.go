package queue

import "testing"

func TestQueue(t *testing.T) {
	var queue Queue[int]

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	if queue.Front() != 1 {
		t.Errorf("Front() = %d; want 1", queue.Front())
	}

	if queue.Back() != 4 {
		t.Errorf("Back() = %d; want 4", queue.Back())
	}

	if queue.Size() != 4 {
		t.Errorf("Size() = %d; want 4", queue.Size())
	}

	queue.Pop()

	if queue.Front() != 2 {
		t.Errorf("Front() = %d; want 2", queue.Front())
	}

	if queue.Size() != 3 {
		t.Errorf("Size() = %d; want 3", queue.Size())
	}
}
