package linkedlist

type Node struct {
	Value interface{}
	Next *Node
}

type SimpleLinkedList struct {
	Head *Node
}

// Will return a pointer for a new Simple LinkedList
func NewSimpleLinkedList() *SimpleLinkedList {
	return &SimpleLinkedList{}
}

// Will make a new node, with the given value at the beginning of the list.
	// If head is nil (list is empty)
		// New node becomes the head
	// Otherwise, 
		// The new node is inserted before the current head and becomes the new head
// Every new node is inserted at the beginning.
func (l *SimpleLinkedList) Insert(value interface{}) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}