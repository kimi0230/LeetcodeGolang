package addtwonumbers

// 時間複雜 O(max(m,n)), 空間複雜 O(1)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 遍歷 l1跟 l2. 講兩個list的val相加, 並且記錄進位的值給next使用
// 最後如果 carry 還有的話, 需要產生一個新的節點
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		if result == nil {
			result = &ListNode{Val: sum, Next: nil}
			tail = result
		} else {
			tail.Next = &ListNode{Val: sum, Next: nil}
			tail = tail.Next
		}
	}
	// 最後如果 carry 還有的話, 需要產生一個新的節點
	if carry > 0 {
		tail.Next = &ListNode{Val: carry, Next: nil}
	}
	return result
}
