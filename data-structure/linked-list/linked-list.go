package linked_list

import "fmt"

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

type LinkedList struct {
	Head   *LinkedListNode
	Length int
}

func New() *LinkedList {
	return &LinkedList{
		Head:   nil,
		Length: 0,
	}
}

func (ll *LinkedList) InsertAtBeginning(value int) {
	newNode := &LinkedListNode{
		Value: value,
		Next:  ll.Head,
	}

	ll.Head = newNode
	ll.Length++
}

func (ll *LinkedList) InsertAtEnd(value int) {
	newNode := &LinkedListNode{
		Value: value,
		Next:  nil,
	}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	currentNode := ll.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = newNode
	ll.Length++
}

func (ll *LinkedList) Display() {
	if ll.Head == nil {
		fmt.Println("Linked List is empty")
	}

	currentNode := ll.Head
	for currentNode != nil {
		fmt.Printf("%d->", currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println("nil")
}
