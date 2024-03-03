// 時間複雜度:
// 空間複雜度:
/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 先用快慢指針找出Linked list的中間點
// 然後把Linked list分成兩半
// 再把後半的Linked list反轉
// 再把兩半的Linked list merge 起來
func reorderList(head *ListNode) {
	mid := middleNode(head)

	// 2.反轉中間節點的下一個節點
	right := resverseNode(mid.Next)
	mid.Next = nil
	left := head
	mergeNode(left, right)
}

// [876. Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/)
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)
func resverseNode(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}

func mergeNode(l1, l2 *ListNode) {
	lcur, rcur := l1, l2
	for lcur != nil && rcur != nil {
		lcur.Next, rcur.Next, lcur, rcur = rcur, lcur.Next, lcur.Next, rcur.Next
	}
}

// @lc code=end

