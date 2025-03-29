package main

import (
	"fmt"
	"leetcode/data-structure/linked-list"
)

func mergeTwoLists(list1 *linked_list.Node, list2 *linked_list.Node) *linked_list.Node {
	llResult := &linked_list.Node{}

	headResult := llResult

	head1 := list1
	head2 := list2

	for head1 != nil && head2 != nil {
		fmt.Println(head1.Val, head2.Val)
		if head1.Val <= head2.Val {
			headResult.Next = &linked_list.Node{Val: head1.Val, Next: nil}
			headResult = headResult.Next
			head1 = head1.Next
		} else {
			headResult.Next = &linked_list.Node{Val: head2.Val, Next: nil}
			headResult = headResult.Next
			head2 = head2.Next
		}
	}

	for head1 != nil {
		headResult.Next = &linked_list.Node{Val: head1.Val, Next: nil}
		headResult = headResult.Next
		head1 = head1.Next
	}

	for head2 != nil {
		headResult.Next = &linked_list.Node{Val: head2.Val, Next: nil}
		headResult = headResult.Next
		head2 = head2.Next
	}

	return llResult.Next
}

func main() {
	ll1 := linked_list.New()
	ll1.InsertAtEnd(1)
	ll1.InsertAtEnd(3)
	ll1.InsertAtEnd(5)
	ll1.InsertAtEnd(7)
	ll1.InsertAtEnd(9)
	ll1.InsertAtEnd(11)
	ll1.InsertAtEnd(13)
	ll1.Display()

	ll2 := linked_list.New()
	ll2.InsertAtEnd(2)
	ll2.InsertAtEnd(4)
	ll2.InsertAtEnd(6)
	ll2.InsertAtEnd(8)
	ll2.Display()

	ll3 := mergeTwoLists(ll1.Head, ll2.Head)
	ll3.Print()
}
