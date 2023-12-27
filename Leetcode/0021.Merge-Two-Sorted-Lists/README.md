---
title: 0021. Merge Two Sorted Lists
subtitle: "https://leetcode.com/problems/merge-two-sorted-lists/description/"
date: 2023-11-05T21:30:00+08:00
lastmod: 2023-11-05T21:30:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0021.Merge-Two-Sorted-Lists"
license: ""
images: []

tags: [LeetCode, Go, Easy/Medium/Hard, Merge Two Sorted Lists]
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
# [0021. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/description/)

## 題目

## 題目大意


## 解題思路

## Big O
時間複雜 : `O( log n)`
空間複雜 : `O(1)`

## 來源
* https://leetcode.com/problems/merge-two-sorted-lists/description/
* https://leetcode.cn/problems/merge-two-sorted-lists/description/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0021.Merge-Two-Sorted-Lists/main.go

```go
package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

// 時間複雜 O(log n), 空間複雜 O(1)
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{Next: nil}
	cur := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return head.Next
}

```

##  Benchmark

```sh

```