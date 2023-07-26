package main

import (
	"fmt"
	singlylists "linked_lists/package/singly_lists"
)

func main() {
	singly_list := &singlylists.List{}

	singly_list.AddAtTheEnd(1)
	singly_list.AddAtTheEnd(2)
	singly_list.AddAtTheEnd(3)
	singly_list.AddAtTheEnd(4)
	singly_list.AddAtTheEnd(5)
	singly_list.AddAtTheBeginning(6)
	singly_list.AddInTheMiddle(7, 3)

	fmt.Println("Initial List: ")
	singlylists.PrintList(singly_list)
}
