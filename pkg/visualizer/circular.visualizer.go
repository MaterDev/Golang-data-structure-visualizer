package visualizer

import (
	"fmt"
	"golang-data-structure-visualizer/pkg/linkedlist"
	"strings"
)

type CircularLinkedListVisualizer struct {}

func NewCircularLinkedListVisualizer() *CircularLinkedListVisualizer {
	return &CircularLinkedListVisualizer{}
}

func (v *CircularLinkedListVisualizer) Visualize(list interface{}) string {
	// Incoming list, which is CircularLinkedList
	// Check incoming CLL

	// Check to make sure incoming list is the correct type.
	circularList, ok := list.(*linkedlist.CircularLinkedList)
	if !ok {
		return "Error: Not a CircularLinkedList"
	}
	// If head is nil then list is empty 
	if circularList.Head == nil {
		return "Empty circular list!"
	}
	// A slice of strings to hold the visualization text
	var result []string

	fmt.Printf("First Head of list: %v", circularList.Head.Value)

	// Start traversing from head.Next
	current := circularList.Head

	/**
		For loop
			Append to 'result' a new string of the current value
			Reassign 'current' to current.Next, in order to move to the next node.
			// If current == circularlist.Head, that means we have completed the circle
	*/
	result = append(result, fmt.Sprintf("%v", current.Value))
	if circularList.Head == circularList.Head.Next {
		result = append(result, "⤶")
		return strings.Join(result, " -> ")
	}

	current = circularList.Head.Next

	for current != circularList.Head {
		result = append(result, fmt.Sprintf("%v", current.Value))
		current = current.Next
	}

	result = append(result, "⤶")
	return strings.Join(result, " -> ")
}