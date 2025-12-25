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
