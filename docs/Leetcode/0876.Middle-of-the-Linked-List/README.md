---
title: 0876.Middle of the Linked List
tags:
  - Easy
  - Linked List
  - Two Pointers
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [876. Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/)

## 題目
Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.

 

Example 1:


```
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.
```

Example 2:

```
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
```

Constraints:

The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100


## 題目大意
尋找無環的 linked list 的中間節點

## 解題思路
使用 two pointer. 讓快指針往前兩步. 慢指針往前一步. 當快指針到達盡頭時, 慢指針就位於linked list的中間位置.
當linked list的長度為奇數時, slow剛好停在中點位置;
如果長度為偶數, slow停在中間偏右

## 來源
* https://leetcode.com/problems/middle-of-the-linked-list/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0876.Middle-of-the-Linked-List/main.go

```go
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
```

###### tags: `Medium` `Leetcode` `two pointers`