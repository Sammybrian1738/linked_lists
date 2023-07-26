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

/* Singly List Insertion */
/*
	• Inserting a new node before the head (at the beginning)
	• Inserting a new node after the tail (at the end of the list)
	• Inserting a new node at the middle of the list (random location)

	Time Complexity: O(n), since, in the worst case, we may need to insert the node at the end of the list.

	Space Complexity: O(1), for creating one temporary variable.
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

/*Singly List Deletion*/
/*
	• Deleting the first node
	• Deleting the last node
	• Deleting an intermediate node.
*/

func (list *List) DeleteFirstNode() {
	if list.head == nil {
		fmt.Println("List empty")
		return
	}

	// Create a temporary node which will point to the same node as that of head.
	temp_node := list.head

	// Now, move the head nodes pointer to the next node of the temp node and dispose of the temporary node.
	list.head = temp_node.next

	temp_node = nil
}

func (list *List) DeleteLastNode() {
	if list.head == nil {
		fmt.Println("List empty")
		return
	}

	/* Traverse the list and while traversing maintain the previous node address also. By
	the time we reach the end of the list, we will have two pointers, one pointing to the
	tail node and the other pointing to the node before the tail node. */
	var previous_node, tail_node *Node = nil, list.head

	for tail_node.next != nil {
		previous_node = tail_node

		tail_node = tail_node.next
	}

	// Update previous node’s next pointer with NULL.
	previous_node.next = nil

	// Dispose of the tail node.
	tail_node = nil
}

func (list *List) DeleteIntermidiateNode(position int) {
	if list.head == nil {
		fmt.Println("List empty")
		return
	}

	if position == 1 {
		temp_node := list.head
		list.head = temp_node.next
		temp_node = nil
		return
	}

	// traverse the list until arriving at the position we want to delete
	var count int = 1
	var previous_node, position_node *Node = nil, list.head

	for position_node != nil && count < position {
		count++

		previous_node = position_node
		position_node = position_node.next
	}

	// check if the position node is nil, i.e at the end
	if position_node == nil {
		fmt.Println("Position does not exist")
		return
	}

	// change the previous node’s next pointer to the next pointer of the node to be deleted.
	previous_node.next = position_node.next

	position_node = nil
}

func (list *List) DeleteAllNodes() {
	var auxiliary_node, current_node *Node = nil, list.head

	for current_node != nil {
		auxiliary_node = current_node.next

		// free the next pointer of the current node
		current_node.next = nil

		// assign auxiliary node to the current node
		current_node = auxiliary_node
	}

	list.head = nil

}
