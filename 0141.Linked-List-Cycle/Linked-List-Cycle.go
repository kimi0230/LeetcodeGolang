/*
# [141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/description/)

## 題目

Given a linked list, determine if it has a cycle in it.

Follow up:
Can you solve it without using extra space?

## 題目大意

判斷鍊表是否有環，不能使用額外的空間。

## 解題思路

給 2 個指針，一個指針是另外一個指針的下一個指針。快指針一次走 2 格，慢指針一次走 1 格。如果存在環，那麼前一個指針一定會經過若干圈之後追上慢的指針。
*/

package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}
