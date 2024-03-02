---
title: 0002.Add Two Numbers
subtitle: "https://leetcode.com/problems/add-two-numbers/description/"
date: 2024-03-02T15:57:00+08:00
lastmod: 2024-03-02T15:57:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0002.Add-Two-Numbers"
license: ""
images: []

tags: [LeetCode, Go, Medium, Add Two Numbers, Linked List, Math, Recursion, Amazon, Apple,Facebook,Microsoft,Bloomberg]
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
# [0002.Add Two Numbers](https://leetcode.com/problems/add-two-numbers/description/)

## 題目

## 題目大意


## 解題思路
[鏈表雙指標技巧](https://labuladong.github.io/article/fname.html?fname=%E9%93%BE%E8%A1%A8%E6%8A%80%E5%B7%A7) 和加法運算過程中對進位的處理。 注意這個 carry 變數的處理，在我們手動類比加法過程的時候會經常用到。

* 遍歷 l1跟 l2. 講兩個list的val相加, 並且記錄進位的值給next使用
* 最後如果 carry 還有的話, 需要產生一個新的節點

## Big O

* 時間複雜 : `O(max⁡(m,n)`
時間複雜度： O(max⁡(m,n)) ，其中 m 和 n 分別為兩個鏈表的長度。 我們要遍歷兩個鏈表的全部位置，而處理每個位置只需要 O(1) 的時間

* 空間複雜 : `O(1)`
O(1) 。 注意返回值不計入空間複雜度

## 來源
* https://leetcode.com/problems/add-two-numbers/description/
* https://leetcode.cn/problems/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0002.Add-Two-Numbers/main.go

```go
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
```

##  Benchmark

```sh

```