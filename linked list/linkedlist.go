package main

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	value int
	next  *Node
}

func NewList() LinkedList {
	return LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (list *LinkedList) Add(value int) {
	newNode := Node{
		value: value,
		next:  nil,
	}

	if list.size == 0 {
		list.head = &newNode
		list.tail = &newNode
		list.size++
		return
	}

	if list.size == 1 {
		list.head.next = &newNode
		list.tail = &newNode
		list.size++
		return
	}

	list.tail.next = &newNode
	list.tail = &newNode
	list.size++
}

func (list *LinkedList) Display() {
	if list.size == 0 {
		fmt.Println("list is empty")
	}

	iterator := new(Node)
	for iterator = list.head; iterator != nil; iterator = iterator.next {
		fmt.Println(iterator.value)
	}
}

func main() {
	list := NewList()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	list.Display()
}
