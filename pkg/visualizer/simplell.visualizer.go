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

func (v *SimpleLinkedListVisualizer) Visualize(list interface{}) string {
	simpleList, ok := list.(*linkedlist.SimpleLinkedList)
	if !ok {
		return "Error: Not a SimpleLinkedList"
	}
	var result []string
	current := simpleList.Head
	for current != nil {
		result = append(result, fmt.Sprintf("%v", current.Value))
		current = current.Next
	}
	result = append(result, "nil")
	return strings.Join(result, " -> ")
}