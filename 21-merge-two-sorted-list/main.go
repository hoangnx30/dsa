package main

import (
	"leetcode/data-structure/linked-list"
)

func mergeTwoLists(list1 *linked_list.Node, list2 *linked_list.Node) *linked_list.Node {
	llResult := &linked_list.Node{}

	headResult := llResult

	head1 := list1
	head2 := list2

	for head1 != nil && head2 != nil {
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

func mergeTwoListsEnhance(list1 *linked_list.Node, list2 *linked_list.Node) *linked_list.Node {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	list1Head := list1
	list2Head := list2

	resultList := list1
	if list1Head.Val > list2.Val {
		resultList = list2
		list2Head = list2Head.Next
	} else {
		list1Head = list1Head.Next
	}

	resultListHead := resultList

	for list1Head != nil || list2Head != nil {
		if list1Head == nil {
			resultListHead.Next = list2Head
			break
		}

		if list2Head == nil {
			resultListHead.Next = list1Head
			break
		}

		if list1Head.Val <= list2Head.Val {
			resultListHead.Next = list1Head
			resultListHead = resultListHead.Next
			list1Head = list1Head.Next
		} else {
			resultListHead.Next = list2Head
			resultListHead = resultListHead.Next
			list2Head = list2Head.Next
		}
	}

	return resultList
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

	ll3 := mergeTwoListsEnhance(ll1.Head, ll2.Head)
	ll3.Print()
}
