---
title: 105. Construct Binary Tree from Preorder and Inorder Traversal
subtitle: "https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/"
date: 2024-02-20T16:56:00+08:00
lastmod: 2024-02-20T16:56:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0105.Construct-Binary-Tree-from-Preorder-and-Inorder-Traversal"
license: ""
images: []

tags: [LeetCode, Go, Medium, Construct Binary Tree from Preorder and Inorder Traversal, Array,Hash Table,Divide and Conquer,Tree, Binary Tree]
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
# [105. Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/)

## 題目

## 題目大意


## 解題思路

## Big O
* 時間複雜 : `O(n)`
n個樹節點

* 空間複雜 : `O(n)`
遞迴調用的Stack空間是 O(h)，其中 h 是樹的高度。
在每個遞迴層級中，都創建了一個 TreeNode 對象，因此空間複雜度也是 O(n)，其中 n 是節點的數量。 
h< n所以空間複雜為 O(n)

## 來源
* https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
* https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0105.Construct-Binary-Tree-from-Preorder-and-Inorder-Traversal/main.go

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 時間複雜 O(), 空間複雜 O()
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	result := &TreeNode{preorder[0], nil, nil}

	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	result.Left = buildTree(preorder[1:i+1], inorder[:i])
	result.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return result
}
```
![](105.construct%20Binary%20Tree%20From%20Preorder%20And%20In%20order%20Traversal-2.jpg)

##  Benchmark

```sh

```