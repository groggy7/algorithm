package main

import "fmt"

var (
	ErrQueueAlreadyFull  = fmt.Errorf("queue is already full")
	ErrQueueAlreadyEmpty = fmt.Errorf("queue is already empty")
)

type Queue struct {
	elements []int
	size     int
	front    int
	rear     int
}

func NewQueue(size int) Queue {
	return Queue{
		elements: make([]int, 0),
		size:     size,
		front:    -1,
		rear:     -1,
	}
}

func (q *Queue) isEmpty() bool {
	return q.front == -1
}

func (q *Queue) isFull() bool {
	return q.size == len(q.elements)
}

func (q *Queue) Enqueue(element int) error {
	if q.isFull() {
		return ErrQueueAlreadyFull
	}

	if q.front == -1 {
		q.front = 0
	}
	q.elements = append(q.elements, element)
	q.rear++
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.isEmpty() {
		return 0, ErrQueueAlreadyEmpty
	}

	q.front--
	element := q.elements[0]
	q.elements = q.elements[1:]
	if len(q.elements) == 0 {
		q.front = -1
		q.rear = -1
	}
	return element, nil
}

func (q *Queue) Display() {
	fmt.Println("displaying the queue --> ")
	for i, elem := range q.elements {
		fmt.Printf("%d. element: %d\n", i, elem)
	}
}

func main() {
	queue := NewQueue(5)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	a, err := queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	queue.Display()
}
