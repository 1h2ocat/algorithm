package list

import "testing"

func TestLinkedList_Front(t *testing.T) {
	var list LinkedList[int]
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	list.PushFront(4)

	if list.Front() != 4 {
		t.Errorf("Front() = %d; want 4", list.Front())
	}
}

func TestLinkedList_Back(t *testing.T) {
	var list LinkedList[int]
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	if list.Back() != 4 {
		t.Errorf("Back() = %d; want 4", list.Back())
	}
}

func TestLinkedList_Begin(t *testing.T) {
	var list LinkedList[int]
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	list.PushFront(4)

	if list.Begin().Value != 4 {
		t.Errorf("Begin() = %d; want 4", list.Begin().Value)
	}
}

func TestLinkedList_End(t *testing.T) {
	var list LinkedList[int]
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	if list.End().Value != 4 {
		t.Errorf("End() = %d; want 4", list.End().Value)
	}
}

func TestLinkedList_Erase(t *testing.T) {
	var list LinkedList[int]
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	list.PushFront(4)

	list.Erase(list.Begin().Next())

	if list.Begin().Next().Value != 2 {
		t.Errorf("Erase() = %d; want 2", list.Begin().Next().Next().Value)
	}
}

func TestLinkedList_Insert(t *testing.T) {
	var list LinkedList[int]
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	list.Insert(list.Begin(), 5)

	if list.Begin().Next().Value != 5 {
		t.Errorf("Insert() = %d; want 5", list.Begin().Next().Value)
	}
}

func TestLinkedList_PopBack(t *testing.T) {
	var list LinkedList[int]
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	list.PopBack()

	if list.Back() != 3 {
		t.Errorf("PopBack() = %d; want 3", list.Back())
	}
}

func TestLinkedList_PopFront(t *testing.T) {
	var list LinkedList[int]
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	list.PushFront(4)

	list.PopFront()

	if list.Front() != 3 {
		t.Errorf("PopFront() = %d; want 3", list.Front())
	}
}

func TestLinkedList_RemoveIf(t *testing.T) {
	var list LinkedList[int]
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	list.RemoveIf(func(i int) bool {
		return i == 2
	})

	if list.Begin().Next().Value != 3 {
		t.Errorf("RemoveIf() = %d; want 3", list.Begin().Next().Value)
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	var list LinkedList[int]
	list.PushBack(4)
	list.PushBack(3)
	list.PushBack(2)
	list.PushBack(1)

	list.Reverse()

	node := list.Begin()
	for i := 1; i <= 4; i++ {
		if node.Value != i {
			t.Errorf("Reverse() = %d; want %d", node.Value, i)
		}
		node = node.Next()
	}
}

func TestLinkedList_Size(t *testing.T) {
	var list LinkedList[int]
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	if list.Size() != 4 {
		t.Errorf("Size() = %d; want 4", list.Size())
	}
}

func TestLinkedList_Sort(t *testing.T) {
	var list LinkedList[int]
	list.PushBack(4)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(1)

	list.Sort(func(i, j int) bool {
		return i > j
	})

	node := list.Begin()
	for i := 1; i <= 4; i++ {
		if node.Value != i {
			t.Errorf("Sort() = %d; want %d", node.Value, i)
		}
		node = node.Next()
	}
}

func TestLinkedList_Swap(t *testing.T) {
	var list LinkedList[int]
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	var list2 LinkedList[int]
	list2.PushBack(5)
	list2.PushBack(6)
	list2.PushBack(7)
	list2.PushBack(8)

	list.Swap(&list2)

	if list.Begin().Value != 5 {
		t.Errorf("Swap() = %d; want 2", list.Begin().Value)
	}
}
