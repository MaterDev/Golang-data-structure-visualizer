package visualizer

import (
	"testing"
	"golang-data-structure-visualizer/pkg/linkedlist"
)

func TestSimpleLinkedListVisualizer(t *testing.T) {
	list := linkedlist.NewSimpleLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	visualizer := NewSimpleLinkedListVisualizer()
	result := visualizer.Visualize(list)

	expected := "3 -> 2 -> 1 -> nil"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}