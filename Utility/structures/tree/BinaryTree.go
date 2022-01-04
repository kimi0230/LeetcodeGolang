package tree

import (
	"fmt"
)

type BinaryTree struct {
	root *Node
}

// type Node struct {
// 	data  interface{}
// 	left  *Node
// 	right *Node
// }

func NewBinaryTree(rootV interface{}) *BinaryTree {
	return &BinaryTree{NewNode(rootV)}
}

// BFSTraverse: 廣度優先, Breadth First Traverse 廣度優先
// 用 Queue實現
func (this *BinaryTree) BFSTraverse() {
	p := this.root
	s := NewQueueStack()

	if p == nil {
		return
	}
	s.Push(p)
	for !s.IsEmpty() {
		node := s.Head().(*Node)

		if node.left != nil {
			s.Push(node.left)
		}
		if node.right != nil {
			s.Push(node.right)
		}

		tmp := s.Pop().(*Node)
		fmt.Printf("%+v ", tmp.data)
	}
	fmt.Println()
}

// PreOrderTraverse : 深度優先, 前序遍歷
// 若二元樹為空回傳空, 否則先根節點-> 左子樹 -> 右子樹
func (this *BinaryTree) PreOrderTraverse() {
	p := this.root
	s := NewArrayStack()

	for !s.IsEmpty() || nil != p {
		if nil != p {
			fmt.Printf("%+v ", p.data)
			s.Push(p)
			p = p.left
		} else {
			p = s.Pop().(*Node).right
		}
	}
}

func (this *Node) PreOrderTraverseRecursive() {
	p := this
	if p == nil {
		return
	}

	fmt.Printf("%+v ", p.data)
	left := p.left
	right := p.right
	left.PreOrderTraverseRecursive()
	right.PreOrderTraverseRecursive()
}

// InOrderTraverse : 深度優先, 中序遍歷
// 若二元樹為空回傳空, 否則從根結點開始, 先走訪根節點的左子樹 -> 根節點 -> 右子樹
func (this *BinaryTree) InOrderTraverse() {
	p := this.root
	s := NewArrayStack()

	for !s.IsEmpty() || nil != p {
		if nil != p {
			s.Push(p)
			p = p.left
		} else {
			tmp := s.Pop().(*Node)
			fmt.Printf("%+v ", tmp.data)
			p = tmp.right
		}
	}
	fmt.Println()
}

// InOrderTraverseRecursive : 中序遍歷 遞迴
func (this *Node) InOrderTraverseRecursive() {
	p := this
	if p == nil {
		return
	}

	left := p.left
	left.PreOrderTraverseRecursive()

	fmt.Printf("%+v ", p.data)

	right := p.right
	right.PreOrderTraverseRecursive()
}

// PostOrderTraverse : 深度優先, 後序遍歷
// 若二元樹為空回傳空, 否則從左到右誒並從葉子節點後續走訪左子樹到右子樹, 最後是拜訪根節點
func (this *BinaryTree) PostOrderTraverse() {
	s1 := NewArrayStack()
	s2 := NewArrayStack()
	s1.Push(this.root)

	for !s1.IsEmpty() {
		p := s1.Pop().(*Node)
		s2.Push(p)
		if p.left != nil {
			s1.Push(p.left)
		}
		if p.right != nil {
			s1.Push(p.right)
		}
	}

	for !s2.IsEmpty() {
		fmt.Printf("%+v ", s2.Pop().(*Node).data)
	}

	fmt.Println()
}

// PostOrderTraverse2 : 深度優先, 後序遍歷
// use one stack, pre cursor to traverse from post order
// 若二元樹為空回傳空, 否則從左到右誒並從葉子節點後續走訪左子樹到右子樹, 最後是拜訪根節點
func (this *BinaryTree) PostOrderTraverse2() {
	r := this.root
	s := NewArrayStack()

	//point to last visit node
	var pre *Node

	s.Push(r)

	for !s.IsEmpty() {
		r = s.Top().(*Node)
		if (r.left == nil && r.right == nil) || (pre != nil && (pre == r.left || pre == r.right)) {
			fmt.Printf("%+v ", r.data)
			s.Pop()
			pre = r
		} else {
			if r.right != nil {
				s.Push(r.right)
			}

			if r.left != nil {
				s.Push(r.left)
			}
		}
	}
	fmt.Println()
}

// PostOrderTraverseRecursive : 後序遍歷 遞迴
func (this *Node) PostOrderTraverseRecursive() {
	p := this
	if p == nil {
		return
	}

	left := p.left
	right := p.right
	left.PreOrderTraverseRecursive()
	right.PreOrderTraverseRecursive()
	fmt.Printf("%+v ", p.data)
}
