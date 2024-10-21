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
		newNode.Next = l.Head.Next
		l.Head.Next = newNode
		l.Head = newNode
	}
}