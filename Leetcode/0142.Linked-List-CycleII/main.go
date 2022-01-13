package linkedlistcycleII

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

func DetectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			// 找到相遇點
			break
		}
	}

	slow = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
