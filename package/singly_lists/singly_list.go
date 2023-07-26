package singlylists

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func PrintList(l *List) {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

/* List INsertion */
/*
	• Inserting a new node before the head (at the beginning)
	• Inserting a new node after the tail (at the end of the list)
	• Inserting a new node at the middle of the list (random location)
*/
func (list *List) AddAtTheBeginning(value int) {
	newNode := &Node{data: value}

	// Update the next pointer of new node, to point to the current head.
	newNode.next = list.head

	// Update head pointer to point to the new node.
	list.head = newNode

}

func (list *List) AddAtTheEnd(value int) {
	newNode := &Node{data: value}

	// check if the the head node pointer in nil
	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head

	// traverse the list in the until you find the last node
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *List) AddInTheMiddle(value int, position int) { // position starts from 1 not 0 like indices
	newNode := &Node{data: value}

	if list.head == nil {
		list.head = newNode
		return
	}

	// traverse the list until the position where we want to insert
	var before_node, position_node *Node = nil, list.head

	var count int = 1

	for position_node != nil && count < position {
		count++

		before_node = position_node

		position_node = position_node.next
	}

	before_node.next = newNode
	newNode.next = position_node
}
