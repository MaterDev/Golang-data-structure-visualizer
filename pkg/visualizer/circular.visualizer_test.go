package visualizer

import (
	"golang-data-structure-visualizer/pkg/linkedlist"
	"testing"
)

// Will run the visualizer against an expected output.
func TestCircularLinkedListVisualizer(t *testing.T) {
	list := linkedlist.NewCircularLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	vis := NewCircularLinkedListVisualizer()
	result := vis.Visualize(list)

	expected := "3 -> 1 -> 2 -> ⤶" // the tail icon represents next being the new head.
	if result != expected {
		t.Errorf("Expected %v, but got %s", expected, result)
	}
}