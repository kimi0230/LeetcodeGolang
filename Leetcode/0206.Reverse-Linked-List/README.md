---
title: 206. Reverse Linked List
subtitle: "Reverse Linked List"
date: 2023-01-22T21:20:00+08:00
lastmod: 2023-01-22T21:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Reverse Linked List"
license: ""
images: []

tags: [LeetCode, Go, Easy, Reverse Linked List]
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
# [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)

## 題目
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:
![](https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg)

```
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
```

Example 2:
![](https://assets.leetcode.com/uploads/2021/02/19/rev1ex2.jpg)
```
Input: head = [1,2]
Output: [2,1]
```

Example 3:
```
Input: head = []
Output: []
```

Constraints:

* The number of nodes in the list is the range [0, 5000].
* -5000 <= Node.val <= 5000

Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?

## 題目大意
將 Linked List 反向

## 解題思路

## 來源
* [https://leetcode.com/problems/Reverse Linked List/](https://leetcode.com/problems/reverse-linked-list/)

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0000.kkkk/main.go

```go
package reverselinkedlist

import (
	. "LeetcodeGolang/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func ReverseListRecursively(head *ListNode) *ListNode {
	return ListRecursivelyChild(head, nil)
}

func ListRecursivelyChild(current *ListNode, prev *ListNode) *ListNode {
	if current == nil {
		return prev
	}
	next := current.Next
	current.Next = prev
	return ListRecursivelyChild(next, current)
}

```

##  Benchmark
```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0206.Reverse-Linked-List
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkReverseList-8                  1000000000               0.7960 ns/op          0 B/op          0 allocs/op
BenchmarkReverseListRecursively-8       276534334                4.374 ns/op           0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0206.Reverse-Linked-List        2.597s
```