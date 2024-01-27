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
	} else if list.size == 1 {
		list.head.next = &newNode
		list.tail = &newNode
		list.size++
		return
	}

	list.tail.next = &newNode
	list.tail = &newNode
	list.size++
}

func (list *LinkedList) Remove(index int) error {
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
	if list.size == 0 {
		fmt.Println("list is empty")
	}

	fmt.Printf("list has %d items\n", list.size)
	for iterator := list.head; iterator != nil; iterator = iterator.next {
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

	list.Remove(1)
	list.Display()

	list.Remove(0)
	list.Display()

	list.Add(5)
	list.Display()

	list.Remove(2)
	list.Display()
}
