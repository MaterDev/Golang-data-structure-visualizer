package main

import (
	"fmt"
	"golang-data-structure-visualizer/pkg/linkedlist"
	"golang-data-structure-visualizer/pkg/visualizer"
)

func main() {
	// Create factory
	factory := visualizer.NewVisualizerFactory()

	// Demo simple linked list
	fmt.Println("\nSimple Linked List Demo:")
	simpleList := linkedlist.NewSimpleLinkedList()
	simpleList.Insert(1)
	simpleList.Insert(2)
	simpleList.Insert(3)
	simpleList.Insert(4)
	simpleList.Insert(5)
	simpleList.Insert(6)
	simpleList.Insert(7)
	simpleList.Insert(8)
	simpleList.Insert(9)

	if vis, err := factory.CreateVisualizer(simpleList); err == nil {
		fmt.Println(vis.Visualize((simpleList)))
	}

	// Demo Doubly linked list
	fmt.Println("\nDoubly Linked List Demo:")
	doublyList := linkedlist.NewDoublyLinkedList()
	doublyList.Insert(1)
	doublyList.Insert(2)
	doublyList.Insert(3)
	doublyList.Insert(4)
	doublyList.Insert(5)
	doublyList.Insert(6)
	doublyList.Insert(7)
	doublyList.Insert(8)
	doublyList.Insert(9)

	if vis, err := factory.CreateVisualizer(doublyList); err == nil {
		fmt.Println(vis.Visualize((doublyList)))
	}

	// Demo circular linked list
	fmt.Println("\nCircular Linked List Demo:")
	circularList := linkedlist.NewCircularLinkedList()
	circularList.Insert(1)
	circularList.Insert(2)
	circularList.Insert(3)
	circularList.Insert(4)
	circularList.Insert(5)
	circularList.Insert(6)
	circularList.Insert(7)
	circularList.Insert(8)
	circularList.Insert(9)

	if vis, err := factory.CreateVisualizer(circularList); err == nil {
		fmt.Println(vis.Visualize((circularList)))
	}


}