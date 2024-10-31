package visualizer

import (
	"fmt"
	"golang-data-structure-visualizer/pkg/linkedlist"
	"strings"
)

type DoublyLinkedListVisualizer struct{}

func NewDoublyLinkedListVisualizer() *DoublyLinkedListVisualizer {
	return &DoublyLinkedListVisualizer{}
}

// Will return a result string by:
	// Traverse the doubly linked from the head towards next, until reaching nil for next
	// Append each to a result []strings, which will be joined and returned at the end
func (v *DoublyLinkedListVisualizer) Visualize(list interface{}) string {
	doublyList, ok := list.(*linkedlist.DoublyLinkedList)
	if !ok {
		return "Error: Not a DoublyLinkedList"
	}
	var result []string
	result = append(result, "nil")
	current := doublyList.Head

	// Loop over list by going from next to next.
	for current != nil {
		result = append(result, fmt.Sprintf("%v", current.Value))
		current = current.Next
	}

	result = append(result, "nil")
	return strings.Join(result, " <-> ")
}