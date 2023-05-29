package main

import "fmt"

// Element are the elements that our stack will hold inside
type Element string

// String representation of our Element
func (element Element) String() string {
	return string(element)
}

// Stack is our data structure
type Stack struct {
	elements []Element
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{
		elements: make([]Element, 0),
	}
}

// Push adds a element to the stack.
func (s *Stack) Push(e Element) {
	s.elements = append(s.elements, e)
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Element {
	var e *Element
	e, s.elements = &s.elements[len(s.elements)-1], s.elements[:len(s.elements)-1]
	return e
}

// Iter will return a channel that we will use to iterate over our elements in a for loop.
// The for loop iterate go to remove elements from the stack using the Pop method.
// Also, you can use the Pop method alone if you want.
func (s *Stack) Iter() chan *Element {
	c := make(chan *Element)
	go func() {
		for range s.elements {
			c <- s.Pop()
		}
		close(c)
	}()
	return c
}

func main() {

	// The stack data structure is a linear data structure that accompanies a principle known as  LIFO (Last In First Out)
	// or FILO (First In Last Out). Real-life examples of a stack are a deck of cards, piles of books, piles of money, and
	// many more.
	stack := NewStack()
	// Add elements to the stack
	stack.Push("Primero")
	stack.Push("Segundo")
	stack.Push("Tercero")
	stack.Push("Cuarto")
	stack.Push("Quinto")

	fmt.Println("Initial stack: ", stack.elements)

	// With this loop, remove all the elements from the stack, iterating over all elements, using the Pop method.
	for element := range stack.Iter() {
		fmt.Println(element)
	}
	fmt.Println("End stack: ", stack.elements)
}
