package main

import "fmt"

type Queue struct {
	names []string
}

func NewQueue() *Queue {
	return &Queue{
		names: []string{},
	}
}

func (q *Queue) Push(name string) {
	q.names = append(q.names, name)
}

func (q *Queue) Pop() string {
	var name string
	name, q.names = q.names[0], q.names[1:]
	return name
}

func (q *Queue) IsEmpty() bool {
	return len(q.names) == 0
}

func main() {

	q := NewQueue()
	q.Push("Juan")
	q.Push("Pedro")
	q.Push("Gabriel")

	for !q.IsEmpty() {
		fmt.Println(q.Pop())
	}
}
