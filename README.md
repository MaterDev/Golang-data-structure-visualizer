# GoLang Data Structure Visualizer

A Go package implementing the Factory Method pattern to visualize different types of linked lists. This project demonstrates test-driven development (TDD), Go interfaces, and clean code organization through visualization of simple, doubly, and circular linked lists.

![cover](./images/cover.png)

## Features

- Three types of linked lists:
  - Simple Linked List
  - Doubly Linked List 
  - Circular Linked List
- Visualizers for each list type using ASCII art
- Factory Method pattern for visualizer creation
- Test coverage for all components
- Command-line demo application

## Installation

```bash
# Clone the repository
git clone https://github.com/MaterDev/Golang-data-structure-visualizer.git

# Navigate to project directory
cd golang-data-structure-visualizer

# Install dependencies
go mod tidy
```

## Usage

### Running the Demo

```bash
go run cmd/visualizer/main.go
```

This will demonstrate visualization of all three list types with 9 nodes each:

```
Simple Linked List Demo:
9 -> 8 -> 7 -> 6 -> 5 -> 4 -> 3 -> 2 -> 1 -> nil

Doubly Linked List Demo:
nil <-> 9 <-> 8 <-> 7 <-> 6 <-> 5 <-> 4 <-> 3 <-> 2 <-> 1 <-> nil

Circular Linked List Demo:
9 -> 8 -> 7 -> 6 -> 5 -> 4 -> 3 -> 2 -> 1 -> ⤶
```

### Implementation Example

```go
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
```

## Running Tests

Run all tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## Design Patterns Used

### Factory Method Pattern

- Interface: `LinkedListVisualizer`
- Concrete Types: `SimpleLinkedListVisualizer`, `DoublyLinkedListVisualizer`, `CircularLinkedListVisualizer`
- Factory: `VisualizerFactory`

The Factory Method pattern allows for:

- Encapsulation of object creation
- Easy addition of new list types and visualizers
- Separation of concerns between list implementation and visualization