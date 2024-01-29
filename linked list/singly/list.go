package main

import "fmt"

type SinglyLinkedList struct {
	head *Node
	size int
}

type Node struct {
	value int
	next  *Node
}

func NewList() SinglyLinkedList {
	return SinglyLinkedList{
		head: nil,
		size: 0,
	}
}

func (list *SinglyLinkedList) Insert(value int, index int) error {
	if index < 0 || index > list.size {
		return fmt.Errorf("index out of range")
	}

	node := &Node{value: value, next: nil}

	if index == 0 {
		node.next = list.head
		list.head = node
		list.size++
		return nil
	}

	current := list.head
	for i := 1; i < index; i++ {
		current = current.next
	}

	if index == list.size-1 {
		current.next = node
		list.size++
	}

	current.next = node
	node.next = current.next.next
	list.size++
	return nil
}

func (list *SinglyLinkedList) Remove(index int) error {
	if list.size == 0 {
		return fmt.Errorf("cannot remove from an empty list")
	}

	if index < 0 || index >= list.size {
		return fmt.Errorf("index out of range")
	}

	if index == 0 {
		list.head = list.head.next
		list.size--
		return nil
	}

	iterator := list.head
	for i := 1; i < index-1; i++ {
		iterator = iterator.next
		i++
	}

	if index == list.size-1 {
		iterator.next = nil
		list.size--
		return nil
	}

	iterator.next = iterator.next.next
	list.size--
	return nil
}

func (list *SinglyLinkedList) Display() {
	if list.head == nil {
		fmt.Println("list is empty")
		return
	}

	fmt.Printf("list has %d items: \n", list.size)
	for iterator := list.head; iterator != nil; iterator = iterator.next {
		fmt.Printf("%d---->", iterator.value)
	}
	fmt.Println()
}

func main() {
	list := NewList()

	list.Insert(1, 0)
	list.Insert(2, 1)
	list.Insert(3, 2)
	list.Insert(4, 3)
	list.Display()

	list.Remove(1)
	list.Display()

	list.Remove(0)
	list.Display()

	list.Remove(1)

	list.Insert(5, 1)
	list.Display()

	list.Remove(1)
	list.Display()
}
