---
title: 0094.Binary Tree Inorder Traversal
tags: Medium, Stack
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)

## 題目
Given a binary tree, return the inorder traversal of its nodes' values.

Example :
```c
Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
```

Follow up: Recursive solution is trivial, could you do it iteratively?

## 題目大意
中序遍歷一顆樹

## 解題思路
深度優先 中序遍歷
	若二元樹為空回傳空, 否則從根結點開始, 先走訪根節點的左子樹 -> 根節點 -> 右子樹
深度優先, 前序遍歷
 	若二元樹為空回傳空, 否則先根節點-> 左子樹 -> 右子樹
深度優先, 後序遍歷
	若二元樹為空回傳空, 否則從左到右誒並從葉子節點後續走訪左子樹到右子樹, 最後是拜訪根節點

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0094.Binary-Tree-Inorder-Traversal/
* https://leetcode-cn.com/problems/binary-tree-inorder-traversal/


## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0094.Binary-Tree-Inorder-Traversal/Binary-Tree-Inorder-Traversal.go

```go
package binarytreeinordertraversal

import (
	"LeetcodeGolang/Utility/structures"
)

type TreeNode = structures.TreeNode

func InorderTraversalStack(root *TreeNode) []int {
	var result []int
	inorderStack(root, &result)
	return result
}

func InorderTraversalRecursion(root *TreeNode) []int {
	var result []int
	inorderRecursion(root, &result)
	return result
}

func inorderRecursion(root *TreeNode, r *[]int) {
	if root != nil {
		inorderRecursion(root.Left, r)
		*r = append(*r, root.Val)
		inorderRecursion(root.Right, r)
	}
}

func inorderStack(root *TreeNode, r *[]int) {
	s := structures.NewArrayStack()
	p := root

	for !s.IsEmpty() || p != nil {
		if p != nil {
			// 把中間放入stack, 往左節點繼續往下找
			s.Push(p)
			p = p.Left
		} else {
			// 找不到時,印出
			tmp := s.Pop().(*TreeNode)
			*r = append(*r, tmp.Val)
			p = tmp.Right
		}
	}
}
```