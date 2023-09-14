package main

import "fmt"

/*
	Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

	Example 1:

	Input: l1 = [2,4,3], l2 = [5,6,4]
	Output: [7,0,8]
	Explanation: 342 + 465 = 807.

	Example 2:

	Input: l1 = [0], l2 = [0]
	Output: [0]

	Example 3:

	Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	Output: [8,9,9,9,0,0,0,1]

	Input: l1 = [2,4,9] l2 = [5,6,4,9]
	Output: [7,0,4,0,1]

Constraints:

	The number of nodes in each linked list is in the range [1, 100].
	0 <= Node.val <= 9
	It is guaranteed that the list represents a number that does not have leading zeros.
*/
func main() {

	l1 := NewListNode(2, 4, 9)
	l2 := NewListNode(5, 6, 4, 9)
	expected := []int{7, 0, 4, 0, 1}
	fmt.Printf("L1 = %s, L2 = %s\n", l1, l2)
	result := addTwoNumbers(l1, l2)
	fmt.Printf("RESULT: %s, EXPECTED: %d\n\n", result, expected)

	l1 = NewListNode(9, 9, 9, 9, 9, 9, 9)
	l2 = NewListNode(9, 9, 9, 9)
	expected = []int{8, 9, 9, 9, 0, 0, 0, 1}
	fmt.Printf("L1 = %s, L2 = %s\n", l1, l2)
	result = addTwoNumbers(l1, l2)
	fmt.Printf("RESULT: %s, EXPECTED: %d\n\n", result, expected)

	l1 = NewListNode(0)
	l2 = NewListNode(0)
	expected = []int{0}
	fmt.Printf("L1 = %s, L2 = %s\n", l1, l2)
	result = addTwoNumbers(l1, l2)
	fmt.Printf("RESULT: %s, EXPECTED: %d\n\n", result, expected)

	l1 = NewListNode(2, 4, 3)
	l2 = NewListNode(5, 6, 4)
	expected = []int{7, 0, 8}
	fmt.Printf("L1 = %s, L2 = %s\n", l1, l2)
	result = addTwoNumbers(l1, l2)
	fmt.Printf("RESULT: %s, EXPECTED: %d\n", result, expected)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(v int, list ...int) *ListNode {
	listNode := &ListNode{
		Val:  v,
		Next: nil,
	}

	if len(list) > 0 {
		for _, v := range list {
			listNode.Add(v)
		}
	}

	return listNode
}

func (l *ListNode) Add(v int) {
	newl := NewListNode(v)
	for l.Next != nil {
		l = l.Next
	}
	l.Next = newl
}

func (l *ListNode) String() string {
	list := []int{l.Val}
	for node := l.Next; node != nil; node = node.Next {
		list = append(list, node.Val)
	}
	return fmt.Sprintf("%v", list)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		num := n1 + n2 + carry
		carry = num / 10
		cur.Next = &ListNode{num % 10, nil}
		cur = cur.Next
	}
	return head.Next
}
