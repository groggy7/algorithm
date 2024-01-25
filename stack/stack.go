package main

import "fmt"

var (
	ErrStackAlreadyFull  = fmt.Errorf("stack is already full")
	ErrStackAlreadyEmpty = fmt.Errorf("stack is already empty")
)

type Stack struct {
	elements []int
	top      int
	capacity int
}

func NewStack(cap int) Stack {
	return Stack{
		elements: make([]int, cap),
		top:      -1,
		capacity: cap,
	}
}

func (s *Stack) Push(a int) error {
	if s.isFull() {
		return ErrStackAlreadyFull
	}

	s.top++
	//After pop operation length of the slice decreases and setting element
	//by index throws index out of range expression. That's for it.
	if len(s.elements) < s.capacity {
		s.elements = append(s.elements, a)
		return nil
	}

	s.elements[s.top] = a
	return nil
}

func (s *Stack) Pop() error {
	if s.isEmpty() {
		return ErrStackAlreadyEmpty
	}

	s.elements = s.elements[:len(s.elements)-1]
	s.top--
	return nil
}

func (s *Stack) isFull() bool {
	return s.top == s.capacity-1
}

func (s *Stack) isEmpty() bool {
	return s.top == -1
}

func (s *Stack) printStack() {
	for i := 0; i < len(s.elements); i++ {
		fmt.Println(s.elements[i])
	}
}

func main() {
	stack := NewStack(4)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	stack.Pop()
	stack.Push(5)
	stack.printStack()
}
