package stack

import "testing"

func TestStack(t *testing.T) {
	var stack Stack[int]

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	if stack.Top() != 4 {
		t.Errorf("Top() = %d; want 4", stack.Top())
	}

	if stack.Size() != 4 {
		t.Errorf("Size() = %d; want 4", stack.Size())
	}

	stack.Pop()

	if stack.Top() != 3 {
		t.Errorf("Top() = %d; want 3", stack.Top())
	}

	if stack.Size() != 3 {
		t.Errorf("Size() = %d; want 3", stack.Size())
	}
}
