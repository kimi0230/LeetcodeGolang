---
title: 143. Reorder List
subtitle: "https://leetcode.com/problems/reorder-list/description/"
date: 2024-03-03T16:53:00+08:00
lastmod: 2024-03-03T16:53:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0143.Reorder-List"
license: ""
images: []

tags: [LeetCode, Go, /Medium, 143. Reorder List, Amazon,Microsoft,Adobe,Facebook,Bloomberg,Linked List,Two Pointers,Stack, Recursion]
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
# [143. Reorder List](https://leetcode.com/problems/reorder-list/description/)

## 題目

## 題目大意


## 解題思路

## Big O

* 時間複雜 : ``
* 空間複雜 : ``

## 來源
* https://leetcode.com/problems/reorder-list/description/
* https://leetcode.cn/problems/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0143.Reorder-List/main.go

```go
package reorderlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 時間複雜 O(), 空間複雜 O()
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

```

##  Benchmark

```sh

```