package maximumdepthofbinarytree

import (
	"LeetcodeGolang/structures"
	"math"
)

type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 時間複雜 O(n), 空間複雜 O(1)
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(MaxDepth(root.Left)), float64(MaxDepth(root.Right)))) + 1

}

// 時間複雜 O(n), 空間複雜 O(1)
func MaxDepthKimi(root *TreeNode) int {
	count := 0
	return depth(root, count)

}

func depth(root *TreeNode, count int) int {
	if root != nil {
		count++
		countL := depth(root.Left, count)
		countR := depth(root.Right, count)
		if countL > countR {
			count = countL
		} else {
			count = countR
		}
	}
	return count
}
