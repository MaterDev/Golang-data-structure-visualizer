package visualizer

import (
	"fmt"
	"golang-data-structure-visualizer/pkg/linkedlist"
	"strings"
)

type SimpleLinkedListVisualizer struct {}

func NewSimpleLinkedListVisualizer() *SimpleLinkedListVisualizer {
	return &SimpleLinkedListVisualizer{}
}

func (v *SimpleLinkedListVisualizer) Visualize(list *linkedlist.SimpleLinkedList) string {
	var result []string
	current := list.Head
	for current != nil {
		result = append(result, fmt.Sprintf("%v", current.Value))
		current = current.Next
	}
	result = append(result, "nil")
	return strings.Join(result, " -> ")
}