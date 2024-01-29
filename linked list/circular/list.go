// singly circular linked list
package main

import "fmt"

type CircularLinkedList struct {
	head *Node
	size int
}

type Node struct {
	value int
	next  *Node
}

func NewCircularLinkedList() CircularLinkedList {
	return CircularLinkedList{
		head: nil,
		size: 0,
	}
}

func (list *CircularLinkedList) Insert(value int) {
	node := &Node{value: value, next: nil}

	if list.size == 0 {
		list.head = node
		list.head.next = list.head
		list.size++
		return
	}

	iterator := list.head
	for iterator.next != list.head {
		iterator = iterator.next
	}

	iterator.next = node
	node.next = list.head
	list.size++
}

func (list *CircularLinkedList) Remove(index int) error {
	if list.size == 0 {
		return fmt.Errorf("cannot remove from an empty list")
	}

	if index < 0 || index >= list.size {
		return fmt.Errorf("index out of range")
	}

	if list.size == 1 {
		list.head = nil
		list.size = 0
		return nil
	}

	iterator := list.head
	var prev *Node

	for i := 0; i < index; i++ {
		prev = iterator
		iterator = iterator.next

		if iterator == list.head {
			return fmt.Errorf("index out of range")
		}
	}

	if prev != nil {
		prev.next = iterator.next
	} else {
		list.head = iterator.next
		lastNode := list.head
		for lastNode.next != iterator {
			lastNode = lastNode.next
		}
		lastNode.next = list.head
	}

	list.size--

	return nil
}

func (list *CircularLinkedList) Display() error {
	if list.size == 0 {
		return fmt.Errorf("list is empty")
	}

	iterator := list.head
	for {
		fmt.Printf("%d-->", iterator.value)
		iterator = iterator.next

		if iterator == list.head {
			break
		}
	}

	fmt.Println()
	return nil
}

func main() {
	list := NewCircularLinkedList()

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	list.Remove(1)
	list.Remove(3)
	list.Remove(0)
	list.Display()
}
