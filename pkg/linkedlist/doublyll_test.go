package linkedlist

import "testing"

func TestNewDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList()
	if list == nil{
		t.Fatal("NewDoublyLinkedList() should not return nil")
	}
	if list.Head != nil {
		t.Error("New list should have nil head")
	}
	if list.Tail != nil {
		t.Error("New list should have nil tail")
	}
}

func TestDoublyLinkedListInsert(t *testing.T) {
	list := NewDoublyLinkedList()
	// Insert a value of 1 in a new node (This should be the head)
	list.Insert(1)
	if list.Head == nil {
		t.Error("After insert, head should not be nil")
	}
	if list.Tail == nil {
		t.Error("After insert, tail should not be nil")
	}
	if list.Head.Value != 1 || list.Tail.Value != 1 {
		t.Errorf("Expected head and tail value to be 1, got head %v, tail %v", list.Head.Value, list.Tail.Value)
	}

	list.Insert(2)
	if list.Head.Value != 2 || list.Tail.Value != 1 {
		t.Errorf("Expected head value to be 2 and tail value to be 1, got head %v, tail %v", list.Head.Value, list.Tail.Value)
	}
	// ! If the current head's next doesnt point to the tail
		// ! or if the tail's previous doesnt point to the current head...
	if list.Head.Next != list.Tail || list.Tail.Prev != list.Head {
		t.Error("Links between nodes are not correct.")
	}

}