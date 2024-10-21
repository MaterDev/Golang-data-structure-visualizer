package linkedlist

import "testing"

func TestNewCircularLinkedList(t *testing.T) {
	list := NewCircularLinkedList()
	if list == nil {
		t.Fatal("NewCircularLinkedList() should not return nil")
	}
	if list.Head != nil {
		t.Error("New list shoud have nil head")
	}
}

func TestCircularLinkedListInsert(t *testing.T) {
	list := NewCircularLinkedList()
	
	list.Insert(1)
	if list.Head == nil {
		t.Error("After insert, head should not be nil")
	}
	if list.Head.Value != 1 {
		t.Errorf("Expected head value to be 1, got %v", list.Head.Value)
	}
	if list.Head.Next != list.Head {
		t.Error("In a single-element circular list, head should point to itself")
}

	list.Insert(2)
	if list.Head.Value != 2 {
		t.Errorf("Expected head value to be 2, got %v", list.Head.Value)
	}
	if list.Head.Next.Value != 1 {
		t.Errorf("Expected second node value to be 1, got %v", list.Head.Next.Value)
	}
	if list.Head.Next.Next != list.Head {
		t.Error("Last node should point back to head in a circular list")
	}
}
