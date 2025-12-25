package invertbinarytree

import (
	"LeetcodeGolang/Utility/structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func InvertTree(root *structures.TreeNode) *structures.TreeNode {
	if root == nil {
		return nil
	}

	InvertTree(root.Left)
	InvertTree(root.Right)

	root.Left, root.Right = root.Right, root.Left
	return root
}

func InvertTree2(root *structures.TreeNode) *structures.TreeNode {

	if root != nil {
		root.Left, root.Right = InvertTree2(root.Right), InvertTree2(root.Left)
	}

	return root
}

func InvertTree3(root *structures.TreeNode) *structures.TreeNode {
	queue := make([]*structures.TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		current.Left, current.Right = current.Right, current.Left

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return root
}
