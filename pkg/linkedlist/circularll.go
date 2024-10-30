package linkedlist

type CircularLinkedList struct {
	Head *Node
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{}
}

func (l *CircularLinkedList) Insert(value interface{}) {
	newNode := &Node{Value: value}
	
	if l.Head == nil {
		l.Head = newNode
		l.Head.Next = l.Head
	} else {
		// Make new node point to the first node
		newNode.Next = l.Head.Next
		// Make head point to the new node
		l.Head.Next = newNode 
		// Make new node the head
		l.Head = newNode

		// newNode.Next = l.Head.Next
		// l.Head.Next = newNode
		// l.Head = newNode
	}
}