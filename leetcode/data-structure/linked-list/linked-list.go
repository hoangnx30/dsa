package linked_list

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func New() *LinkedList {
	return &LinkedList{
		Head:   nil,
		Length: 0,
	}
}

func (ll *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{
		Val:  value,
		Next: ll.Head,
	}

	ll.Head = newNode
	ll.Length++
}

func (ll *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{
		Val:  value,
		Next: nil,
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

func (node *Node) Print() {
	if node == nil {
		fmt.Println("Empty Node")
	}

	currentNode := node
	for currentNode != nil {
		fmt.Printf("%d->", currentNode.Val)
		currentNode = currentNode.Next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) Display() {
	if ll.Head == nil {
		fmt.Println("Linked List is empty")
	}

	currentNode := ll.Head
	for currentNode != nil {
		fmt.Printf("%d->", currentNode.Val)
		currentNode = currentNode.Next
	}
	fmt.Println("nil")
}
