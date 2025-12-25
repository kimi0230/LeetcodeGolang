package constructbinarytreefrompreorderandinordertraversal

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
