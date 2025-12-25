---
title: 110. Balanced Binary Tree
subtitle: "https://leetcode.com/problems/balanced-binary-tree/"
date: 2024-01-02T17:21:00+08:00
lastmod: 2024-01-02T17:21:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0110.Balanced-Binary-Tree"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Easy
  - Balanced Binary Tree
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
# [110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/)

## 題目
Given a binary tree, determine if it is 
height-balanced
.


Example 1:
![](balance_1.jpg)


Input: root = [3,9,20,null,null,15,7]
Output: true



Example 2:
![](balance_2.jpg)

Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
Example 3:

Input: root = []
Output: true
 

Constraints:

The number of nodes in the tree is in the range [0, 5000].
-104 <= Node.val <= 104

## 題目大意

判斷一棵樹是不是平衡二叉樹。 平衡二叉樹的定義是：樹中每個節點都滿足左右兩個子樹的高度差 <= 1 的這個條件。

## 解題思路

高度運算可使用 [0104.Maximum-Depth-of-Binary-Tree](../0104.Maximum-Depth-of-Binary-Tree/README.md)

平衡二叉樹（Balanced Binary Tree）是一種二叉樹，其左右子樹的高度差不超過一的二叉樹。換句話說，對於樹中的每個節點，其左子樹和右子樹的高度差不得大於1。

平衡二叉樹的主要目的是確保樹的高度平衡，這有助於保持在最壞情況下的查找、插入和刪除操作的時間複雜度在O(log n)範圍內，其中n是樹中節點的數量。這種平衡性有助於確保樹的性能良好，並減少操作的平均時間複雜度。

## Big O
時間複雜 : `O(n)`
空間複雜 : `O(1)`

## 來源
* https://leetcode.com/problems/balanced-binary-tree/
* https://leetcode.cn/problems/balanced-binary-tree/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0110.Balanced-Binary-Tree/main.go

```go
package balancedbinarytree

import "LeetcodeGolang/structures"

// 時間複雜 O(), 空間複雜 O()

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = structures.TreeNode

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHight := depth(root.Left)
	rightHight := depth(root.Right)

	return abs(leftHight-rightHight) <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```

##  Benchmark

```sh

```