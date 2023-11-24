package main

import (
	"fmt"
	"strconv"
)

/*
   Remove Nth node from end of list.

   Given the head of a linked list, remove the nth node from the end of the list and return its head.

   Example 1:

   Input: head = [1,2,3,4,5], n = 2
   Output: [1,2,3,5]

   Example 2:

   Input: head = [1], n = 1
   Output: []

   Example 3:

   Input: head = [1,2], n = 1
   Output: [1]

   Constraints:

    The number of nodes in the list is sz.
    1 <= sz <= 30
    0 <= Node.val <= 100
    1 <= n <= sz

*/

func main() {
	input := []int{1, 2, 3, 4, 5}
	nodeInput := newListNode(input)
	n := 3
	output := []int{1, 2, 4, 5}
	nodeOutput := newListNode(output)
	fmt.Printf("Input %v n %d Output %v Result %v\n\n", show(nodeInput), n, show(nodeOutput), show(removeNthFromEnd(nodeInput, n)))

	input = []int{1}
	nodeInput = newListNode(input)
	n = 1
	output = []int{}
	nodeOutput = newListNode(output)
	fmt.Printf("Input %v n %d Output %v Result %v\n\n", show(nodeInput), n, show(nodeOutput), show(removeNthFromEnd(nodeInput, n)))

	input = []int{1, 2}
	nodeInput = newListNode(input)
	n = 1
	output = []int{1}
	nodeOutput = newListNode(output)
	fmt.Printf("Input %v n %d Output %v Result %v\n\n", show(nodeInput), n, show(nodeOutput), show(removeNthFromEnd(nodeInput, n)))
}

func newListNode(n []int) *ListNode {
	if len(n) <= 0 {
		return &ListNode{}
	}
	l := ListNode{Val: n[0]}
	n = n[1:]
	for _, v := range n {
		l.Add(v)
	}
	return &l
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func show(list *ListNode) string {
	l := *list

	if l.Next == nil && l.Val == 0 {
		return "[]"
	}

	r := "["
	for {
		if l.Next != nil {
			r += strconv.Itoa(l.Val) + ","
		} else if l.Next == nil {
			r += strconv.Itoa(l.Val)
			break
		}
		l = *l.Next
	}
	r += "]"
	return r
}

func (l *ListNode) Add(n int) {
	newNode := ListNode{Val: n}
	for {
		if l.Next == nil {
			l.Next = &newNode
			break
		}
		l = l.Next
	}
}

func (l *ListNode) Lenght() int {
	count := 0
	for {
		if l.Next != nil {
			count++
		} else {
			break
		}
	}

	return count
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listNodes := &ListNode{-1, head}

	current, prevOfRemoval := listNodes, listNodes

	for current.Next != nil {

		// n step delay for prevOfRemoval
		if n <= 0 {
			prevOfRemoval = prevOfRemoval.Next
		}

		current = current.Next

		n -= 1
	}

	// Remove the N-th node from end of list
	nthNode := prevOfRemoval.Next
	prevOfRemoval.Next = nthNode.Next

	if listNodes.Next == nil {
		return &ListNode{}
	}
	return listNodes.Next
}
