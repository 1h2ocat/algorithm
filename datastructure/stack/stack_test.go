package stack

import "testing"

func TestStack_Push(t *testing.T) {
	var stack Stack[int]

	stack.Push(1)

	if stack.Size() != 1 {
		t.Errorf("Size() = %d; want 1", stack.Size())
	}

	if stack.Top() != 1 {
		t.Errorf("Top() = %d; want 1", stack.Top())
	}
}

func TestStack_Top(t *testing.T) {
	var stack Stack[int]

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Top() != 3 {
		t.Errorf("Top() = %d; want 3", stack.Top())
	}
}

func TestStack_Pop(t *testing.T) {
	var stack Stack[int]

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	stack.Pop()

	if stack.Size() != 3 {
		t.Errorf("Size() = %d; want 3", stack.Size())
	}
}

func TestStack_Size(t *testing.T) {
	var stack Stack[int]

	stack.Push(1)
	stack.Push(2)

	if stack.Size() != 2 {
		t.Errorf("Size() = %d; want 2", stack.Size())
	}
}

func TestStack_Empty(t *testing.T) {
	var stack Stack[int]

	if !stack.Empty() {
		t.Errorf("Empty() = false; want true")
	}
}
