---
title: 0543. Diameter of Binary Tree
subtitle: "https://leetcode.com/problems/diameter-of-binary-tree/"
date: 2023-01-22T21:20:00+08:00
lastmod: 2023-01-22T21:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0543.Diameter-of-Binary-Tree"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Easy
  - Diameter of Binary Tree
  - DFS
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
# [543. Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/)

## 題目
Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.

 

Example 1:


Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
Example 2:

Input: root = [1,2]
Output: 1
 

Constraints:

The number of nodes in the tree is in the range [1, 104].
-100 <= Node.val <= 100

## 題目大意


## 解題思路
左邊的最高高度與右邊的最高高度相加


## Big O
時間複雜 `O(n)`, 
空間複雜: 最壞 `O(n)`, 平衡樹 `O(log(n))`

## 來源
* https://leetcode.com/problems/diameter-of-binary-tree/
* https://leetcode.cn/problems/diameter-of-binary-tree/
* 
## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0543.Diameter-of-Binary-Tree/main.go

```go
package diameterofbinarytree

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

var maxDiameter int

// 時間複雜 O(n), 空間複雜: 最壞 O(n), 平衡樹 O(log(n))
func DiameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDiameter = 0
	maxDepth(root)
	return maxDiameter
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := maxDepth(node.Left)
	right := maxDepth(node.Right)

	maxDiameter = max(maxDiameter, left+right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

```

##  Benchmark

```sh

```