package main

import (
	"fmt"
	"strings"
)

/*
https://leetcode.com/problems/swap-nodes-in-pairs/description/

Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

Input: head = [1,2,3,4]
Output: [2,1,4,3]

Example 2:

Input: head = []
Output: []

Example 3:

Input: head = [1]
Output: [1]

Constraints:

The number of nodes in the list is in the range [0, 100].
0 <= Node.val <= 100
*/
func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	output := []int{2, 1, 4, 3, 6, 5}
	head := NewListNode(input)
	fmt.Printf("head: %d output %d result %s\n", input, output, swapPairs(head))

	input = []int{}
	output = []int{}
	head = NewListNode(input)
	fmt.Printf("head: %d output %d result %s\n", input, output, swapPairs(head))

	input = []int{1}
	output = []int{1}
	head = NewListNode(input)
	fmt.Printf("head: %d output %d result %s\n", input, output, swapPairs(head))
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head, head.Next, head.Next.Next = head.Next, swapPairs(head.Next.Next), head

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "[]" // Empty list
	}

	values := []string{}
	for curr := l; curr != nil; curr = curr.Next {
		values = append(values, fmt.Sprintf("%d", curr.Val))
	}
	return fmt.Sprintf("[%s]", strings.Join(values, ", "))
}

func NewListNode(data []int) *ListNode {
	var head *ListNode // Head of the linked list
	for _, val := range data {
		newNode := &ListNode{Val: val} // Create a new ListNode with the value
		if head == nil {
			head = newNode // Set head for the first element
		} else {
			// Traverse to the last node and append the new node
			curr := head
			for curr.Next != nil {
				curr = curr.Next
			}
			curr.Next = newNode
		}
	}
	return head
}
