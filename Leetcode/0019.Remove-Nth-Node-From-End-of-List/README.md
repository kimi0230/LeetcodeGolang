---
title: 0019. Remove Nth Node From End of List
subtitle: "https://leetcode.com/problems/remove-nth-node-from-end-of-list/"
date: 2024-03-12T16:26:00+08:00
lastmod: 2024-03-12T16:26:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0015.3Sum"
license: ""
images: []

tags: [LeetCode, Go, Medium, Linked List, Two Pointers, Facebook, Amazon, Microsoft, Google, Bloomberg]
categories: [LeetCode]

featuredImage: ""
featuredImagePreview: ""

hiddenFromHomePage: false
hiddenFromSearch: false
twemoji: false
lightgallery: true
ruby: true
fraction: true
fontawesome: true
linkToMarkdown: false
rssFullText: false

toc:
  enable: true
  auto: true
code:
  copy: true
  maxShownLines: 200
math:
  enable: false
  # ...
mapbox:
  # ...
share:
  enable: true
  # ...
comment:
  enable: true
  # ...
library:
  css:
    # someCSS = "some.css"
    # located in "assets/"
    # Or
    # someCSS = "https://cdn.example.com/some.css"
  js:
    # someJS = "some.js"
    # located in "assets/"
    # Or
    # someJS = "https://cdn.example.com/some.js"
seo:
  images: []
  # ...
---
# [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

## 題目
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
![](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)
```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```

Example 2:
```
Input: head = [1], n = 1
Output: []
```

Example 3:
```
Input: head = [1,2], n = 1
Output: [1]
```

Constraints:
* The number of nodes in the list is sz.
* 1 <= sz <= 30
* 0 <= Node.val <= 100
* 1 <= n <= sz


## 題目大意
找尋單linked list的 倒數第 n 個元素並刪除.
返回該 linked list的頭節點

## 解題思路
先讓 fast走 `k` 步, 然後 `fast slow 同速前進`
這樣當fast走到nil時, slow所在位置就是在倒數第 k 的節點

## 來源
* https://leetcode.com/problems/remove-nth-node-from-end-of-list/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0019.Remove-Nth-Node-From-End-of-List/main.go

```go
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
```


###### tags: `Medium` `Leetcode` `Two Pointers`