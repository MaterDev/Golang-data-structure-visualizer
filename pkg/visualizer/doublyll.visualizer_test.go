package visualizer

import (
	"golang-data-structure-visualizer/pkg/linkedlist"
	"testing"
)

func TestDoublyLinkedListVisualizer(t *testing.T) {
	// Make a new doubly linked list
	list := linkedlist.NewDoublyLinkedList()
	// Insert 3 nodes
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	// Make new visualzier
	visualizer := NewDoublyLinkedListVisualizer()
	// Capture result of visualizer for the list
	result := visualizer.Visualize(list)

	// Check that result is same as expected.
	expected := "nil <-> 3 <-> 2 <-> 1 <-> nil"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}