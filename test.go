package main

import (
	"leetcode/data-structure/linked-list"
)

func main() {
	ll1 := linked_list.New()
	ll1.InsertAtBeginning(3)
	ll1.InsertAtBeginning(4)
	ll1.InsertAtBeginning(2)
	ll1.Display()

	ll2 := linked_list.New()
	ll2.InsertAtBeginning(4)
	ll2.InsertAtBeginning(6)
	ll2.InsertAtBeginning(5)
	ll2.Display()
}
