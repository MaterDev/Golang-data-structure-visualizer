package visualizer

import (
	"fmt"
	"golang-data-structure-visualizer/pkg/linkedlist"
)

// VisualizerFactory is reponsible for creating the appropriate visualizer
type VisualizerFactory interface {
	CreateVisualizer(list interface{}) (LinkedListVisualizer, error)
}

type ConcreteVisualizerFactory struct{}

func NewVisualizerFactory() *ConcreteVisualizerFactory{
	return &ConcreteVisualizerFactory{}
}

func (f *ConcreteVisualizerFactory) CreateVisualizer(list interface{}) (LinkedListVisualizer, error) {
	switch list.(type) {
	case *linkedlist.SimpleLinkedList:
		return NewSimpleLinkedListVisualizer(), nil
	case *linkedlist.DoublyLinkedList:
		return NewDoublyLinkedListVisualizer(), nil
	case *linkedlist.CircularLinkedList:
		return NewCircularLinkedListVisualizer(), nil
	default:
		return nil, fmt.Errorf("unsupported list type")
	}
}
