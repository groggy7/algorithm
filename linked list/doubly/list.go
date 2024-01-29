package main

import (
	"fmt"
	"log"
)

var (
	ErrEmptyList = fmt.Errorf("list is empty")
)

type DoublyLinkedList struct {
	head *Node
	size int
}

type Node struct {
	value int
	prev  *Node
	next  *Node
}

func NewDoublyLinkedList() DoublyLinkedList {
	return DoublyLinkedList{
		size: 0,
	}
}

func (list *DoublyLinkedList) Insert(value int) {
	node := &Node{value: value, prev: nil, next: nil}

	if list.size == 0 {
		list.head = node
		list.size++
		return
	}

	iterator := list.head
	for iterator.next != nil {
		iterator = iterator.next
	}

	iterator.next = node
	node.prev = iterator
	list.size++
}

func (list *DoublyLinkedList) Remove(index int) error {
	if list.size == 0 {
		return ErrEmptyList
	}

	if index < 0 || index >= list.size {
		return fmt.Errorf("index out of range")
	}

	if index == 0 {
		list.head = list.head.next
		list.head.prev = nil
		list.size--
		return nil
	}

	iterator := list.head
	for i := 0; i < index; i++ {
		iterator = iterator.next
	}

	if index == list.size-1 {
		iterator.prev.next = nil
		list.size--
		return nil
	}

	iterator.prev.next = iterator.next
	iterator.next.prev = iterator.prev
	list.size--
	return nil
}

func (list *DoublyLinkedList) Display() error {
	if list.size == 0 {
		return ErrEmptyList
	}

	for iterator := list.head; iterator != nil; iterator = iterator.next {
		fmt.Printf("%d-->", iterator.value)
	}
	fmt.Println()
	return nil
}
func main() {
	list := NewDoublyLinkedList()

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	list.Display()

	if err := list.Remove(0); err != nil {
		log.Fatal(err)
	}
	list.Display()

	if err := list.Remove(3); err != nil {
		log.Fatal(err)
	}
	list.Display()

	if err := list.Remove(1); err != nil {
		log.Fatal(err)
	}
	list.Display()

	list.Insert(3)
	list.Display()
}
