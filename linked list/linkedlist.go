package main

import "fmt"

type LinkedList struct {
	head *Node
	size int
}

type Node struct {
	value int
	next  *Node
}

func NewList() LinkedList {
	return LinkedList{
		head: nil,
		size: 0,
	}
}

func (list *LinkedList) Add(value int, index int) error {
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

func (list *LinkedList) Remove(index int) error {
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

	i := 1
	for iterator := list.head; iterator != nil; iterator = iterator.next {
		if index == i {
			iterator.next = iterator.next.next
			list.size--
			return nil
		}
		i++
	}
	return nil
}

func (list *LinkedList) Display() {
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

	list.Add(1, 0)
	list.Add(2, 1)
	list.Add(3, 2)
	list.Add(4, 3)
	list.Display()

	list.Remove(1)
	list.Display()

	list.Remove(0)
	list.Display()

	list.Add(5, 2)
	list.Display()

	list.Remove(2)
	list.Display()
}
