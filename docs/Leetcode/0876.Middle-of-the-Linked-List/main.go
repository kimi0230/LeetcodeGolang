package middleofthelinkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 算出長度
	curr := head
	length := 0
	for curr != nil {
		length++
		curr = curr.Next
	}
	if length%2 == 0 {
		// 偶數
		return slow.Next
	} else {
		// 奇數
		return slow
	}
}
