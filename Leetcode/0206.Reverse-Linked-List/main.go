package reverselinkedlist

import (
	. "LeetcodeGolang/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func ReverseListRecursively(head *ListNode) *ListNode {
	return ListRecursivelyChild(head, nil)
}

func ListRecursivelyChild(current *ListNode, prev *ListNode) *ListNode {
	if current == nil {
		return prev
	}
	next := current.Next
	current.Next = prev
	return ListRecursivelyChild(next, current)
}
