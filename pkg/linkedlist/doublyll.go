package linkedlist

type DoublyNode struct {
	Value interface{}
	Next *DoublyNode
	Prev *DoublyNode
}

type DoublyLinkedList struct {
	Head *DoublyNode
	Tail *DoublyNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) Insert(value interface{}) {
	// New doubly node with the incoming value
	newNode := &DoublyNode{Value: value}
	// Place new node in the data structure
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		// New node's Next will point to head
		newNode.Next = l.Head
		// Old head's Prev will point to new node
		l.Head.Prev = newNode
		// Head will become the new node
		l.Head = newNode

		// ! This will efectively make it where the new node is at the head of the list and the old head is shifted one place back.
	}
}