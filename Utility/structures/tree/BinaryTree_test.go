package tree

import (
	"testing"
)

// 	   1
//   /   \
//  3    4
//      /
// 	   5
func TestBinaryTree_BFSTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)
	binaryTree.BFSTraverse() // 1 3 4 5
}

// 前序遍歷
func TestBinaryTree_PreOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.PreOrderTraverse() //1 3 4 5
}

func TestBinaryTree_PreOrderTraverseRecursive(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.root.PreOrderTraverseRecursive() //1 3 4 5
}

//中序遍歷
func TestBinaryTree_InOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.InOrderTraverse() // 3 1 5 4
}

//中序遍歷 遞迴
func TestBinaryTree_InOrderTraverseRecursive(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.root.InOrderTraverseRecursive() // 3 1 5 4
}

//後序遍歷
func TestBinaryTree_PostOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.PostOrderTraverse() // 3 5 4 1
}

//後序遍歷
func TestBinaryTree_PostOrderTraverse2(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.PostOrderTraverse2() // 3 5 4 1
}

//後序遍歷 遞迴
func TestBinaryTree_PostOrderTraverseRecursive(t *testing.T) {
	binaryTree := NewBinaryTree(1)
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.root.PostOrderTraverseRecursive() // 3 5 4 1
}
