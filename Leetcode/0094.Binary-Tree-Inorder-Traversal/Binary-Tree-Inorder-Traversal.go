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
