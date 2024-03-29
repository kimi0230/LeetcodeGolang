---
title: 0203.Remove Linked List Elements
tags: Easy, Linked List
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [203. Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/)

## 題目

Remove all elements from a linked list of integers that have value val.

Example :

```c
Input: 1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
```


## 題目大意

刪除鍊錶中所有指定值的結點。

## 解題思路

按照題意做即可

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0203.Remove-Linked-List-Elements/
* https://leetcode-cn.com/problems/remove-linked-list-elements/
* https://mp.weixin.qq.com/s/slM1CH5Ew9XzK93YOQYSjA

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0203.Remove-Linked-List-Elements/Remove-Linked-List-Elements.go

```go
package removelinkedlistelements

import (
	"LeetcodeGolang/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = structures.ListNode

// 方法一: 另外考慮刪除head節點
func RemoveElements(head *ListNode, val int) *ListNode {
	// 刪除值相同的head
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	tmpHead := head
	for tmpHead.Next != nil {
		if tmpHead.Next.Val == val {
			tmpHead.Next = tmpHead.Next.Next
		} else {
			tmpHead = tmpHead.Next
		}
	}
	return head
}

/*
方法二 添加虛擬節點, 效能較好
可以設置一個虛擬頭結點」，這樣原鍊錶的所有節點就都可以按照統一的方式進行移除了
return newHead.Next
*/
func RemoveElementsVirtualNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	// 建立一個虛擬 Head 節點
	newHead := &ListNode{Val: 0, Next: head}
	preHead := newHead
	curHead := head
	for curHead != nil {
		if curHead.Val == val {
			preHead.Next = curHead.Next
		} else {
			preHead = curHead
		}
		curHead = curHead.Next
	}
	return newHead.Next
}

/*
方法二 遞迴
*/
func RemoveElementsRecurse(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = RemoveElementsRecurse(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}
```