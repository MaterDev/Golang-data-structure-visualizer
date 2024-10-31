package visualizer

import (
	"testing"
	"golang-data-structure-visualizer/pkg/linkedlist"
)

func TestVisualizerFactory(t *testing.T) {
	factory := NewVisualizerFactory()

	t.Run("create simple linked list visualizer", func(t *testing.T) {
		list := linkedlist.NewSimpleLinkedList()
		vis, err := factory.CreateVisualizer(list)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if vis == nil {
			t.Error("Expected visualizer, got nil")
		}
	})

	t.Run("create doubly linked list visualizer", func(t *testing.T) {
		list := linkedlist.NewDoublyLinkedList()
		vis, err := factory.CreateVisualizer(list)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if vis == nil {
			t.Error("Expected visualizer, got nil")
		}
	})

	t.Run("create circular linked list visualizer", func(t *testing.T) {
		list := linkedlist.NewCircularLinkedList()
		vis, err := factory.CreateVisualizer(list)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if vis == nil {
			t.Error("Expected visualizer, got nil")
		}
	})

	t.Run("create unsupported linked list visualizer", func(t *testing.T) {
		invalidList := "not a list"
		vis, err := factory.CreateVisualizer(invalidList)
		if err == nil {
			t.Error("Expected error for unsupported type, got nil")
		}
		if vis != nil {
			t.Error("Expected nil visualizer for unsupported type")
		}
	})
}
