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
