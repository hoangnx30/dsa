package main

import (
	"leetcode/data-structure/linked-list"
)

func addTwoNumbers(l1 *linked_list.Node, l2 *linked_list.Node) *linked_list.Node {
	carry := 0
	sum := 0
	l3 := &linked_list.Node{}
	head := l3

	for l1 != nil || l2 != nil {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		carry = sum / 10

		l3.Next = &linked_list.Node{
			Val:  sum % 10,
			Next: nil,
		}
		l3 = l3.Next
	}

	if carry > 0 {
		l3.Next = &linked_list.Node{
			Val: carry,
		}
	}

	return head.Next
}

func addTwoNumbersEnhance(l1 *linked_list.Node, l2 *linked_list.Node) *linked_list.Node {
	carry := 0
	sum := 0
	l3 := &linked_list.Node{}
	head := l3

	for l1 != nil || l2 != nil || carry != 0 {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		carry = sum / 10

		l3.Next = &linked_list.Node{
			Val:  sum % 10,
			Next: nil,
		}
		l3 = l3.Next
	}

	return head.Next
}

func addTwoNumbersEnhance2(l1 *linked_list.Node, l2 *linked_list.Node) *linked_list.Node {
	carry := 0
	sum := 0
	l3 := &linked_list.Node{}
	head := l3

	for l1 != nil || l2 != nil || carry != 0 {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		carry = sum / 10

		l3.Val = sum % 10
		if l1 != nil || l2 != nil || carry != 0 {
			l3.Next = &linked_list.Node{}
			l3 = l3.Next
		}
	}

	return head
}

func main() {
	ll1 := linked_list.New()
	ll1.InsertAtBeginning(9)
	ll1.InsertAtBeginning(9)
	ll1.InsertAtBeginning(9)
	ll1.InsertAtBeginning(9)
	ll1.InsertAtBeginning(9)
	ll1.InsertAtBeginning(9)
	ll1.InsertAtBeginning(9)
	ll1.Display()

	ll2 := linked_list.New()
	ll2.InsertAtBeginning(9)
	ll2.InsertAtBeginning(9)
	ll2.InsertAtBeginning(9)
	ll2.InsertAtBeginning(9)
	ll2.Display()

	ll3 := addTwoNumbersEnhance2(ll1.Head, ll2.Head)
	ll3.Print()
}
