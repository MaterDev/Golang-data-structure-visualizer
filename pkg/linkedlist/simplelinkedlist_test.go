package linkedlist

import "testing"

// Test for a simple linkedlist
func TestNewSimpleLinkedList(t *testing.T) {
	list := NewSimpleLinkedList()

	if list == nil {
		t.Error("NewSimpleLinkedList() should not return nil")
	}
	if list.Head != nil {
		t.Error("New list should have nil head")
	}
}

func TestSimpleLinkedListInsert(t *testing.T) {
	list := NewSimpleLinkedList()
	list.Insert(1)
	if list.Head == nil {
		t.Error("After insert, head should not be nil")
	}
	if list.Head.Value != 1 {
		t.Errorf("Expected head value to be 1, got %v", list.Head.Value)
	}
}

func TestNewDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList()
	if list == nil{
		t.Error("NewDoublyLinkedList() should not return nil")
	}
	if list.Head != nil {
		t.Error("New list should have nil head")
	}
	if list.Tail != nil {
		t.Error("New list should have nil tail")
	}
}

// TODO Test insert into doubly linked list
