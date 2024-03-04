package removenthnodefromendoflist

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

// 產生 dummyHead,跟 preslow
// 使用雙指針, 先讓 fast走 `k` 步, 然後 `fast slow 同速前進`
// 這樣當fast走到nil時, slow所在位置就是在倒數第 k 的節點
// 將 slow的前一步(preslow)的next 指向 slow.Next
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	preSlow, slow, fast := dummyHead, head, head
	for fast != nil {
		if n <= 0 {
			// 先讓 fast走 `k` 步, 然後 `fast slow 同速前進`
			// 這樣當fast走到nil時, slow所在位置就是在倒數第 k 的節點
			preSlow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next
	return dummyHead.Next
}
