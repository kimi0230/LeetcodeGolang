// 時間複雜度:
// 空間複雜度:
/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 產生 dummyHead,跟 preslow
// 使用雙指針, 先讓 fast走 `k` 步, 然後 `fast slow 同速前進`
// 這樣當fast走到nil時, slow所在位置就是在倒數第 k 的節點
// 將 slow的前一步(preslow)的next 指向 slow.Next
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	preSlow, slow, fast := dummyHead, head, head
	for fast != nil {
		if n <= 0 {
			preSlow = slow
			slow = slow.Next
		}
		fast = fast.Next
		n--
	}
	preSlow.Next = slow.Next
	return dummyHead.Next
}

// @lc code=end

