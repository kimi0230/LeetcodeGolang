/*
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
深度優先
若二元樹為空回傳空, 否則從根結點開始, 先走訪根節點的左子樹 -> 根節點 -> 右子樹
*/

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
			s.Push(p)
			p = p.Left
		} else {
			tmp := s.Pop().(*TreeNode)
			*r = append(*r, tmp.Val)
			p = tmp.Right
		}
	}
}
