---
title: 100. Same Tree
subtitle: "https://leetcode.com/problems/same-tree/"
date: 2024-01-03T16:13:00+08:00
lastmod: 2024-01-03T16:13:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0100.Same-Tree"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Easy
  - Same Tree
  - Tree
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
# [100. Same Tree](https://leetcode.com/problems/same-tree/)

## 題目

Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

 

Example 1:
![](ex1.jpg)

Input: p = [1,2,3], q = [1,2,3]
Output: true

Example 2:
![](ex2.jpg)

Input: p = [1,2], q = [1,null,2]
Output: false

Example 3:
![](ex3.jpg)

Input: p = [1,2,1], q = [1,1,2]
Output: false
 

Constraints:

The number of nodes in both trees is in the range [0, 100].
-104 <= Node.val <= 104

* Accepted: 1.8M
* Submissions: 3.1M
* Acceptance Rate: 60.3%

## 題目大意
判斷 2 顆樹是否是完全相等的

## 解題思路

遞歸判斷即可

## Big O
時間複雜 : `O(n)`
空間複雜 : `O(1)`

## 來源
* https://leetcode.com/problems/same-tree/
* https://leetcode.cn/problems/same-tree/
* https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0100.Same-Tree/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0100.Same-Tree/main.go

```go
package sametree

import "LeetcodeGolang/structures"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = structures.TreeNode

// 時間複雜 O(n), 空間複雜 O(1)
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val == q.Val {
			return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
		}
		return false
	} else {
		return false
	}
}

```

##  Benchmark

```sh

```