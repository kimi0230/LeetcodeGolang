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

// 時間複雜 : `O(n)` ; 空間複雜 : `O(1)`
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHight := depth(root.Left)
	rightHight := depth(root.Right)

	// return abs(leftHight-rightHight) <= 1
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
