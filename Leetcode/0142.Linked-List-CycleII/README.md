# [142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)

## 題目

Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.

 
Example 1:
![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)
```
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```

Example 2:
![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)
```
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```

Example 3:
![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)
```
Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
``` 

Constraints:

* The number of the nodes in the list is in the range [0, 104].
* -105 <= Node.val <= 105
* pos is -1 or a valid index in the linked-list.

## 題目大意
已知linked list 中有含有環, 返回這個環的起始位置

## 解題思路
假設第一次相遇, slow 走了`k`步, 那麼 `fast就走了 2k`, 也就是說fast比slow多走了k步(環長度的整數倍)
設相遇點與環的起點距離為 `m`, 那麼`環的起點`與`Head`的距離為 `k-m`
![](https://github.com/kimi0230/assets/blob/master/leetcode/images/0142.LinkedListCycleII.jpg?raw=true)

## 來源
* https://leetcode.com/problems/linked-list-cycle-ii/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0142.Linked-List-CycleII/main.go

```go
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

```