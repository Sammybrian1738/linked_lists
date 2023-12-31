package main

import (
	"fmt"
	singlylists "linked_lists/package/singly_lists"
)

func main() {
	singly_list := &singlylists.List{}

	// insertion
	singly_list.AddAtTheEnd(1)
	singly_list.AddAtTheEnd(2)
	singly_list.AddAtTheEnd(3)
	singly_list.AddAtTheEnd(4)
	singly_list.AddAtTheEnd(5)
	singly_list.AddAtTheBeginning(6)
	singly_list.AddInTheMiddle(7, 3)

	fmt.Println("Initial List: ")
	singlylists.PrintList(singly_list) // prints 6 1 7 2 3 4 5

	// deletion
	singly_list.DeleteFirstNode() // remains 1 7 2 3 4 5
	singly_list.DeleteLastNode()  // remains 1 7 2 3 4
	singly_list.DeleteIntermidiateNode(1)
	fmt.Println("After Deletion List: ")
	singlylists.PrintList(singly_list)

	singly_list.DeleteAllNodes()
	fmt.Println("After Deleting all Nodes: ")
	singlylists.PrintList(singly_list)
}
